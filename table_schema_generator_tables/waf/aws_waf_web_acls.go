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

type TableAwsWafWebAclsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsWafWebAclsGenerator{}

func (x *TableAwsWafWebAclsGenerator) GetTableName() string {
	return "aws_waf_web_acls"
}

func (x *TableAwsWafWebAclsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsWafWebAclsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsWafWebAclsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsWafWebAclsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			service := c.AwsServices().Waf
			config := waf.ListWebACLsInput{}
			for {
				output, err := service.ListWebACLs(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, webAcl := range output.WebACLs {
					webAclConfig := waf.GetWebACLInput{WebACLId: webAcl.WebACLId}
					webAclOutput, err := service.GetWebACL(ctx, &webAclConfig, func(options *waf.Options) {
						options.Region = c.Region
					})
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}

					cfg := waf.GetLoggingConfigurationInput{
						ResourceArn: webAclOutput.WebACL.WebACLArn,
					}
					loggingConfigurationOutput, err := service.GetLoggingConfiguration(ctx, &cfg, func(options *waf.Options) {
						options.Region = c.Region
					})
					if err != nil {

					}

					var webAclLoggingConfiguration *types.LoggingConfiguration
					if loggingConfigurationOutput != nil {
						webAclLoggingConfiguration = loggingConfigurationOutput.LoggingConfiguration
					}

					resultChannel <- &WebACLWrapper{
						webAclOutput.WebACL,
						webAclLoggingConfiguration,
					}
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

type WebACLWrapper struct {
	*types.WebACL
	LoggingConfiguration	*types.LoggingConfiguration
}

func (x *TableAwsWafWebAclsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsWafWebAclsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("logging_configuration").ColumnType(schema.ColumnTypeJSON).Description("`The LoggingConfiguration for the specified web ACL.`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_acl_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WebACLId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WebACLArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsWafWebAclsGenerator) GetSubTables() []*schema.Table {
	return nil
}
