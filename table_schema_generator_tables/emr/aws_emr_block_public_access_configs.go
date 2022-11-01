package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEmrBlockPublicAccessConfigsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEmrBlockPublicAccessConfigsGenerator{}

func (x *TableAwsEmrBlockPublicAccessConfigsGenerator) GetTableName() string {
	return "aws_emr_block_public_access_configs"
}

func (x *TableAwsEmrBlockPublicAccessConfigsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEmrBlockPublicAccessConfigsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEmrBlockPublicAccessConfigsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
		},
	}
}

func (x *TableAwsEmrBlockPublicAccessConfigsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().EMR
			out, err := svc.GetBlockPublicAccessConfiguration(ctx, &emr.GetBlockPublicAccessConfigurationInput{})
			if err != nil {
				if aws_client.IgnoreNotAvailableRegion(err) {

					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- out
			return nil
		},
	}
}

func (x *TableAwsEmrBlockPublicAccessConfigsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticmapreduce")
}

func (x *TableAwsEmrBlockPublicAccessConfigsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("block_public_access_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("block_public_access_configuration_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsEmrBlockPublicAccessConfigsGenerator) GetSubTables() []*schema.Table {
	return nil
}
