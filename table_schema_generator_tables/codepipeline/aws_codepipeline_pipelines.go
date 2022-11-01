package codepipeline

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCodepipelinePipelinesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCodepipelinePipelinesGenerator{}

func (x *TableAwsCodepipelinePipelinesGenerator) GetTableName() string {
	return "aws_codepipeline_pipelines"
}

func (x *TableAwsCodepipelinePipelinesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCodepipelinePipelinesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCodepipelinePipelinesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsCodepipelinePipelinesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().CodePipeline
			config := codepipeline.ListPipelinesInput{}
			for {
				response, err := svc.ListPipelines(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.Pipelines, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().CodePipeline
					item := result.(types.PipelineSummary)
					response, err := svc.GetPipeline(ctx, &codepipeline.GetPipelineInput{Name: item.Name})
					if err != nil {
						return nil, err
					}
					return response, nil

				})
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsCodepipelinePipelinesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("codepipeline")
}

func (x *TableAwsCodepipelinePipelinesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					pipeline := result.(*codepipeline.GetPipelineOutput)

					a := arn.ARN{
						Partition:	cl.Partition,
						Service:	"codepipeline",
						Region:		cl.Region,
						AccountID:	cl.AccountID,
						Resource:	"pipelines/" + *pipeline.Metadata.PipelineArn,
					}

					return a.String(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pipeline").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsCodepipelinePipelinesGenerator) GetSubTables() []*schema.Table {
	return nil
}
