package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEc2TransitGatewayRouteTablesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEc2TransitGatewayRouteTablesGenerator{}

func (x *TableAwsEc2TransitGatewayRouteTablesGenerator) GetTableName() string {
	return "aws_ec2_transit_gateway_route_tables"
}

func (x *TableAwsEc2TransitGatewayRouteTablesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEc2TransitGatewayRouteTablesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEc2TransitGatewayRouteTablesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsEc2TransitGatewayRouteTablesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.TransitGateway)

			config := ec2.DescribeTransitGatewayRouteTablesInput{
				Filters: []types.Filter{
					{
						Name:	aws.String("transit-gateway-id"),
						Values:	[]string{*r.TransitGatewayId},
					},
				},
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().EC2
			for {
				output, err := svc.DescribeTransitGatewayRouteTables(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.TransitGatewayRouteTables
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEc2TransitGatewayRouteTablesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ec2")
}

func (x *TableAwsEc2TransitGatewayRouteTablesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_ec2_transit_gateways_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_ec2_transit_gateways.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transit_gateway_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transit_gateway_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_association_route_table").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_propagation_route_table").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transit_gateway_route_table_id").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsEc2TransitGatewayRouteTablesGenerator) GetSubTables() []*schema.Table {
	return nil
}
