package apigatewayv2

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsApigatewayv2ApiIntegrationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsApigatewayv2ApiIntegrationsGenerator{}

func (x *TableAwsApigatewayv2ApiIntegrationsGenerator) GetTableName() string {
	return "aws_apigatewayv2_api_integrations"
}

func (x *TableAwsApigatewayv2ApiIntegrationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsApigatewayv2ApiIntegrationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsApigatewayv2ApiIntegrationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsApigatewayv2ApiIntegrationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.Api)
			config := apigatewayv2.GetIntegrationsInput{
				ApiId: r.ApiId,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Apigatewayv2
			for {
				response, err := svc.GetIntegrations(ctx, &config)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Items
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsApigatewayv2ApiIntegrationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("apigateway")
}

func (x *TableAwsApigatewayv2ApiIntegrationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("connection_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_selection_expression").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timeout_in_millis").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integration_response_selection_expression").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("request_parameters").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("response_parameters").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tls_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("passthrough_behavior").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integration_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integration_subtype").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integration_uri").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("payload_format_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("request_templates").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integration_method").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("integration_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_gateway_managed").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_handling_strategy").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				idsComputer := func() ([]string, error) {
					r := result.(types.Integration)
					p := task.ParentRawResult.(types.Api)
					return []string{"/apis",

						*p.ApiId, "integrations",

						*r.IntegrationId}, nil
				}

				ids, err := idsComputer()
				if err != nil {
					return nil, diagnostics.AddErrorColumnValueExtractor(task.Table, column, err)
				}

				cl := client.(*aws_client.Client)
				return arn.ARN{
					Partition:	cl.Partition,
					Service:	"apigateway",
					Region:		cl.Region,
					AccountID:	"",
					Resource:	strings.Join(ids, "/"),
				}.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_apigatewayv2_apis_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_apigatewayv2_apis.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsApigatewayv2ApiIntegrationsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsApigatewayv2ApiIntegrationResponsesGenerator{}),
	}
}
