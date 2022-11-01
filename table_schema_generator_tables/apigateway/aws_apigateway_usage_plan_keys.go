package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApigatewayUsagePlanKeysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayUsagePlanKeysGenerator{}

func (x *TableAwsApigatewayUsagePlanKeysGenerator) GetTableName() string {
	return "aws_apigateway_usage_plan_keys"
}

func (x *TableAwsApigatewayUsagePlanKeysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayUsagePlanKeysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayUsagePlanKeysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsApigatewayUsagePlanKeysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.UsagePlan)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigateway
			config := apigateway.GetUsagePlanKeysInput{UsagePlanId: r.Id}
			for p := apigateway.NewGetUsagePlanKeysPaginator(svc, &config); p.HasMorePages(); {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Items
			}
			return nil
		},
	}
}

func (x *TableAwsApigatewayUsagePlanKeysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayUsagePlanKeysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigateway_usage_plans_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigateway_usage_plans.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usage_plan_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					up := task.ParentRawResult.(types.UsagePlan)
					key := result.(types.UsagePlanKey)
					arn := cl.RegionGlobalARN("apigateway", "/usageplans",

						*up.Id, "keys", *key.Id)
					return arn, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsApigatewayUsagePlanKeysGenerator) GetSubTables() []*schema.Table {
	return nil
}
