package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAthenaDataCatalogDatabaseTablesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAthenaDataCatalogDatabaseTablesGenerator{}

func (x *TableAwsAthenaDataCatalogDatabaseTablesGenerator) GetTableName() string {
	return "aws_athena_data_catalog_database_tables"
}

func (x *TableAwsAthenaDataCatalogDatabaseTablesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAthenaDataCatalogDatabaseTablesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAthenaDataCatalogDatabaseTablesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"data_catalog_arn",
			"data_catalog_database_name",
			"name",
		},
	}
}

func (x *TableAwsAthenaDataCatalogDatabaseTablesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Athena
			input := athena.ListTableMetadataInput{
				CatalogName:	task.ParentTask.ParentRawResult.(types.DataCatalog).Name,
				DatabaseName:	task.ParentRawResult.(types.Database).Name,
			}
			for {
				response, err := svc.ListTableMetadata(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.TableMetadataList

				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsAthenaDataCatalogDatabaseTablesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("athena")
}

func (x *TableAwsAthenaDataCatalogDatabaseTablesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_access_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameters").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("partition_keys").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_athena_data_catalog_databases_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_athena_data_catalog_databases.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("columns").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("table_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_catalog_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("data_catalog_arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_catalog_database_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("name")).Build(),
	}
}

func (x *TableAwsAthenaDataCatalogDatabaseTablesGenerator) GetSubTables() []*schema.Table {
	return nil
}
