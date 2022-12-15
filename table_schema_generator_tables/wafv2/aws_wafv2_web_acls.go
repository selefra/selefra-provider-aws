package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsWafv2WebAclsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsWafv2WebAclsGenerator{}

func (x *TableAwsWafv2WebAclsGenerator) GetTableName() string {
	return "aws_wafv2_web_acls"
}

func (x *TableAwsWafv2WebAclsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsWafv2WebAclsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsWafv2WebAclsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsWafv2WebAclsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			service := c.AwsServices().WafV2

			config := wafv2.ListWebACLsInput{
				Scope: c.WAFScope,
				Limit: aws.Int32(100),
			}
			for {
				output, err := service.ListWebACLs(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.WebACLs, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().WafV2
					webAcl := result.(types.WebACLSummary)

					webAclConfig := wafv2.GetWebACLInput{Id: webAcl.Id, Name: webAcl.Name, Scope: c.WAFScope}
					webAclOutput, err := svc.GetWebACL(ctx, &webAclConfig, func(options *wafv2.Options) {
						options.Region = c.Region
					})
					if err != nil {
						return nil, err
					}

					cfg := wafv2.GetLoggingConfigurationInput{
						ResourceArn: webAclOutput.WebACL.ARN,
					}

					loggingConfigurationOutput, err := svc.GetLoggingConfiguration(ctx, &cfg, func(options *wafv2.Options) {
						options.Region = c.Region
					})
					if err != nil {
						if aws_client.IsAWSError(err, "WAFNonexistentItemException") {

						} else {

						}
					}

					var webAclLoggingConfiguration *types.LoggingConfiguration
					if loggingConfigurationOutput != nil {
						webAclLoggingConfiguration = loggingConfigurationOutput.LoggingConfiguration
					}
					return &WebACLWrapper{
						WebACL:               webAclOutput.WebACL,
						LoggingConfiguration: webAclLoggingConfiguration,
					}, nil

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

type WebACLWrapper struct {
	*types.WebACL
	LoggingConfiguration *types.LoggingConfiguration
}

func (x *TableAwsWafv2WebAclsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegionAndScope("waf-regional")
}

func (x *TableAwsWafv2WebAclsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("visibility_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("post_process_firewall_manager_rule_groups").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logging_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label_namespace").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pre_process_firewall_manager_rule_groups").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources_for_web_acl").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					webACL := result.(*WebACLWrapper)

					cl := client.(*aws_client.Client)
					service := cl.AwsServices().WafV2

					resourceArns := []string{}
					if cl.WAFScope == types.ScopeCloudfront {
						cloudfrontService := cl.AwsServices().Cloudfront
						params := &cloudfront.ListDistributionsByWebACLIdInput{
							WebACLId: webACL.Id,
							MaxItems: aws.Int32(100),
						}
						for {
							output, err := cloudfrontService.ListDistributionsByWebACLId(ctx, params, func(options *cloudfront.Options) {
								options.Region = cl.Region
							})
							if err != nil {
								return nil, err
							}
							for _, item := range output.DistributionList.Items {
								resourceArns = append(resourceArns, *item.ARN)
							}
							if aws.ToString(output.DistributionList.NextMarker) == "" {
								break
							}
							params.Marker = output.DistributionList.NextMarker
						}
					} else {
						output, err := service.ListResourcesForWebACL(ctx, &wafv2.ListResourcesForWebACLInput{WebACLArn: webACL.ARN})
						if err != nil {
							return nil, err
						}
						resourceArns = output.ResourceArns
					}
					return resourceArns, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("captcha_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_response_bodies").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("managed_by_firewall_manager").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_action").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsWafv2WebAclsGenerator) GetSubTables() []*schema.Table {
	return nil
}
