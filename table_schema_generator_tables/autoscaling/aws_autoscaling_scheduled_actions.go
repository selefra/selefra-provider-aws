package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAutoscalingScheduledActionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAutoscalingScheduledActionsGenerator{}

func (x *TableAwsAutoscalingScheduledActionsGenerator) GetTableName() string {
	return "aws_autoscaling_scheduled_actions"
}

func (x *TableAwsAutoscalingScheduledActionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAutoscalingScheduledActionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAutoscalingScheduledActionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAutoscalingScheduledActionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Autoscaling
			params := &autoscaling.DescribeScheduledActionsInput{
				MaxRecords: aws.Int32(100),
			}
			for {
				output, err := svc.DescribeScheduledActions(ctx, params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, scheduledUpdateGroupAction := range output.ScheduledUpdateGroupActions {
					resultChannel <- scheduledUpdateGroupAction
				}
				if aws.ToString(output.NextToken) == "" {
					break
				}
				params.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsAutoscalingScheduledActionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("autoscaling")
}

func (x *TableAwsAutoscalingScheduledActionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("desired_capacity").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_scaling_group_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_size").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recurrence").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_zone").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ScheduledActionARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_size").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scheduled_action_name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsAutoscalingScheduledActionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
