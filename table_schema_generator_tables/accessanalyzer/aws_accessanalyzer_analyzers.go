package accessanalyzer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/smithy-go/middleware"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsAccessanalyzerAnalyzersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsAccessanalyzerAnalyzersGenerator{}

func (x *TableAwsAccessanalyzerAnalyzersGenerator) GetTableName() string {
	return "aws_accessanalyzer_analyzers"
}

func (x *TableAwsAccessanalyzerAnalyzersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsAccessanalyzerAnalyzersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsAccessanalyzerAnalyzersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsAccessanalyzerAnalyzersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			config := accessanalyzer.ListAnalyzersInput{}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Analyzer
			for {
				response, err := svc.ListAnalyzers(ctx, &config, func(options *accessanalyzer.Options) {
					options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
						if err := stack.Initialize.Add(&awsmiddleware.RegisterServiceMetadata{
							Region:		c.Region,
							ServiceID:	accessanalyzer.ServiceID,
							SigningName:	"access-analyzer",
							OperationName:	"ListAnalyzers",
						}, middleware.Before); err != nil {
							return nil
						}
						return nil
					})
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- response.Analyzers
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsAccessanalyzerAnalyzersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("access-analyzer")
}

func (x *TableAwsAccessanalyzerAnalyzersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_resource_analyzed").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_resource_analyzed_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_reason").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsAccessanalyzerAnalyzersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsAccessanalyzerAnalyzerArchiveRulesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsAccessanalyzerAnalyzerFindingsGenerator{}),
	}
}
