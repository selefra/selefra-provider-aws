package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsWafSubscribedRuleGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsWafSubscribedRuleGroupsGenerator{}

func (x *TableAwsWafSubscribedRuleGroupsGenerator) GetTableName() string {
	return "aws_waf_subscribed_rule_groups"
}

func (x *TableAwsWafSubscribedRuleGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsWafSubscribedRuleGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsWafSubscribedRuleGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"rule_group_id",
		},
	}
}

func (x *TableAwsWafSubscribedRuleGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			service := c.AwsServices().Waf
			config := waf.ListSubscribedRuleGroupsInput{}
			for {
				output, err := service.ListSubscribedRuleGroups(ctx, &config, func(options *waf.Options) {

					options.Region = "us-east-1"
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.RuleGroups

				if aws.ToString(output.NextMarker) == "" {
					break
				}
				config.NextMarker = output.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsWafSubscribedRuleGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsWafSubscribedRuleGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).Description("`The AWS Account ID of the resource.`").
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_group_id").ColumnType(schema.ColumnTypeString).Description("`A unique identifier for a RuleGroup.`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsWafSubscribedRuleGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
