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

type TableAwsApigatewayDomainNameBasePathMappingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayDomainNameBasePathMappingsGenerator{}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetTableName() string {
	return "aws_apigateway_domain_name_base_path_mappings"
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.DomainName)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigateway
			config := apigateway.GetBasePathMappingsInput{DomainName: r.DomainName}
			for p := apigateway.NewGetBasePathMappingsPaginator(svc, &config); p.HasMorePages(); {
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

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("base_path").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rest_api_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stage").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					domain := task.ParentRawResult.(types.DomainName)
					mapping := result.(types.BasePathMapping)
					arn := cl.RegionGlobalARN("apigateway", "/domainnames",

						*domain.DomainName, "basepathmappings", *mapping.BasePath)
					return arn, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigateway_domain_names_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigateway_domain_names.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsApigatewayDomainNameBasePathMappingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
