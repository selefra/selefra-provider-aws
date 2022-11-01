package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApigatewayRestApiDocumentationVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayRestApiDocumentationVersionsGenerator{}

func (x *TableAwsApigatewayRestApiDocumentationVersionsGenerator) GetTableName() string {
	return "aws_apigateway_rest_api_documentation_versions"
}

func (x *TableAwsApigatewayRestApiDocumentationVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayRestApiDocumentationVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayRestApiDocumentationVersionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsApigatewayRestApiDocumentationVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.RestApi)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigateway
			config := apigateway.GetDocumentationVersionsInput{RestApiId: r.Id}
			for {
				response, err := svc.GetDocumentationVersions(ctx, &config)
				if err != nil {
					if c.IsNotFoundError(err) {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Items
				if aws.ToString(response.Position) == "" {
					break
				}
				config.Position = response.Position
			}
			return nil
		},
	}
}

func (x *TableAwsApigatewayRestApiDocumentationVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayRestApiDocumentationVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					v := result.(types.DocumentationVersion)
					rapi := task.ParentRawResult.(types.RestApi)
					arn := cl.RegionGlobalARN("apigateway", "/restapis",

						*rapi.Id, "documentation/versions", *v.Version)
					return arn, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigateway_rest_apis_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigateway_rest_apis.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rest_api_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsApigatewayRestApiDocumentationVersionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
