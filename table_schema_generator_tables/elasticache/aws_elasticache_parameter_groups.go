package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticacheParameterGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticacheParameterGroupsGenerator{}

func (x *TableAwsElasticacheParameterGroupsGenerator) GetTableName() string {
	return "aws_elasticache_parameter_groups"
}

func (x *TableAwsElasticacheParameterGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticacheParameterGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticacheParameterGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElasticacheParameterGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			awsProviderClient := client.(*aws_client.Client)
			svc := awsProviderClient.AwsServices().ElastiCache

			var describeCacheParameterGroupsInput elasticache.DescribeCacheParameterGroupsInput

			for {
				describeCacheParameterGroupsOutput, err := svc.DescribeCacheParameterGroups(ctx, &describeCacheParameterGroupsInput)

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- describeCacheParameterGroupsOutput.CacheParameterGroups

				if aws.ToString(describeCacheParameterGroupsOutput.Marker) == "" {
					return nil
				}

				describeCacheParameterGroupsInput.Marker = describeCacheParameterGroupsOutput.Marker
			}
		},
	}
}

func (x *TableAwsElasticacheParameterGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticache")
}

func (x *TableAwsElasticacheParameterGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_parameter_group_family").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_parameter_group_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_global").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsElasticacheParameterGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
