package elbv1

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElbv1LoadBalancersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElbv1LoadBalancersGenerator{}

func (x *TableAwsElbv1LoadBalancersGenerator) GetTableName() string {
	return "aws_elbv1_load_balancers"
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().ELBv1
			processLoadBalancers := func(loadBalancers []types.LoadBalancerDescription) error {
				tagsCfg := &elbv1.DescribeTagsInput{LoadBalancerNames: make([]string, 0, len(loadBalancers))}
				for _, lb := range loadBalancers {
					tagsCfg.LoadBalancerNames = append(tagsCfg.LoadBalancerNames, *lb.LoadBalancerName)
				}
				tagsResponse, err := svc.DescribeTags(ctx, tagsCfg)
				if err != nil {
					return err
				}
				for _, lb := range loadBalancers {
					loadBalancerAttributes, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv1.DescribeLoadBalancerAttributesInput{LoadBalancerName: lb.LoadBalancerName})
					if err != nil {
						if c.IsNotFoundError(err) {
							continue
						}
						return err
					}

					wrapper := ELBv1LoadBalancerWrapper{
						LoadBalancerDescription:	lb,
						Tags:				aws_client.TagsToMap(getTagsByLoadBalancerName(*lb.LoadBalancerName, tagsResponse.TagDescriptions)),
						Attributes:			loadBalancerAttributes.LoadBalancerAttributes,
					}

					resultChannel <- wrapper
				}
				return nil
			}

			var config elbv1.DescribeLoadBalancersInput
			for {
				response, err := svc.DescribeLoadBalancers(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				for i := 0; i < len(response.LoadBalancerDescriptions); i += 20 {
					end := i + 20

					if end > len(response.LoadBalancerDescriptions) {
						end = len(response.LoadBalancerDescriptions)
					}
					loadBalancers := response.LoadBalancerDescriptions[i:end]
					if err := processLoadBalancers(loadBalancers); err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
				}

				if aws.ToString(response.NextMarker) == "" {
					break
				}
				config.Marker = response.NextMarker
			}

			return nil
		},
	}
}

func getTagsByLoadBalancerName(id string, tagsResponse []types.TagDescription) []types.Tag {
	for _, t := range tagsResponse {
		if id == *t.LoadBalancerName {
			return t.Tags
		}
	}
	return nil
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticloadbalancing")
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("canonical_hosted_zone_name_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CanonicalHostedZoneNameID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instances").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("listener_descriptions").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zones").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DNSName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancer_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attributes").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backend_server_descriptions").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policies").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scheme").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_groups").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				idsComputer := func() ([]string, error) {
					return []string{"loadbalancer", *result.(ELBv1LoadBalancerWrapper).LoadBalancerName}, nil
				}

				ids, err := idsComputer()
				if err != nil {
					return nil, diagnostics.AddErrorColumnValueExtractor(task.Table, column, err)
				}

				cl := client.(*aws_client.Client)
				return arn.ARN{
					Partition:	cl.Partition,
					Service:	"elasticloadbalancing",
					Region:		cl.Region,
					AccountID:	cl.AccountID,
					Resource:	strings.Join(ids, "/"),
				}.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("canonical_hosted_zone_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_check").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_security_group").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnets").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VPCId")).Build(),
	}
}

func (x *TableAwsElbv1LoadBalancersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsElbv1LoadBalancerPoliciesGenerator{}),
	}
}
