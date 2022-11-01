package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLightsailDatabaseParametersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailDatabaseParametersGenerator{}

func (x *TableAwsLightsailDatabaseParametersGenerator) GetTableName() string {
	return "aws_lightsail_database_parameters"
}

func (x *TableAwsLightsailDatabaseParametersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailDatabaseParametersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailDatabaseParametersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLightsailDatabaseParametersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.RelationalDatabase)
			input := lightsail.GetRelationalDatabaseParametersInput{
				RelationalDatabaseName: r.Name,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			for {
				response, err := svc.GetRelationalDatabaseParameters(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Parameters
				if aws.ToString(response.NextPageToken) == "" {
					break
				}
				input.PageToken = response.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableAwsLightsailDatabaseParametersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailDatabaseParametersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameter_value").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_values").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("apply_method").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("apply_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_modifiable").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parameter_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_lightsail_databases_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_lightsail_databases.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsLightsailDatabaseParametersGenerator) GetSubTables() []*schema.Table {
	return nil
}
