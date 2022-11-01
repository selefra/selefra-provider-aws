package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticbeanstalkApplicationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticbeanstalkApplicationsGenerator{}

func (x *TableAwsElasticbeanstalkApplicationsGenerator) GetTableName() string {
	return "aws_elasticbeanstalk_applications"
}

func (x *TableAwsElasticbeanstalkApplicationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticbeanstalkApplicationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticbeanstalkApplicationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
			"date_created",
		},
	}
}

func (x *TableAwsElasticbeanstalkApplicationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config elasticbeanstalk.DescribeApplicationsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().ElasticBeanstalk
			output, err := svc.DescribeApplications(ctx, &config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output.Applications
			return nil
		},
	}
}

func (x *TableAwsElasticbeanstalkApplicationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticbeanstalk")
}

func (x *TableAwsElasticbeanstalkApplicationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_created").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration_templates").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("versions").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApplicationArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("application_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_updated").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_lifecycle_config").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsElasticbeanstalkApplicationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
