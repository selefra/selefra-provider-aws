package backup

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsBackupVaultsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsBackupVaultsGenerator{}

func (x *TableAwsBackupVaultsGenerator) GetTableName() string {
	return "aws_backup_vaults"
}

func (x *TableAwsBackupVaultsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsBackupVaultsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsBackupVaultsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsBackupVaultsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Backup
			params := backup.ListBackupVaultsInput{MaxResults: aws.Int32(1000)}
			for {
				result, err := svc.ListBackupVaults(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.BackupVaultList
				if aws.ToString(result.NextToken) == "" {
					break
				}
				params.NextToken = result.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsBackupVaultsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("backup")
}

func (x *TableAwsBackupVaultsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("notifications").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					vault := result.(types.BackupVaultListMember)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Backup
					result, err := svc.GetBackupVaultNotifications(
						ctx,
						&backup.GetBackupVaultNotificationsInput{BackupVaultName: vault.BackupVaultName},
						func(o *backup.Options) {
							o.Region = cl.Region
						},
					)
					if err != nil {
						return nil, err
					}
					return result, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_vault_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("number_of_recovery_points").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locked").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_retention_days").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					vault := result.(types.BackupVaultListMember)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().Backup
					result, err := svc.GetBackupVaultAccessPolicy(
						ctx,
						&backup.GetBackupVaultAccessPolicyInput{BackupVaultName: vault.BackupVaultName},
						func(o *backup.Options) {
							o.Region = cl.Region
						},
					)
					if err != nil {
						if cl.IsNotFoundError(err) {
							return nil, nil
						}
						return nil, err
					}
					if result.Policy == nil {
						return nil, nil
					}

					var p map[string]interface{}
					err = json.Unmarshal([]byte(*result.Policy), &p)
					if err != nil {
						return nil, err
					}
					return p, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_key_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lock_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BackupVaultArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator_request_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_retention_days").ColumnType(schema.ColumnTypeInt).Build(),
	}
}

func (x *TableAwsBackupVaultsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsBackupVaultRecoveryPointsGenerator{}),
	}
}
