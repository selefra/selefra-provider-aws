package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRedshiftClusterParametersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRedshiftClusterParametersGenerator{}

func (x *TableAwsRedshiftClusterParametersGenerator) GetTableName() string {
	return "aws_redshift_cluster_parameters"
}

func (x *TableAwsRedshiftClusterParametersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRedshiftClusterParametersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRedshiftClusterParametersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"cluster_arn",
			"parameter_name",
		},
	}
}

func (x *TableAwsRedshiftClusterParametersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			group := task.ParentRawResult.(types.ClusterParameterGroupStatus)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Redshift

			config := redshift.DescribeClusterParametersInput{
				ParameterGroupName: group.ParameterGroupName,
			}
			for {
				response, err := svc.DescribeClusterParameters(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Parameters
				if aws.ToString(response.Marker) == "" {
					break
				}
				config.Marker = response.Marker
			}

			return nil
		},
	}
}

func (x *TableAwsRedshiftClusterParametersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("redshift")
}

func (x *TableAwsRedshiftClusterParametersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_values").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum_engine_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameter_value").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameter_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("apply_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_modifiable").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_arn").ColumnType(schema.ColumnTypeString).Description("`The Amazon Resource Name (ARN) for the resource.`").
			Extractor(column_value_extractor.ParentColumnValue("cluster_arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_redshift_cluster_parameter_groups_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_redshift_cluster_parameter_groups.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsRedshiftClusterParametersGenerator) GetSubTables() []*schema.Table {
	return nil
}
