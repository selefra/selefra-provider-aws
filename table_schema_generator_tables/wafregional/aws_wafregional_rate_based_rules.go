package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsWafregionalRateBasedRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsWafregionalRateBasedRulesGenerator{}

func (x *TableAwsWafregionalRateBasedRulesGenerator) GetTableName() string {
	return "aws_wafregional_rate_based_rules"
}

func (x *TableAwsWafregionalRateBasedRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsWafregionalRateBasedRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsWafregionalRateBasedRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsWafregionalRateBasedRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().WafRegional
			var params wafregional.ListRateBasedRulesInput
			for {
				result, err := svc.ListRateBasedRules(ctx, &params, func(o *wafregional.Options) {
					o.Region = cl.Region
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, item := range result.Rules {
					detail, err := svc.GetRateBasedRule(
						ctx,
						&wafregional.GetRateBasedRuleInput{RuleId: item.RuleId},
						func(o *wafregional.Options) {
							o.Region = cl.Region
						},
					)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					if detail.Rule == nil {
						continue
					}
					resultChannel <- *detail.Rule
				}
				if aws.ToString(result.NextMarker) == "" {
					break
				}
				params.NextMarker = result.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsWafregionalRateBasedRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("waf-regional")
}

func (x *TableAwsWafregionalRateBasedRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rate_limit").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				rateBasedRuleARN := func(meta any, id string) string {
					cl := meta.(*aws_client.Client)
					return cl.ARN("waf-regional", "ratebasedrule", id)
				}
				extractor := func() (any, error) {
					return rateBasedRuleARN(client, *result.(types.RateBasedRule).RuleId), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("match_predicates").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rate_key").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsWafregionalRateBasedRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}
