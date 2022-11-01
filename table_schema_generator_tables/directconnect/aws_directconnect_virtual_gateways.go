package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDirectconnectVirtualGatewaysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDirectconnectVirtualGatewaysGenerator{}

func (x *TableAwsDirectconnectVirtualGatewaysGenerator) GetTableName() string {
	return "aws_directconnect_virtual_gateways"
}

func (x *TableAwsDirectconnectVirtualGatewaysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDirectconnectVirtualGatewaysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDirectconnectVirtualGatewaysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"id",
		},
	}
}

func (x *TableAwsDirectconnectVirtualGatewaysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config directconnect.DescribeVirtualGatewaysInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Directconnect
			output, err := svc.DescribeVirtualGateways(ctx, &config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output.VirtualGateways
			return nil
		},
	}
}

func (x *TableAwsDirectconnectVirtualGatewaysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("directconnect")
}

func (x *TableAwsDirectconnectVirtualGatewaysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VirtualGatewayId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("virtual_gateway_state").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsDirectconnectVirtualGatewaysGenerator) GetSubTables() []*schema.Table {
	return nil
}
