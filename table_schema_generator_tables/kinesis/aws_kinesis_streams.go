package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsKinesisStreamsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsKinesisStreamsGenerator{}

func (x *TableAwsKinesisStreamsGenerator) GetTableName() string {
	return "aws_kinesis_streams"
}

func (x *TableAwsKinesisStreamsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsKinesisStreamsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsKinesisStreamsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsKinesisStreamsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Kinesis
			input := kinesis.ListStreamsInput{}
			for {
				response, err := svc.ListStreams(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.StreamNames, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					streamName := result.(string)
					svc := c.AwsServices().Kinesis
					streamSummary, err := svc.DescribeStreamSummary(ctx, &kinesis.DescribeStreamSummaryInput{
						StreamName: aws.String(streamName),
					})
					if err != nil {
						return nil, err
					}
					return streamSummary.StreamDescriptionSummary, nil

				})
				if !aws.ToBool(response.HasMoreStreams) {
					break
				}
				input.ExclusiveStartStreamName = aws.String(response.StreamNames[len(response.StreamNames)-1])
			}
			return nil
		},
	}
}

func (x *TableAwsKinesisStreamsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("kinesis")
}

func (x *TableAwsKinesisStreamsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stream_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stream_creation_timestamp").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stream_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stream_mode_details").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StreamARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enhanced_monitoring").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_period_hours").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("open_shard_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("consumer_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsKinesisStreamsGenerator) GetSubTables() []*schema.Table {
	return nil
}
