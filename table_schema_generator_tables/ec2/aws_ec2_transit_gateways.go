package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEc2TransitGatewaysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEc2TransitGatewaysGenerator{}

func (x *TableAwsEc2TransitGatewaysGenerator) GetTableName() string {
	return "aws_ec2_transit_gateways"
}

func (x *TableAwsEc2TransitGatewaysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEc2TransitGatewaysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEc2TransitGatewaysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
			"arn",
		},
	}
}

func (x *TableAwsEc2TransitGatewaysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config ec2.DescribeTransitGatewaysInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().EC2
			for {
				output, err := svc.DescribeTransitGateways(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.TransitGateways
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEc2TransitGatewaysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ec2")
}

func (x *TableAwsEc2TransitGatewaysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TransitGatewayId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TransitGatewayArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("options").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsEc2TransitGatewaysGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsEc2TransitGatewayAttachmentsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsEc2TransitGatewayRouteTablesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsEc2TransitGatewayVpcAttachmentsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsEc2TransitGatewayPeeringAttachmentsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsEc2TransitGatewayMulticastDomainsGenerator{}),
	}
}
