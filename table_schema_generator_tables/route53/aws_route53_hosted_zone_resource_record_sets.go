package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRoute53HostedZoneResourceRecordSetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRoute53HostedZoneResourceRecordSetsGenerator{}

func (x *TableAwsRoute53HostedZoneResourceRecordSetsGenerator) GetTableName() string {
	return "aws_route53_hosted_zone_resource_record_sets"
}

func (x *TableAwsRoute53HostedZoneResourceRecordSetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRoute53HostedZoneResourceRecordSetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRoute53HostedZoneResourceRecordSetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsRoute53HostedZoneResourceRecordSetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(*Route53HostedZoneWrapper)
			svc := client.(*aws_client.Client).AwsServices().Route53
			config := route53.ListResourceRecordSetsInput{HostedZoneId: r.Id}
			for {
				response, err := svc.ListResourceRecordSets(ctx, &config, func(options *route53.Options) {})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- response.ResourceRecordSets
				if !response.IsTruncated {
					break
				}

				config.StartRecordIdentifier = response.NextRecordIdentifier
				config.StartRecordType = response.NextRecordType
				config.StartRecordName = response.NextRecordName
			}

			return nil
		},
	}
}

func (x *TableAwsRoute53HostedZoneResourceRecordSetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsRoute53HostedZoneResourceRecordSetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("multi_value_answer").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_records").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr_routing_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("set_identifier").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hosted_zone_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alias_target").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("geo_location").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ttl").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TTL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("traffic_policy_instance_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_route53_hosted_zones_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_route53_hosted_zones.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failover").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_check_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("weight").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsRoute53HostedZoneResourceRecordSetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
