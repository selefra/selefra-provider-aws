package xray

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsXrayEncryptionConfigGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsXrayEncryptionConfigGenerator{}

func (x *TableAwsXrayEncryptionConfigGenerator) GetTableName() string {
	return "aws_xray_encryption_config"
}

func (x *TableAwsXrayEncryptionConfigGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsXrayEncryptionConfigGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsXrayEncryptionConfigGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsXrayEncryptionConfigGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Xray
			input := xray.GetEncryptionConfigInput{}
			output, err := svc.GetEncryptionConfig(ctx, &input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output.EncryptionConfig
			return nil
		},
	}
}

func (x *TableAwsXrayEncryptionConfigGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("xray")
}

func (x *TableAwsXrayEncryptionConfigGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_id").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsXrayEncryptionConfigGenerator) GetSubTables() []*schema.Table {
	return nil
}
