package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsGlueSecurityConfigurationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlueSecurityConfigurationsGenerator{}

func (x *TableAwsGlueSecurityConfigurationsGenerator) GetTableName() string {
	return "aws_glue_security_configurations"
}

func (x *TableAwsGlueSecurityConfigurationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlueSecurityConfigurationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlueSecurityConfigurationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"name",
		},
	}
}

func (x *TableAwsGlueSecurityConfigurationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().Glue
			input := glue.GetSecurityConfigurationsInput{}
			for {
				result, err := svc.GetSecurityConfigurations(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- result.SecurityConfigurations
				if aws.ToString(result.NextToken) == "" {
					break
				}
				input.NextToken = result.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsGlueSecurityConfigurationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glue")
}

func (x *TableAwsGlueSecurityConfigurationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_time_stamp").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsGlueSecurityConfigurationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
