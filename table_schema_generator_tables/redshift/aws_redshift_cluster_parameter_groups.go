package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRedshiftClusterParameterGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRedshiftClusterParameterGroupsGenerator{}

func (x *TableAwsRedshiftClusterParameterGroupsGenerator) GetTableName() string {
	return "aws_redshift_cluster_parameter_groups"
}

func (x *TableAwsRedshiftClusterParameterGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRedshiftClusterParameterGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRedshiftClusterParameterGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"cluster_arn",
			"parameter_group_name",
		},
	}
}

func (x *TableAwsRedshiftClusterParameterGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cluster := task.ParentRawResult.(types.Cluster)
			resultChannel <- cluster.ClusterParameterGroups
			return nil
		},
	}
}

func (x *TableAwsRedshiftClusterParameterGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("redshift")
}

func (x *TableAwsRedshiftClusterParameterGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("parameter_apply_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_redshift_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_redshift_clusters.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_arn").ColumnType(schema.ColumnTypeString).Description("`The Amazon Resource Name (ARN) for the resource.`").
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameter_group_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_parameter_status_list").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsRedshiftClusterParameterGroupsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsRedshiftClusterParametersGenerator{}),
	}
}
