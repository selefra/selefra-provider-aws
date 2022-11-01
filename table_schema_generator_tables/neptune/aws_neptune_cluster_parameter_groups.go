package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsNeptuneClusterParameterGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsNeptuneClusterParameterGroupsGenerator{}

func (x *TableAwsNeptuneClusterParameterGroupsGenerator) GetTableName() string {
	return "aws_neptune_cluster_parameter_groups"
}

func (x *TableAwsNeptuneClusterParameterGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsNeptuneClusterParameterGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsNeptuneClusterParameterGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsNeptuneClusterParameterGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Neptune
			var input neptune.DescribeDBClusterParameterGroupsInput

			for {
				output, err := svc.DescribeDBClusterParameterGroups(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.DBClusterParameterGroups
				if aws.ToString(output.Marker) == "" {
					break
				}
				input.Marker = output.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsNeptuneClusterParameterGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("neptune")
}

func (x *TableAwsNeptuneClusterParameterGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("db_parameter_group_family").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBParameterGroupFamily")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterParameterGroupArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_parameter_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterParameterGroupName")).Build(),
	}
}

func (x *TableAwsNeptuneClusterParameterGroupsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsNeptuneClusterParameterGroupParametersGenerator{}),
	}
}
