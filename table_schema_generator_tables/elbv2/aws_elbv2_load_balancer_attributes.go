package elbv2

import (
	"context"

	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElbv2LoadBalancerAttributesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElbv2LoadBalancerAttributesGenerator{}

func (x *TableAwsElbv2LoadBalancerAttributesGenerator) GetTableName() string {
	return "aws_elbv2_load_balancer_attributes"
}

func (x *TableAwsElbv2LoadBalancerAttributesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElbv2LoadBalancerAttributesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElbv2LoadBalancerAttributesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsElbv2LoadBalancerAttributesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			lb := task.ParentRawResult.(types.LoadBalancer)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().ELBv2
			result, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv2.DescribeLoadBalancerAttributesInput{LoadBalancerArn: lb.LoadBalancerArn})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- result.Attributes
			return nil
		},
	}
}

func (x *TableAwsElbv2LoadBalancerAttributesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticloadbalancing")
}

func (x *TableAwsElbv2LoadBalancerAttributesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancer_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_elbv2_load_balancers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_elbv2_load_balancers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsElbv2LoadBalancerAttributesGenerator) GetSubTables() []*schema.Table {
	return nil
}
