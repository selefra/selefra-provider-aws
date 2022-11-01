package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLambdaLayerVersionPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLambdaLayerVersionPoliciesGenerator{}

func (x *TableAwsLambdaLayerVersionPoliciesGenerator) GetTableName() string {
	return "aws_lambda_layer_version_policies"
}

func (x *TableAwsLambdaLayerVersionPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLambdaLayerVersionPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLambdaLayerVersionPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLambdaLayerVersionPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(types.LayerVersionsListItem)

			pp := task.ParentTask.ParentRawResult.(types.LayersListItem)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lambda

			config := lambda.GetLayerVersionPolicyInput{
				LayerName:	pp.LayerName,
				VersionNumber:	p.Version,
			}

			output, err := svc.GetLayerVersionPolicy(ctx, &config)
			if err != nil {
				if aws_client.IsAWSError(err, "ResourceNotFoundException") {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output

			return nil
		},
	}
}

func (x *TableAwsLambdaLayerVersionPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lambda")
}

func (x *TableAwsLambdaLayerVersionPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("layer_version_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("revision_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_lambda_layer_versions_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_lambda_layer_versions.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("layer_version").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.ParentColumnValue("version")).Build(),
	}
}

func (x *TableAwsLambdaLayerVersionPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
