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

type TableAwsWafRuleGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsWafRuleGroupsGenerator{}

func (x *TableAwsWafRuleGroupsGenerator) GetTableName() string {
	return "aws_waf_rule_groups"
}

func (x *TableAwsWafRuleGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsWafRuleGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsWafRuleGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsWafRuleGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			service := c.AwsServices().Waf
			config := waf.ListRuleGroupsInput{}
			for {
				output, err := service.ListRuleGroups(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, rG := range output.RuleGroups {
					ruleGroup, err := service.GetRuleGroup(ctx, &waf.GetRuleGroupInput{RuleGroupId: rG.RuleGroupId}, func(options *waf.Options) {
						options.Region = c.Region
					})
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- ruleGroup.RuleGroup
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

func (x *TableAwsWafRuleGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsWafRuleGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					ruleGroup := result.(*types.RuleGroup)
					return cl.ARN("waf", "rulegroup", aws.ToString(ruleGroup.RuleGroupId)), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_ids").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					ruleGroup := result.(*types.RuleGroup)

					awsClient := client.(*aws_client.Client)
					service := awsClient.AwsServices().Waf
					listActivatedRulesConfig := waf.ListActivatedRulesInRuleGroupInput{RuleGroupId: ruleGroup.RuleGroupId}
					var ruleIDs []string
					for {
						rules, err := service.ListActivatedRulesInRuleGroup(ctx, &listActivatedRulesConfig)
						if err != nil {
							return nil, err
						}
						for _, rule := range rules.ActivatedRules {
							ruleIDs = append(ruleIDs, aws.ToString(rule.RuleId))
						}

						if aws.ToString(rules.NextMarker) == "" {
							break
						}
						listActivatedRulesConfig.NextMarker = rules.NextMarker
					}
					return ruleIDs, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_group_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsWafRuleGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
