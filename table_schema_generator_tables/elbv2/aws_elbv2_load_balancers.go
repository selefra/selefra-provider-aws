package elbv2

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElbv2LoadBalancersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElbv2LoadBalancersGenerator{}

func (x *TableAwsElbv2LoadBalancersGenerator) GetTableName() string {
	return "aws_elbv2_load_balancers"
}

func (x *TableAwsElbv2LoadBalancersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElbv2LoadBalancersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElbv2LoadBalancersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElbv2LoadBalancersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config elbv2.DescribeLoadBalancersInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().ELBv2
			for {
				response, err := svc.DescribeLoadBalancers(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.LoadBalancers
				if aws.ToString(response.NextMarker) == "" {
					break
				}
				config.Marker = response.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsElbv2LoadBalancersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticloadbalancing")
}

func (x *TableAwsElbv2LoadBalancersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("web_acl_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					p := result.(types.LoadBalancer)

					if p.Type != types.LoadBalancerTypeEnumApplication {
						return nil, nil
					}
					cl := client.(*aws_client.Client).AwsServices().WafV2
					input := wafv2.GetWebACLForResourceInput{ResourceArn: p.LoadBalancerArn}
					response, err := cl.GetWebACLForResource(ctx, &input, func(options *wafv2.Options) {})
					if err != nil {
						var exc *wafv2types.WAFNonexistentItemException
						if errors.As(err, &exc) {
							if exc.ErrorCode() == "WAFNonexistentItemException" {
								return nil, nil
							}
						}

						return nil, err
					}
					if response.WebACL == nil {
						return nil, nil
					}

					return response.WebACL.ARN, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_owned_ipv4_pool").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LoadBalancerArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scheme").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("canonical_hosted_zone_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_groups").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zones").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DNSName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancer_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsElbv2LoadBalancersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsElbv2ListenersGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsElbv2LoadBalancerAttributesGenerator{}),
	}
}
