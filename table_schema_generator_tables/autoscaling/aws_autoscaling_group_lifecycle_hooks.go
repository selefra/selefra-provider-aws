package autoscaling

import (
	"context"
	"errors"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAutoscalingGroupLifecycleHooksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAutoscalingGroupLifecycleHooksGenerator{}

func (x *TableAwsAutoscalingGroupLifecycleHooksGenerator) GetTableName() string {
	return "aws_autoscaling_group_lifecycle_hooks"
}

func (x *TableAwsAutoscalingGroupLifecycleHooksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAutoscalingGroupLifecycleHooksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAutoscalingGroupLifecycleHooksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsAutoscalingGroupLifecycleHooksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(AutoScalingGroupWrapper)
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Autoscaling
			config := autoscaling.DescribeLifecycleHooksInput{AutoScalingGroupName: p.AutoScalingGroupName}

			output, err := svc.DescribeLifecycleHooks(ctx, &config)
			if err != nil {
				if isAutoScalingGroupNotExistsError(err) {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output.LifecycleHooks
			return nil
		},
	}
}

var groupNotFoundRegex = regexp.MustCompile(`AutoScalingGroup name not found|Group .* not found`)

type AutoScalingGroupWrapper struct {
	types.AutoScalingGroup
	NotificationConfigurations []types.NotificationConfiguration
}

func isAutoScalingGroupNotExistsError(err error) bool {
	var ae smithy.APIError
	if errors.As(err, &ae) {
		if ae.ErrorCode() == "ValidationError" && groupNotFoundRegex.MatchString(ae.ErrorMessage()) {
			return true
		}
	}
	return false
}

func (x *TableAwsAutoscalingGroupLifecycleHooksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("autoscaling")
}

func (x *TableAwsAutoscalingGroupLifecycleHooksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_hook_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notification_target_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NotificationTargetARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_autoscaling_groups_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_autoscaling_groups.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("heartbeat_timeout").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notification_metadata").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_scaling_group_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_result").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RoleARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("global_timeout").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_transition").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsAutoscalingGroupLifecycleHooksGenerator) GetSubTables() []*schema.Table {
	return nil
}
