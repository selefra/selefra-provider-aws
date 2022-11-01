package elbv2

import (
	"context"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElbv2ListenerCertificatesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElbv2ListenerCertificatesGenerator{}

func (x *TableAwsElbv2ListenerCertificatesGenerator) GetTableName() string {
	return "aws_elbv2_listener_certificates"
}

func (x *TableAwsElbv2ListenerCertificatesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElbv2ListenerCertificatesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElbv2ListenerCertificatesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsElbv2ListenerCertificatesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			region := c.Region
			svc := c.AwsServices().ELBv2
			listener := task.ParentRawResult.(types.Listener)
			config := elbv2.DescribeListenerCertificatesInput{ListenerArn: listener.ListenerArn}
			for {
				response, err := svc.DescribeListenerCertificates(ctx, &config, func(options *elbv2.Options) {
					options.Region = region
				})
				if err != nil {
					if aws_client.IsErrorRegex(err, "ValidationError", notSupportedGatewayLB) {

						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Certificates
				if aws.ToString(response.NextMarker) == "" {
					break
				}
				config.Marker = response.NextMarker
			}
			return nil
		},
	}
}

var notSupportedGatewayLB = regexp.MustCompile("This operation does not support Gateway Load Balancer Listeners")

func (x *TableAwsElbv2ListenerCertificatesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticloadbalancing")
}

func (x *TableAwsElbv2ListenerCertificatesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("aws_elbv2_listeners_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_elbv2_listeners.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("listener_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_default").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsElbv2ListenerCertificatesGenerator) GetSubTables() []*schema.Table {
	return nil
}
