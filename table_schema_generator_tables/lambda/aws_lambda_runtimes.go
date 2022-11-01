package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLambdaRuntimesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLambdaRuntimesGenerator{}

func (x *TableAwsLambdaRuntimesGenerator) GetTableName() string {
	return "aws_lambda_runtimes"
}

func (x *TableAwsLambdaRuntimesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLambdaRuntimesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLambdaRuntimesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"name",
		},
	}
}

func (x *TableAwsLambdaRuntimesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			runtimes := make([]RuntimeWrapper, len(types.RuntimeProvidedal2.Values()))
			for i, runtime := range types.RuntimeProvidedal2.Values() {
				runtimes[i] = RuntimeWrapper{
					Name: string(runtime),
				}
			}
			resultChannel <- runtimes
			return nil
		},
	}
}

type RuntimeWrapper struct {
	Name string
}

func (x *TableAwsLambdaRuntimesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lambda")
}

func (x *TableAwsLambdaRuntimesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsLambdaRuntimesGenerator) GetSubTables() []*schema.Table {
	return nil
}
