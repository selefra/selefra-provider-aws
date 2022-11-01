package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsBackupGlobalSettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsBackupGlobalSettingsGenerator{}

func (x *TableAwsBackupGlobalSettingsGenerator) GetTableName() string {
	return "aws_backup_global_settings"
}

func (x *TableAwsBackupGlobalSettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsBackupGlobalSettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsBackupGlobalSettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
		},
	}
}

func (x *TableAwsBackupGlobalSettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Backup
			input := backup.DescribeGlobalSettingsInput{}

			output, err := svc.DescribeGlobalSettings(ctx, &input)
			if err != nil {
				if aws_client.IgnoreAccessDeniedServiceDisabled(err) || aws_client.IsAWSError(err, "ERROR_9601") {

					return nil
				}
				if aws_client.IsAWSError(err, "ERROR_2502") {

					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output
			return nil
		},
	}
}

func (x *TableAwsBackupGlobalSettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("backup")
}

func (x *TableAwsBackupGlobalSettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("global_settings").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_update_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsBackupGlobalSettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
