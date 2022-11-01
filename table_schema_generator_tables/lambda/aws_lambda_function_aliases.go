package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLambdaFunctionAliasesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLambdaFunctionAliasesGenerator{}

func (x *TableAwsLambdaFunctionAliasesGenerator) GetTableName() string {
	return "aws_lambda_function_aliases"
}

func (x *TableAwsLambdaFunctionAliasesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLambdaFunctionAliasesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLambdaFunctionAliasesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsLambdaFunctionAliasesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(*lambda.GetFunctionOutput)
			if p.Configuration == nil {
				return nil
			}

			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lambda
			config := lambda.ListAliasesInput{
				FunctionName: p.Configuration.FunctionName,
			}

			for {
				output, err := svc.ListAliases(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				if err != nil {
					if c.IsNotFoundError(err) {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.Aliases, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Lambda
					alias := result.(types.AliasConfiguration)
					p := task.ParentRawResult.(*lambda.GetFunctionOutput)

					urlConfig, err := svc.GetFunctionUrlConfig(ctx, &lambda.GetFunctionUrlConfigInput{
						FunctionName:	p.Configuration.FunctionName,
						Qualifier:	alias.Name,
					})
					if err != nil && !c.IsNotFoundError(err) {
						return nil, err
					}
					return &AliasWrapper{AliasConfiguration: &alias, UrlConfig: urlConfig}, nil

				})
				if output.NextMarker == nil {
					break
				}
				config.Marker = output.NextMarker
			}
			return nil
		},
	}
}

type AliasWrapper struct {
	*types.AliasConfiguration
	UrlConfig	*lambda.GetFunctionUrlConfigOutput
}

func (x *TableAwsLambdaFunctionAliasesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lambda")
}

func (x *TableAwsLambdaFunctionAliasesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("routing_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AliasArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("revision_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("function_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_lambda_functions_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_lambda_functions.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("function_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsLambdaFunctionAliasesGenerator) GetSubTables() []*schema.Table {
	return nil
}
