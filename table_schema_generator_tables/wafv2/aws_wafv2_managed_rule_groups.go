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

type TableAwsWafv2ManagedRuleGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsWafv2ManagedRuleGroupsGenerator{}

func (x *TableAwsWafv2ManagedRuleGroupsGenerator) GetTableName() string {
	return "aws_wafv2_managed_rule_groups"
}

func (x *TableAwsWafv2ManagedRuleGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsWafv2ManagedRuleGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsWafv2ManagedRuleGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"scope",
		},
	}
}

func (x *TableAwsWafv2ManagedRuleGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			service := c.AwsServices().WafV2

			config := wafv2.ListAvailableManagedRuleGroupsInput{Scope: c.WAFScope}
			for {
				output, err := service.ListAvailableManagedRuleGroups(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.ManagedRuleGroups

				if aws.ToString(output.NextMarker) == "" {
					break
				}
				config.NextMarker = output.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsWafv2ManagedRuleGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegionAndScope("waf-regional")
}

func (x *TableAwsWafv2ManagedRuleGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scope").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.WAFScopeExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					managedRuleGroupSum := result.(types.ManagedRuleGroupSummary)

					c := client.(*aws_client.Client)
					service := c.AwsServices().WafV2

					output, err := service.DescribeManagedRuleGroup(ctx, &wafv2.DescribeManagedRuleGroupInput{
						Name:		managedRuleGroupSum.Name,
						VendorName:	managedRuleGroupSum.VendorName,
						Scope:		c.WAFScope,
					}, func(options *wafv2.Options) {
						options.Region = c.Region
					})
					if err != nil {
						return nil, err
					}
					return output, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vendor_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("versioning_supported").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsWafv2ManagedRuleGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
