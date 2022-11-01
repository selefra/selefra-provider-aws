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

type TableAwsWafregionalWebAclsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsWafregionalWebAclsGenerator{}

func (x *TableAwsWafregionalWebAclsGenerator) GetTableName() string {
	return "aws_wafregional_web_acls"
}

func (x *TableAwsWafregionalWebAclsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsWafregionalWebAclsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsWafregionalWebAclsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsWafregionalWebAclsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().WafRegional
			var params wafregional.ListWebACLsInput
			for {
				result, err := svc.ListWebACLs(ctx, &params, func(o *wafregional.Options) {
					o.Region = cl.Region
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, item := range result.WebACLs {
					detail, err := svc.GetWebACL(
						ctx,
						&wafregional.GetWebACLInput{WebACLId: item.WebACLId},
						func(o *wafregional.Options) {
							o.Region = cl.Region
						},
					)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					if detail.WebACL == nil {
						continue
					}
					resultChannel <- *detail.WebACL
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

func (x *TableAwsWafregionalWebAclsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("waf-regional")
}

func (x *TableAwsWafregionalWebAclsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("default_action").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources_for_web_acl").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					service := client.(*aws_client.Client).AwsServices().WafRegional
					output, err := service.ListResourcesForWebACL(ctx, &wafregional.ListResourcesForWebACLInput{
						WebACLId: result.(types.WebACL).WebACLId,
					})
					if err != nil {
						return nil, err
					}
					return output.ResourceArns, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_acl_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WebACLId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WebACLArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("`Web ACL tags.`").Build(),
	}
}

func (x *TableAwsWafregionalWebAclsGenerator) GetSubTables() []*schema.Table {
	return nil
}
