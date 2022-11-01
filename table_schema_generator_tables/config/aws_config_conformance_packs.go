package config

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/smithy-go"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsConfigConformancePacksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsConfigConformancePacksGenerator{}

func (x *TableAwsConfigConformancePacksGenerator) GetTableName() string {
	return "aws_config_conformance_packs"
}

func (x *TableAwsConfigConformancePacksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsConfigConformancePacksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsConfigConformancePacksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsConfigConformancePacksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			config := configservice.DescribeConformancePacksInput{}
			var ae smithy.APIError
			for {
				resp, err := c.AwsServices().ConfigService.DescribeConformancePacks(ctx, &config)

				if (c.Region == "af-south-1" || c.Region == "ap-northeast-3") && errors.As(err, &ae) && ae.ErrorCode() == "AccessDeniedException" {
					return nil
				}

				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- resp.ConformancePackDetails
				if resp.NextToken == nil {
					break
				}
				config.NextToken = resp.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsConfigConformancePacksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("config")
}

func (x *TableAwsConfigConformancePacksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("conformance_pack_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conformance_pack_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delivery_s3_bucket").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_ssm_document_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TemplateSSMDocumentDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ConformancePackArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delivery_s3_key_prefix").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_update_requested_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsConfigConformancePacksGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsConfigConformancePackRuleCompliancesGenerator{}),
	}
}
