package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsBackupPlanSelectionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsBackupPlanSelectionsGenerator{}

func (x *TableAwsBackupPlanSelectionsGenerator) GetTableName() string {
	return "aws_backup_plan_selections"
}

func (x *TableAwsBackupPlanSelectionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsBackupPlanSelectionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsBackupPlanSelectionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsBackupPlanSelectionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			plan := task.ParentRawResult.(*backup.GetBackupPlanOutput)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Backup
			params := backup.ListBackupSelectionsInput{
				BackupPlanId:	plan.BackupPlanId,
				MaxResults:	aws.Int32(1000),
			}
			for {
				result, err := svc.ListBackupSelections(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, m := range result.BackupSelectionsList {
					s, err := svc.GetBackupSelection(
						ctx,
						&backup.GetBackupSelectionInput{BackupPlanId: plan.BackupPlanId, SelectionId: m.SelectionId},
						func(o *backup.Options) {
							o.Region = cl.Region
						},
					)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					if s != nil {
						resultChannel <- *s
					}
				}
				if aws.ToString(result.NextToken) == "" {
					break
				}
				params.NextToken = result.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsBackupPlanSelectionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("backup")
}

func (x *TableAwsBackupPlanSelectionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("plan_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_plan_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator_request_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selection_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_backup_plans_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_backup_plans.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_selection").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsBackupPlanSelectionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
