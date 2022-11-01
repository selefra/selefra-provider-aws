package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSagemakerEndpointConfigurationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSagemakerEndpointConfigurationsGenerator{}

func (x *TableAwsSagemakerEndpointConfigurationsGenerator) GetTableName() string {
	return "aws_sagemaker_endpoint_configurations"
}

func (x *TableAwsSagemakerEndpointConfigurationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSagemakerEndpointConfigurationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSagemakerEndpointConfigurationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSagemakerEndpointConfigurationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().SageMaker
			config := sagemaker.ListEndpointConfigsInput{}
			for {
				response, err := svc.ListEndpointConfigs(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.EndpointConfigs, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().SageMaker
					n := result.(types.EndpointConfigSummary)

					response, err := svc.DescribeEndpointConfig(ctx, &sagemaker.DescribeEndpointConfigInput{
						EndpointConfigName: n.EndpointConfigName,
					})
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

func (x *TableAwsSagemakerEndpointConfigurationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("api.sagemaker")
}

func (x *TableAwsSagemakerEndpointConfigurationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("async_inference_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_capture_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("`The tags associated with the model.`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("explainer_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EndpointConfigArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint_config_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("production_variants").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsSagemakerEndpointConfigurationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
