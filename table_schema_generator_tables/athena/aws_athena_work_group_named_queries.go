package athena

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAthenaWorkGroupNamedQueriesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAthenaWorkGroupNamedQueriesGenerator{}

func (x *TableAwsAthenaWorkGroupNamedQueriesGenerator) GetTableName() string {
	return "aws_athena_work_group_named_queries"
}

func (x *TableAwsAthenaWorkGroupNamedQueriesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAthenaWorkGroupNamedQueriesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAthenaWorkGroupNamedQueriesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsAthenaWorkGroupNamedQueriesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Athena
			wg := task.ParentRawResult.(types.WorkGroup)
			input := athena.ListNamedQueriesInput{WorkGroup: wg.Name}
			for {
				response, err := svc.ListNamedQueries(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, d := range response.NamedQueryIds {
					dc, err := svc.GetNamedQuery(ctx, &athena.GetNamedQueryInput{
						NamedQueryId: aws.String(d),
					})
					if err != nil {
						if c.IsNotFoundError(err) {
							continue
						}
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- *dc.NamedQuery
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

func (x *TableAwsAthenaWorkGroupNamedQueriesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("athena")
}

func (x *TableAwsAthenaWorkGroupNamedQueriesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_athena_work_groups_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_athena_work_groups.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("work_group_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_string").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("named_query_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("work_group").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsAthenaWorkGroupNamedQueriesGenerator) GetSubTables() []*schema.Table {
	return nil
}
