package cloudtrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCloudtrailTrailEventSelectorsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudtrailTrailEventSelectorsGenerator{}

func (x *TableAwsCloudtrailTrailEventSelectorsGenerator) GetTableName() string {
	return "aws_cloudtrail_trail_event_selectors"
}

func (x *TableAwsCloudtrailTrailEventSelectorsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudtrailTrailEventSelectorsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudtrailTrailEventSelectorsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsCloudtrailTrailEventSelectorsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(*CloudTrailWrapper)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudtrail
			response, err := svc.GetEventSelectors(ctx, &cloudtrail.GetEventSelectorsInput{TrailName: r.TrailARN}, func(options *cloudtrail.Options) {
				options.Region = *r.HomeRegion
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response.EventSelectors
			return nil
		},
	}
}

type CloudTrailWrapper struct {
	types.Trail
	Tags	map[string]string
}

func (x *TableAwsCloudtrailTrailEventSelectorsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("cloudtrail")
}

func (x *TableAwsCloudtrailTrailEventSelectorsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_resources").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("exclude_management_event_sources").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("include_management_events").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("read_write_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_cloudtrail_trails_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_cloudtrail_trails.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("trail_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
	}
}

func (x *TableAwsCloudtrailTrailEventSelectorsGenerator) GetSubTables() []*schema.Table {
	return nil
}
