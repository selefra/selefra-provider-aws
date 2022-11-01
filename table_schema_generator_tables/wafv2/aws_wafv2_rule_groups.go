package wafv2

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsWafv2RuleGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsWafv2RuleGroupsGenerator{}

func (x *TableAwsWafv2RuleGroupsGenerator) GetTableName() string {
	return "aws_wafv2_rule_groups"
}

func (x *TableAwsWafv2RuleGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsWafv2RuleGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsWafv2RuleGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsWafv2RuleGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().WafV2

			config := wafv2.ListRuleGroupsInput{Scope: c.WAFScope}
			for {
				output, err := svc.ListRuleGroups(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.RuleGroups, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().WafV2
					ruleGroupOutput := result.(types.RuleGroupSummary)

					ruleGroup, err := svc.GetRuleGroup(ctx, &wafv2.GetRuleGroupInput{
						Name:	ruleGroupOutput.Name,
						Id:	ruleGroupOutput.Id,
						Scope:	c.WAFScope,
					})
					if err != nil {
						return nil, err
					}
					return ruleGroup.RuleGroup, nil

				})
				if aws.ToString(output.NextMarker) == "" {
					break
				}
				config.NextMarker = output.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsWafv2RuleGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegionAndScope("waf-regional")
}

func (x *TableAwsWafv2RuleGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_labels").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_response_bodies").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("visibility_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("consumed_labels").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label_namespace").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					ruleGroup := result.(*types.RuleGroup)

					cl := client.(*aws_client.Client)
					service := cl.AwsServices().WafV2

					policy, err := service.GetPermissionPolicy(ctx, &wafv2.GetPermissionPolicyInput{ResourceArn: ruleGroup.ARN}, func(options *wafv2.Options) {
						options.Region = cl.Region
					})
					if err != nil {

						var e *types.WAFNonexistentItemException
						if errors.As(err, &e) {
							return "null", nil
						}
						return nil, err
					}
					var p map[string]interface{}
					err = json.Unmarshal([]byte(*policy.Policy), &p)
					if err != nil {
						return nil, err
					}
					return p, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsWafv2RuleGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
