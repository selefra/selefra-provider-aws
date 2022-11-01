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

type TableAwsGlueDatabaseTableIndexesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlueDatabaseTableIndexesGenerator{}

func (x *TableAwsGlueDatabaseTableIndexesGenerator) GetTableName() string {
	return "aws_glue_database_table_indexes"
}

func (x *TableAwsGlueDatabaseTableIndexesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlueDatabaseTableIndexesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlueDatabaseTableIndexesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"database_arn",
			"database_table_name",
			"index_name",
		},
	}
}

func (x *TableAwsGlueDatabaseTableIndexesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Glue
			d := task.ParentTask.ParentRawResult.(types.Database)
			t := task.ParentRawResult.(types.Table)
			input := glue.GetPartitionIndexesInput{DatabaseName: d.Name, CatalogId: d.CatalogId, TableName: t.Name}
			for {
				result, err := svc.GetPartitionIndexes(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.PartitionIndexDescriptorList
				if aws.ToString(result.NextToken) == "" {
					break
				}
				input.NextToken = result.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsGlueDatabaseTableIndexesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glue")
}

func (x *TableAwsGlueDatabaseTableIndexesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("index_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("index_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_table_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("keys").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backfill_errors").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_glue_database_tables_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_glue_database_tables.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("database_arn")).Build(),
	}
}

func (x *TableAwsGlueDatabaseTableIndexesGenerator) GetSubTables() []*schema.Table {
	return nil
}
