package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAutoscalingGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAutoscalingGroupsGenerator{}

func (x *TableAwsAutoscalingGroupsGenerator) GetTableName() string {
	return "aws_autoscaling_groups"
}

func (x *TableAwsAutoscalingGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAutoscalingGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAutoscalingGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAutoscalingGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Autoscaling
			processGroupsBundle := func(groups []types.AutoScalingGroup) error {
				input := autoscaling.DescribeNotificationConfigurationsInput{
					MaxRecords: aws.Int32(100),
				}
				for _, h := range groups {
					input.AutoScalingGroupNames = append(input.AutoScalingGroupNames, *h.AutoScalingGroupName)
				}
				var configurations []types.NotificationConfiguration
				for {
					output, err := svc.DescribeNotificationConfigurations(ctx, &input, func(o *autoscaling.Options) {
						o.Region = c.Region
					})
					if err != nil {
						return err
					}
					configurations = append(configurations, output.NotificationConfigurations...)
					if aws.ToString(output.NextToken) == "" {
						break
					}
					input.NextToken = output.NextToken
				}
				for _, gr := range groups {
					wrapper := AutoScalingGroupWrapper{
						AutoScalingGroup:		gr,
						NotificationConfigurations:	getNotificationConfigurationByGroupName(*gr.AutoScalingGroupName, configurations),
					}
					resultChannel <- wrapper
				}
				return nil
			}

			config := autoscaling.DescribeAutoScalingGroupsInput{}
			for {
				output, err := svc.DescribeAutoScalingGroups(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				groups := output.AutoScalingGroups
				for i := 0; i < len(groups); i += 255 {
					end := i + 255

					if end > len(groups) {
						end = len(groups)
					}
					t := groups[i:end]
					err := processGroupsBundle(t)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
				}

				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func getNotificationConfigurationByGroupName(name string, set []types.NotificationConfiguration) []types.NotificationConfiguration {
	var response []types.NotificationConfiguration
	for _, s := range set {
		if *s.AutoScalingGroupName == name {
			response = append(response, s)
		}
	}
	return response
}

func (x *TableAwsAutoscalingGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("autoscaling")
}

func (x *TableAwsAutoscalingGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_scaling_group_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("desired_capacity").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_check_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_instance_lifetime").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					p := result.(AutoScalingGroupWrapper)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Autoscaling
					config := autoscaling.DescribeLoadBalancersInput{AutoScalingGroupName: p.AutoScalingGroupName}
					j := map[string]interface{}{}
					for {
						output, err := svc.DescribeLoadBalancers(ctx, &config)
						if err != nil {
							if isAutoScalingGroupNotExistsError(err) {
								return nil, nil
							}
							return nil, err
						}
						for _, lb := range output.LoadBalancers {
							j[*lb.LoadBalancerName] = *lb.State
						}

						if aws.ToString(output.NextToken) == "" {
							break
						}
						config.NextToken = output.NextToken
					}
					return j, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_check_grace_period").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("warm_pool_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity_rebalance").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_instance_warmup").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled_metrics").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notification_configurations").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("new_instances_protected_from_scale_in").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("suspended_processes").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AutoScalingGroupARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zones").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_size").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_size").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("context").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mixed_instances_policy").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_zone_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VPCZoneIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_group_ar_ns").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("TargetGroupARNs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("desired_capacity_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instances").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_configuration_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_template").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("predicted_capacity").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_cooldown").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancer_names").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("placement_group").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_linked_role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceLinkedRoleARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("termination_policies").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancer_target_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					p := result.(AutoScalingGroupWrapper)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Autoscaling
					config := autoscaling.DescribeLoadBalancerTargetGroupsInput{AutoScalingGroupName: p.AutoScalingGroupName}
					j := map[string]interface{}{}
					for {
						output, err := svc.DescribeLoadBalancerTargetGroups(ctx, &config)
						if err != nil {
							if isAutoScalingGroupNotExistsError(err) {
								return nil, nil
							}
							return nil, err
						}
						for _, lb := range output.LoadBalancerTargetGroups {
							j[*lb.LoadBalancerTargetGroupARN] = *lb.State
						}

						if aws.ToString(output.NextToken) == "" {
							break
						}
						config.NextToken = output.NextToken
					}
					return j, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("warm_pool_size").ColumnType(schema.ColumnTypeInt).Build(),
	}
}

func (x *TableAwsAutoscalingGroupsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsAutoscalingGroupScalingPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsAutoscalingGroupLifecycleHooksGenerator{}),
	}
}
