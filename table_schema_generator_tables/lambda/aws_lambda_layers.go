package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLambdaLayersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLambdaLayersGenerator{}

func (x *TableAwsLambdaLayersGenerator) GetTableName() string {
	return "aws_lambda_layers"
}

func (x *TableAwsLambdaLayersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLambdaLayersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLambdaLayersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsLambdaLayersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input lambda.ListLayersInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lambda
			for {
				response, err := svc.ListLayers(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- response.Layers

				if aws.ToString(response.NextMarker) == "" {
					break
				}
				input.Marker = response.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsLambdaLayersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lambda")
}

func (x *TableAwsLambdaLayersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("layer_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LayerArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_matching_version").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("layer_arn").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsLambdaLayersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsLambdaLayerVersionsGenerator{}),
	}
}
