package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAutoscalingGroupScalingPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAutoscalingGroupScalingPoliciesGenerator{}

func (x *TableAwsAutoscalingGroupScalingPoliciesGenerator) GetTableName() string {
	return "aws_autoscaling_group_scaling_policies"
}

func (x *TableAwsAutoscalingGroupScalingPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAutoscalingGroupScalingPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAutoscalingGroupScalingPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAutoscalingGroupScalingPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(AutoScalingGroupWrapper)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Autoscaling
			config := autoscaling.DescribePoliciesInput{AutoScalingGroupName: p.AutoScalingGroupName}

			for {
				output, err := svc.DescribePolicies(ctx, &config)
				if err != nil {
					if isAutoScalingGroupNotExistsError(err) {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.ScalingPolicies

				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsAutoscalingGroupScalingPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("autoscaling")
}

func (x *TableAwsAutoscalingGroupScalingPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("cooldown").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_tracking_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_scaling_group_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("adjustment_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("estimated_instance_warmup").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_adjustment_magnitude").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scaling_adjustment").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alarms").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_aggregation_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("predictive_scaling_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PolicyARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_adjustment_step").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("step_adjustments").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_autoscaling_groups_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_autoscaling_groups.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsAutoscalingGroupScalingPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
