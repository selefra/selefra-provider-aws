package accessanalyzer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAccessanalyzerAnalyzerArchiveRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAccessanalyzerAnalyzerArchiveRulesGenerator{}

func (x *TableAwsAccessanalyzerAnalyzerArchiveRulesGenerator) GetTableName() string {
	return "aws_accessanalyzer_analyzer_archive_rules"
}

func (x *TableAwsAccessanalyzerAnalyzerArchiveRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAccessanalyzerAnalyzerArchiveRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAccessanalyzerAnalyzerArchiveRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsAccessanalyzerAnalyzerArchiveRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			analyzer := task.ParentRawResult.(types.AnalyzerSummary)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Analyzer
			config := accessanalyzer.ListArchiveRulesInput{
				AnalyzerName: analyzer.Name,
			}
			for {
				response, err := svc.ListArchiveRules(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- response.ArchiveRules
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsAccessanalyzerAnalyzerArchiveRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAwsAccessanalyzerAnalyzerArchiveRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("analyzer_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filter").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_accessanalyzer_analyzers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_accessanalyzer_analyzers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsAccessanalyzerAnalyzerArchiveRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}
