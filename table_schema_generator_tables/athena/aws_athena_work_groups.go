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

type TableAwsAthenaWorkGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAthenaWorkGroupsGenerator{}

func (x *TableAwsAthenaWorkGroupsGenerator) GetTableName() string {
	return "aws_athena_work_groups"
}

func (x *TableAwsAthenaWorkGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAthenaWorkGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAthenaWorkGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAthenaWorkGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Athena
			input := athena.ListWorkGroupsInput{}
			for {
				response, err := svc.ListWorkGroups(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.WorkGroups, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Athena

					wg := result.(types.WorkGroupSummary)
					dc, err := svc.GetWorkGroup(ctx, &athena.GetWorkGroupInput{
						WorkGroup: wg.Name,
					})
					if err != nil {
						return nil, err
					}
					return *dc.WorkGroup, nil

				})
				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}

			return nil
		},
	}
}

func (x *TableAwsAthenaWorkGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("athena")
}

func (x *TableAwsAthenaWorkGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				createWorkGroupArn := func(cl *aws_client.Client, groupName string) string {
					return cl.ARN("athena", "workgroup", groupName)
				}
				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					dc := result.(types.WorkGroup)
					return createWorkGroupArn(cl, *dc.Name), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsAthenaWorkGroupsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsAthenaWorkGroupNamedQueriesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsAthenaWorkGroupPreparedStatementsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsAthenaWorkGroupQueryExecutionsGenerator{}),
	}
}
