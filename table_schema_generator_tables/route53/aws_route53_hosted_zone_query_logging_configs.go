package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRoute53HostedZoneQueryLoggingConfigsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRoute53HostedZoneQueryLoggingConfigsGenerator{}

func (x *TableAwsRoute53HostedZoneQueryLoggingConfigsGenerator) GetTableName() string {
	return "aws_route53_hosted_zone_query_logging_configs"
}

func (x *TableAwsRoute53HostedZoneQueryLoggingConfigsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRoute53HostedZoneQueryLoggingConfigsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRoute53HostedZoneQueryLoggingConfigsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRoute53HostedZoneQueryLoggingConfigsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(*Route53HostedZoneWrapper)
			svc := client.(*aws_client.Client).AwsServices().Route53
			config := route53.ListQueryLoggingConfigsInput{HostedZoneId: r.Id}
			for {
				response, err := svc.ListQueryLoggingConfigs(ctx, &config, func(options *route53.Options) {})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.QueryLoggingConfigs
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

type Route53HostedZoneWrapper struct {
	types.HostedZone
	Tags		map[string]string
	DelegationSetId	*string
	VPCs		[]types.VPC
}

func (x *TableAwsRoute53HostedZoneQueryLoggingConfigsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsRoute53HostedZoneQueryLoggingConfigsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("cloud_watch_logs_log_group_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hosted_zone_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_route53_hosted_zones_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_route53_hosted_zones.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					ql := result.(types.QueryLoggingConfig)
					return cl.PartitionGlobalARN("route53", "queryloggingconfig", *ql.Id), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hosted_zone_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
	}
}

func (x *TableAwsRoute53HostedZoneQueryLoggingConfigsGenerator) GetSubTables() []*schema.Table {
	return nil
}
