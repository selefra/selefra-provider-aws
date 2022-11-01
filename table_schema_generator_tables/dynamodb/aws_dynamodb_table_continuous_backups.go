package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDynamodbTableContinuousBackupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDynamodbTableContinuousBackupsGenerator{}

func (x *TableAwsDynamodbTableContinuousBackupsGenerator) GetTableName() string {
	return "aws_dynamodb_table_continuous_backups"
}

func (x *TableAwsDynamodbTableContinuousBackupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDynamodbTableContinuousBackupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDynamodbTableContinuousBackupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsDynamodbTableContinuousBackupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			par := task.ParentRawResult.(*types.TableDescription)

			c := client.(*aws_client.Client)
			svc := c.AwsServices().DynamoDB

			output, err := svc.DescribeContinuousBackups(ctx, &dynamodb.DescribeContinuousBackupsInput{
				TableName: par.TableName,
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			resultChannel <- output.ContinuousBackupsDescription
			return nil
		},
	}
}

func (x *TableAwsDynamodbTableContinuousBackupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("dynamodb")
}

func (x *TableAwsDynamodbTableContinuousBackupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("aws_dynamodb_tables_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_dynamodb_tables.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("table_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("continuous_backups_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("point_in_time_recovery_description").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsDynamodbTableContinuousBackupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
