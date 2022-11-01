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

type TableAwsElbv2TargetGroupTargetHealthDescriptionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElbv2TargetGroupTargetHealthDescriptionsGenerator{}

func (x *TableAwsElbv2TargetGroupTargetHealthDescriptionsGenerator) GetTableName() string {
	return "aws_elbv2_target_group_target_health_descriptions"
}

func (x *TableAwsElbv2TargetGroupTargetHealthDescriptionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElbv2TargetGroupTargetHealthDescriptionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElbv2TargetGroupTargetHealthDescriptionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsElbv2TargetGroupTargetHealthDescriptionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().ELBv2
			tg := task.ParentRawResult.(types.TargetGroup)
			response, err := svc.DescribeTargetHealth(ctx, &elbv2.DescribeTargetHealthInput{
				TargetGroupArn: tg.TargetGroupArn,
			})
			if err != nil {
				if cl.IsNotFoundError(err) {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response.TargetHealthDescriptions
			return nil
		},
	}
}

func (x *TableAwsElbv2TargetGroupTargetHealthDescriptionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticloadbalancing")
}

func (x *TableAwsElbv2TargetGroupTargetHealthDescriptionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("target_group_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_check_port").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_health").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_elbv2_target_groups_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_elbv2_target_groups.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsElbv2TargetGroupTargetHealthDescriptionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
