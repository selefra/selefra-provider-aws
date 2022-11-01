package firehose

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsFirehoseDeliveryStreamsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsFirehoseDeliveryStreamsGenerator{}

func (x *TableAwsFirehoseDeliveryStreamsGenerator) GetTableName() string {
	return "aws_firehose_delivery_streams"
}

func (x *TableAwsFirehoseDeliveryStreamsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsFirehoseDeliveryStreamsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsFirehoseDeliveryStreamsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsFirehoseDeliveryStreamsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Firehose
			input := firehose.ListDeliveryStreamsInput{}
			for {
				response, err := svc.ListDeliveryStreams(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.DeliveryStreamNames, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					streamName := result.(string)
					svc := c.AwsServices().Firehose
					streamSummary, err := svc.DescribeDeliveryStream(ctx, &firehose.DescribeDeliveryStreamInput{
						DeliveryStreamName: aws.String(streamName),
					})
					if err != nil {
						return nil, err
					}
					return streamSummary.DeliveryStreamDescription, nil

				})
				if !aws.ToBool(response.HasMoreDeliveryStreams) {
					break
				}
				input.ExclusiveStartDeliveryStreamName = aws.String(response.DeliveryStreamNames[len(response.DeliveryStreamNames)-1])
			}
			return nil
		},
	}
}

func (x *TableAwsFirehoseDeliveryStreamsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("firehose")
}

func (x *TableAwsFirehoseDeliveryStreamsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delivery_stream_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_update_timestamp").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeliveryStreamARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("destinations").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failure_description").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delivery_stream_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_more_destinations").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_timestamp").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delivery_stream_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delivery_stream_encryption_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsFirehoseDeliveryStreamsGenerator) GetSubTables() []*schema.Table {
	return nil
}
