package lambda

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLambdaFunctionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLambdaFunctionsGenerator{}

func (x *TableAwsLambdaFunctionsGenerator) GetTableName() string {
	return "aws_lambda_functions"
}

func (x *TableAwsLambdaFunctionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLambdaFunctionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLambdaFunctionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLambdaFunctionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input lambda.ListFunctionsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lambda
			for {
				response, err := svc.ListFunctions(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.Functions, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Lambda
					f := result.(types.FunctionConfiguration)

					funcResponse, err := svc.GetFunction(ctx, &lambda.GetFunctionInput{
						FunctionName: f.FunctionName,
					})
					if err != nil {
						if c.IsNotFoundError(err) || c.IsAccessDeniedError(err) {
							return &lambda.GetFunctionOutput{
								Configuration: &f,
							}, nil

						}
						return nil, err
					}

					return funcResponse, nil

				})
				if aws.ToString(response.NextMarker) == "" {
					break
				}
				input.Marker = response.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsLambdaFunctionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lambda")
}

func (x *TableAwsLambdaFunctionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("concurrency").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_revision_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				c := client.(*aws_client.Client)
				svc := c.AwsServices().Lambda
				r := result.(*lambda.GetFunctionOutput)
				response, err := svc.GetPolicy(ctx, &lambda.GetPolicyInput{
					FunctionName: r.Configuration.FunctionName,
				})
				if err != nil {
					if aws_client.IsAWSError(err, "ResourceNotFoundException") {
						return nil, nil
					}
					return r, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				if response != nil {
					var policyDocument map[string]interface{}
					err = json.Unmarshal([]byte(*response.Policy), &policyDocument)
					if err != nil {
						return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
					}
					task.SetItem("policy_document", policyDocument)
					return response.RevisionId, nil
				}

				return nil, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_document").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.Wrapper("", func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				return task.GetItem("policy_document"), nil
			}, func(ctx context.Context, clientMeta *schema.ClientMeta, parentTable *schema.Table, table *schema.Table, column *schema.Column) []string {
				return []string{"policy_revision_id"}
			}, nil)).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("code_signing_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.Wrapper("", func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				c := client.(*aws_client.Client)
				svc := c.AwsServices().Lambda
				r := result.(*lambda.GetFunctionOutput)

				if row.GetStringOrDefault("code_repository_type", "") == "ECR" {
					return nil, nil
				}

				functionSigning, err := svc.GetFunctionCodeSigningConfig(ctx, &lambda.GetFunctionCodeSigningConfigInput{
					FunctionName: r.Configuration.FunctionName,
				})
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				if functionSigning.CodeSigningConfigArn == nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				signing, err := svc.GetCodeSigningConfig(ctx, &lambda.GetCodeSigningConfigInput{
					CodeSigningConfigArn: functionSigning.CodeSigningConfigArn,
				})
				if err != nil {
					if c.IsNotFoundError(err) {
						return nil, nil
					}
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				if signing.CodeSigningConfig == nil {
					return nil, nil
				}

				return signing.CodeSigningConfig, nil
			}, func(ctx context.Context, clientMeta *schema.ClientMeta, parentTable *schema.Table, table *schema.Table, column *schema.Column) []string {
				return []string{"code_repository_type"}
			}, nil)).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("code_repository_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Code.RepositoryType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Configuration.FunctionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("code").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsLambdaFunctionsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsLambdaFunctionAliasesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsLambdaFunctionVersionsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsLambdaFunctionConcurrencyConfigsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsLambdaFunctionEventSourceMappingsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsLambdaFunctionEventInvokeConfigsGenerator{}),
	}
}
