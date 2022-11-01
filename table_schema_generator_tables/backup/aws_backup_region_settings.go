package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsBackupRegionSettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsBackupRegionSettingsGenerator{}

func (x *TableAwsBackupRegionSettingsGenerator) GetTableName() string {
	return "aws_backup_region_settings"
}

func (x *TableAwsBackupRegionSettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsBackupRegionSettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsBackupRegionSettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
		},
	}
}

func (x *TableAwsBackupRegionSettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Backup
			input := backup.DescribeRegionSettingsInput{}

			output, err := svc.DescribeRegionSettings(ctx, &input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output
			return nil
		},
	}
}

func (x *TableAwsBackupRegionSettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("backup")
}

func (x *TableAwsBackupRegionSettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type_management_preference").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type_opt_in_preference").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsBackupRegionSettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
