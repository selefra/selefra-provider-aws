package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLambdaFunctionEventSourceMappingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLambdaFunctionEventSourceMappingsGenerator{}

func (x *TableAwsLambdaFunctionEventSourceMappingsGenerator) GetTableName() string {
	return "aws_lambda_function_event_source_mappings"
}

func (x *TableAwsLambdaFunctionEventSourceMappingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLambdaFunctionEventSourceMappingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLambdaFunctionEventSourceMappingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLambdaFunctionEventSourceMappingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(*lambda.GetFunctionOutput)
			if p.Configuration == nil {
				return nil
			}

			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Lambda
			config := lambda.ListEventSourceMappingsInput{
				FunctionName: p.Configuration.FunctionName,
			}

			for {
				output, err := svc.ListEventSourceMappings(ctx, &config)
				if err != nil {
					if cl.IsNotFoundError(err) {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.EventSourceMappings
				if output.NextMarker == nil {
					break
				}
				config.Marker = output.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsLambdaFunctionEventSourceMappingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lambda")
}

func (x *TableAwsLambdaFunctionEventSourceMappingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("function_response_types").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_managed_event_source").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_lambda_functions_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_lambda_functions.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("batch_size").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("destination_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_managed_kafka_event_source_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bisect_batch_on_function_error").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum_retry_attempts").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("queues").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("starting_position").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("starting_position_timestamp").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("amazon_managed_kafka_event_source_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("topics").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filter_criteria").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_processing_result").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parallelization_factor").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tumbling_window_in_seconds").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uuid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("function_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum_record_age_in_seconds").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_source_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum_batching_window_in_seconds").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_access_configurations").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_transition_reason").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsLambdaFunctionEventSourceMappingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
