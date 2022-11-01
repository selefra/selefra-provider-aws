package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsGuarddutyDetectorMembersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGuarddutyDetectorMembersGenerator{}

func (x *TableAwsGuarddutyDetectorMembersGenerator) GetTableName() string {
	return "aws_guardduty_detector_members"
}

func (x *TableAwsGuarddutyDetectorMembersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGuarddutyDetectorMembersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGuarddutyDetectorMembersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsGuarddutyDetectorMembersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			detector := task.ParentRawResult.(*DetectorWrapper)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().GuardDuty
			config := &guardduty.ListMembersInput{DetectorId: aws.String(detector.Id)}
			for {
				output, err := svc.ListMembers(ctx, config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Members
				if output.NextToken == nil {
					return nil
				}
				config.NextToken = output.NextToken
			}
		},
	}
}

type DetectorWrapper struct {
	*guardduty.GetDetectorOutput
	Id	string
}

func (x *TableAwsGuarddutyDetectorMembersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("guardduty")
}

func (x *TableAwsGuarddutyDetectorMembersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("invited_at").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_guardduty_detectors_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_guardduty_detectors.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("detector_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("relationship_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("administrator_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("detector_id").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsGuarddutyDetectorMembersGenerator) GetSubTables() []*schema.Table {
	return nil
}
