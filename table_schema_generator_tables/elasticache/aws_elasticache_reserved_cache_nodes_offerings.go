package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticacheReservedCacheNodesOfferingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticacheReservedCacheNodesOfferingsGenerator{}

func (x *TableAwsElasticacheReservedCacheNodesOfferingsGenerator) GetTableName() string {
	return "aws_elasticache_reserved_cache_nodes_offerings"
}

func (x *TableAwsElasticacheReservedCacheNodesOfferingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticacheReservedCacheNodesOfferingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticacheReservedCacheNodesOfferingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElasticacheReservedCacheNodesOfferingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			paginator := elasticache.NewDescribeReservedCacheNodesOfferingsPaginator(client.(*aws_client.Client).AwsServices().ElastiCache, nil)
			for paginator.HasMorePages() {
				v, err := paginator.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- v.ReservedCacheNodesOfferings
			}
			return nil
		},
	}
}

func (x *TableAwsElasticacheReservedCacheNodesOfferingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticache")
}

func (x *TableAwsElasticacheReservedCacheNodesOfferingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("offering_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recurring_charges").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_node_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("duration").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fixed_price").ColumnType(schema.ColumnTypeFloat).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					item := result.(types.ReservedCacheNodesOffering)
					a := arn.ARN{
						Partition:	cl.Partition,
						Service:	"elasticache",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	"elasticache/" + aws.ToString(item.ReservedCacheNodesOfferingId),
					}
					return a.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reserved_cache_nodes_offering_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usage_price").ColumnType(schema.ColumnTypeFloat).Build(),
	}
}

func (x *TableAwsElasticacheReservedCacheNodesOfferingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
