package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDynamodbTableReplicaAutoScalingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDynamodbTableReplicaAutoScalingsGenerator{}

func (x *TableAwsDynamodbTableReplicaAutoScalingsGenerator) GetTableName() string {
	return "aws_dynamodb_table_replica_auto_scalings"
}

func (x *TableAwsDynamodbTableReplicaAutoScalingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDynamodbTableReplicaAutoScalingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDynamodbTableReplicaAutoScalingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsDynamodbTableReplicaAutoScalingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			par := task.ParentRawResult.(*types.TableDescription)

			if aws.ToString(par.GlobalTableVersion) == "" {

				return nil
			}

			c := client.(*aws_client.Client)
			svc := c.AwsServices().DynamoDB

			output, err := svc.DescribeTableReplicaAutoScaling(ctx, &dynamodb.DescribeTableReplicaAutoScalingInput{
				TableName: par.TableName,
			})
			if err != nil {
				if c.IsNotFoundError(err) {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			for i := range output.TableAutoScalingDescription.Replicas {
				resultChannel <- output.TableAutoScalingDescription.Replicas[i]
			}
			return nil
		},
	}
}

func (x *TableAwsDynamodbTableReplicaAutoScalingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("dynamodb")
}

func (x *TableAwsDynamodbTableReplicaAutoScalingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("table_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("global_secondary_indexes").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replica_provisioned_read_capacity_auto_scaling_settings").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replica_provisioned_write_capacity_auto_scaling_settings").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replica_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_dynamodb_tables_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_dynamodb_tables.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsDynamodbTableReplicaAutoScalingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
