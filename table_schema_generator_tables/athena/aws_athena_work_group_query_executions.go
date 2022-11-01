package athena

import (
	"context"
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAthenaWorkGroupQueryExecutionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAthenaWorkGroupQueryExecutionsGenerator{}

func (x *TableAwsAthenaWorkGroupQueryExecutionsGenerator) GetTableName() string {
	return "aws_athena_work_group_query_executions"
}

func (x *TableAwsAthenaWorkGroupQueryExecutionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAthenaWorkGroupQueryExecutionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAthenaWorkGroupQueryExecutionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsAthenaWorkGroupQueryExecutionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Athena
			wg := task.ParentRawResult.(types.WorkGroup)
			input := athena.ListQueryExecutionsInput{WorkGroup: wg.Name}
			for {
				response, err := svc.ListQueryExecutions(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, d := range response.QueryExecutionIds {
					dc, err := svc.GetQueryExecution(ctx, &athena.GetQueryExecutionInput{
						QueryExecutionId: aws.String(d),
					})
					if err != nil {
						if c.IsNotFoundError(err) || isQueryExecutionNotFound(err) {
							continue
						}
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- *dc.QueryExecution
					return nil
				}
				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func isQueryExecutionNotFound(err error) bool {
	var ae smithy.APIError
	if !errors.As(err, &ae) {
		return false
	}
	return ae.ErrorCode() == "InvalidRequestException" && strings.Contains(ae.ErrorMessage(), "was not found")
}

func (x *TableAwsAthenaWorkGroupQueryExecutionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("athena")
}

func (x *TableAwsAthenaWorkGroupQueryExecutionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("work_group_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("execution_parameters").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_execution_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("statistics").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("work_group").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_execution_context").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("statement_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_athena_work_groups_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_athena_work_groups.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsAthenaWorkGroupQueryExecutionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
