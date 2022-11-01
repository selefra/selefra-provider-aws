package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDirectconnectGatewayAssociationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDirectconnectGatewayAssociationsGenerator{}

func (x *TableAwsDirectconnectGatewayAssociationsGenerator) GetTableName() string {
	return "aws_directconnect_gateway_associations"
}

func (x *TableAwsDirectconnectGatewayAssociationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDirectconnectGatewayAssociationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDirectconnectGatewayAssociationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsDirectconnectGatewayAssociationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			gateway := task.ParentRawResult.(types.DirectConnectGateway)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Directconnect
			config := directconnect.DescribeDirectConnectGatewayAssociationsInput{DirectConnectGatewayId: gateway.DirectConnectGatewayId}
			for {
				output, err := svc.DescribeDirectConnectGatewayAssociations(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.DirectConnectGatewayAssociations
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsDirectconnectGatewayAssociationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("directconnect")
}

func (x *TableAwsDirectconnectGatewayAssociationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("direct_connect_gateway_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("virtual_gateway_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("virtual_gateway_owner_account").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_directconnect_gateways_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_directconnect_gateways.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gateway_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("association_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_change_error").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gateway_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("associated_gateway").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("association_state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_prefixes_to_direct_connect_gateway").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("direct_connect_gateway_owner_account").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("virtual_gateway_region").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsDirectconnectGatewayAssociationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
