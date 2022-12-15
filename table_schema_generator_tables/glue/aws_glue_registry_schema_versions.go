package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsGlueRegistrySchemaVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlueRegistrySchemaVersionsGenerator{}

func (x *TableAwsGlueRegistrySchemaVersionsGenerator) GetTableName() string {
	return "aws_glue_registry_schema_versions"
}

func (x *TableAwsGlueRegistrySchemaVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlueRegistrySchemaVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlueRegistrySchemaVersionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsGlueRegistrySchemaVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			s := task.ParentRawResult.(*glue.GetSchemaOutput)
			svc := cl.AwsServices().Glue
			schemaId := types.SchemaId{
				SchemaArn: s.SchemaArn,
			}
			input := glue.ListSchemaVersionsInput{
				SchemaId:   &schemaId,
				MaxResults: aws.Int32(100),
			}
			for {
				result, err := svc.ListSchemaVersions(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, result.Schemas, func(result any) (any, error) {
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Glue
					item := result.(types.SchemaVersionListItem)

					s, err := svc.GetSchemaVersion(ctx, &glue.GetSchemaVersionInput{
						SchemaVersionId: item.SchemaVersionId,
					})
					if err != nil {
						return nil, err
					}
					return s, nil

				})
				if aws.ToString(result.NextToken) == "" {
					break
				}
				input.NextToken = result.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsGlueRegistrySchemaVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glue")
}

func (x *TableAwsGlueRegistrySchemaVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("registry_schema_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Glue
					s := result.(*glue.GetSchemaVersionOutput)
					input := &glue.QuerySchemaVersionMetadataInput{
						SchemaVersionId: s.SchemaVersionId,
					}
					metadata := make(map[string]types.MetadataInfo)
					for {
						result, err := svc.QuerySchemaVersionMetadata(ctx, input)
						if err != nil {
							if cl.IsNotFoundError(err) {
								return nil, nil
							}
							return nil, err
						}

						for k, v := range result.MetadataInfoMap {
							metadata[k] = v
						}

						if aws.ToString(result.NextToken) == "" {
							break
						}
						input.NextToken = result.NextToken
					}
					return metadata, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_glue_registry_schemas_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_glue_registry_schemas.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_definition").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_format").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema_version_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_number").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsGlueRegistrySchemaVersionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
