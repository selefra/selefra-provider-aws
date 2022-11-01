package docdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDocdbClusterSnapshotsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDocdbClusterSnapshotsGenerator{}

func (x *TableAwsDocdbClusterSnapshotsGenerator) GetTableName() string {
	return "aws_docdb_cluster_snapshots"
}

func (x *TableAwsDocdbClusterSnapshotsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDocdbClusterSnapshotsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDocdbClusterSnapshotsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsDocdbClusterSnapshotsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			item := task.ParentRawResult.(types.DBCluster)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().DocDB

			input := &docdb.DescribeDBClusterSnapshotsInput{
				DBClusterIdentifier: item.DBClusterIdentifier,
			}
			p := docdb.NewDescribeDBClusterSnapshotsPaginator(svc, input)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.DBClusterSnapshots
			}
			return nil
		},
	}
}

func (x *TableAwsDocdbClusterSnapshotsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("docdb")
}

func (x *TableAwsDocdbClusterSnapshotsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("attributes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					item := result.(types.DBClusterSnapshot)
					cli := client.(*aws_client.Client)
					svc := cli.AwsServices().DocDB

					input := &docdb.DescribeDBClusterSnapshotAttributesInput{
						DBClusterSnapshotIdentifier: item.DBClusterSnapshotIdentifier,
					}

					output, err := svc.DescribeDBClusterSnapshotAttributes(ctx, input)
					if err != nil {
						return nil, err
					}
					return output.DBClusterSnapshotAttributesResult.DBClusterSnapshotAttributes, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zones").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_db_cluster_snapshot_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceDBClusterSnapshotArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_encrypted").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterSnapshotArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("percent_progress").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_create_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_username").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_create_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_cluster_snapshot_identifier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBClusterSnapshotIdentifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_docdb_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_docdb_clusters.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsDocdbClusterSnapshotsGenerator) GetSubTables() []*schema.Table {
	return nil
}
