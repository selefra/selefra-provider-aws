package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsWafv2RegexPatternSetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsWafv2RegexPatternSetsGenerator{}

func (x *TableAwsWafv2RegexPatternSetsGenerator) GetTableName() string {
	return "aws_wafv2_regex_pattern_sets"
}

func (x *TableAwsWafv2RegexPatternSetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsWafv2RegexPatternSetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsWafv2RegexPatternSetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsWafv2RegexPatternSetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().WafV2

			params := wafv2.ListRegexPatternSetsInput{
				Scope:	cl.WAFScope,
				Limit:	aws.Int32(100),
			}
			for {
				result, err := svc.ListRegexPatternSets(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, result.RegexPatternSets, func(result any) (any, error) {
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().WafV2
					s := result.(types.RegexPatternSetSummary)

					info, err := svc.GetRegexPatternSet(
						ctx,
						&wafv2.GetRegexPatternSetInput{
							Id:	s.Id,
							Name:	s.Name,
							Scope:	cl.WAFScope,
						},
						func(options *wafv2.Options) {
							options.Region = cl.Region
						},
					)
					if err != nil {
						return nil, err
					}
					return info.RegexPatternSet, nil

				})
				if aws.ToString(result.NextMarker) == "" {
					break
				}
				params.NextMarker = result.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsWafv2RegexPatternSetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegionAndScope("waf-regional")
}

func (x *TableAwsWafv2RegexPatternSetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("regular_expression_list").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsWafv2RegexPatternSetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
