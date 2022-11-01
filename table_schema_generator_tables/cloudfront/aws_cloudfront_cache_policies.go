package cloudfront

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCloudfrontCachePoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudfrontCachePoliciesGenerator{}

func (x *TableAwsCloudfrontCachePoliciesGenerator) GetTableName() string {
	return "aws_cloudfront_cache_policies"
}

func (x *TableAwsCloudfrontCachePoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudfrontCachePoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudfrontCachePoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsCloudfrontCachePoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config cloudfront.ListCachePoliciesInput
			c := client.(*aws_client.Client)
			s := c.AwsServices()
			svc := s.Cloudfront
			for {
				response, err := svc.ListCachePolicies(ctx, nil)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				if response.CachePolicyList != nil {
					resultChannel <- response.CachePolicyList.Items
				}

				if aws.ToString(response.CachePolicyList.NextMarker) == "" {
					break
				}
				config.Marker = response.CachePolicyList.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsCloudfrontCachePoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsCloudfrontCachePoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				idsComputer := func() ([]string, error) {
					return []string{"cache-policy", *result.(types.CachePolicySummary).CachePolicy.Id}, nil
				}

				ids, err := idsComputer()
				if err != nil {
					return nil, diagnostics.AddErrorColumnValueExtractor(task.Table, column, err)
				}

				cl := client.(*aws_client.Client)
				return arn.ARN{
					Partition:	cl.Partition,
					Service:	"cloudfront",
					Region:		"",
					AccountID:	cl.AccountID,
					Resource:	strings.Join(ids, "/"),
				}.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_policy").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsCloudfrontCachePoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
