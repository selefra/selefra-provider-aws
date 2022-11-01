package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsWafRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsWafRulesGenerator{}

func (x *TableAwsWafRulesGenerator) GetTableName() string {
	return "aws_waf_rules"
}

func (x *TableAwsWafRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsWafRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsWafRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsWafRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			service := c.AwsServices().Waf
			config := waf.ListRulesInput{}
			for {
				output, err := service.ListRules(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, ruleSum := range output.Rules {
					rule, err := service.GetRule(ctx, &waf.GetRuleInput{RuleId: ruleSum.RuleId}, func(options *waf.Options) {
						options.Region = c.Region
					})
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- rule.Rule
				}

				if aws.ToString(output.NextMarker) == "" {
					break
				}
				config.NextMarker = output.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsWafRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsWafRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					rule := result.(*types.Rule)
					return cl.ARN("waf", "rule", aws.ToString(rule.RuleId)), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_id").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsWafRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}
