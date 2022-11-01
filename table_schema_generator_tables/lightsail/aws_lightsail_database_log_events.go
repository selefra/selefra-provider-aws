package lightsail

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"golang.org/x/sync/errgroup"
)

type TableAwsLightsailDatabaseLogEventsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailDatabaseLogEventsGenerator{}

func (x *TableAwsLightsailDatabaseLogEventsGenerator) GetTableName() string {
	return "aws_lightsail_database_log_events"
}

func (x *TableAwsLightsailDatabaseLogEventsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailDatabaseLogEventsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailDatabaseLogEventsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLightsailDatabaseLogEventsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.RelationalDatabase)
			input := lightsail.GetRelationalDatabaseLogStreamsInput{
				RelationalDatabaseName: r.Name,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			streams, err := svc.GetRelationalDatabaseLogStreams(ctx, &input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			endTime := time.Now()
			startTime := endTime.Add(-time.Hour * 24 * 14)
			errs, ctx := errgroup.WithContext(ctx)
			errs.SetLimit(MaxGoroutines)
			for _, s := range streams.LogStreams {
				func(database, stream string, startTime, endTime time.Time) {
					errs.Go(func() error {
						return fetchLogEvents(ctx, resultChannel, c, database, stream, startTime, endTime)
					})
				}(*r.Name, s, startTime, endTime)
			}
			err = errs.Wait()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			return nil
		},
	}
}

var MaxGoroutines = 10

type LogEventWrapper struct {
	types.LogEvent

	LogStreamName	string
}

func fetchLogEvents(ctx context.Context, res chan<- interface{}, c *aws_client.Client, database, stream string, startTime, endTime time.Time) error {
	svc := c.AwsServices().Lightsail
	input := lightsail.GetRelationalDatabaseLogEventsInput{
		RelationalDatabaseName:	&database,
		LogStreamName:		&stream,
		StartTime:		&startTime,
		EndTime:		&endTime,
	}
	for {
		response, err := svc.GetRelationalDatabaseLogEvents(ctx, &input)
		if err != nil {
			return err
		}
		for _, e := range response.ResourceLogEvents {
			res <- LogEventWrapper{
				LogEvent:	e,
				LogStreamName:	stream,
			}
		}
		if aws.ToString(response.NextForwardToken) == "" || len(response.ResourceLogEvents) == 0 {
			break
		}
		input.PageToken = response.NextForwardToken
	}
	return nil
}

func (x *TableAwsLightsailDatabaseLogEventsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailDatabaseLogEventsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("database_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_stream_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_lightsail_databases_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_lightsail_databases.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
	}
}

func (x *TableAwsLightsailDatabaseLogEventsGenerator) GetSubTables() []*schema.Table {
	return nil
}
