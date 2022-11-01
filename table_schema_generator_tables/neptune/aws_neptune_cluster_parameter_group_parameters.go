package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsNeptuneClusterParameterGroupParametersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsNeptuneClusterParameterGroupParametersGenerator{}

func (x *TableAwsNeptuneClusterParameterGroupParametersGenerator) GetTableName() string {
	return "aws_neptune_cluster_parameter_group_parameters"
}

func (x *TableAwsNeptuneClusterParameterGroupParametersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsNeptuneClusterParameterGroupParametersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsNeptuneClusterParameterGroupParametersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsNeptuneClusterParameterGroupParametersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Neptune
			g := task.ParentRawResult.(types.DBClusterParameterGroup)
			input := neptune.DescribeDBClusterParametersInput{DBClusterParameterGroupName: g.DBClusterParameterGroupName}
			for {
				output, err := svc.DescribeDBClusterParameters(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Parameters
				if aws.ToString(output.Marker) == "" {
					break
				}
				input.Marker = output.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsNeptuneClusterParameterGroupParametersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("neptune")
}

func (x *TableAwsNeptuneClusterParameterGroupParametersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("parameter_value").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("apply_method").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_values").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_parameter_group_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("apply_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_modifiable").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum_engine_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameter_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_neptune_cluster_parameter_groups_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_neptune_cluster_parameter_groups.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsNeptuneClusterParameterGroupParametersGenerator) GetSubTables() []*schema.Table {
	return nil
}
