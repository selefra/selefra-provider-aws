package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticbeanstalkApplicationVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticbeanstalkApplicationVersionsGenerator{}

func (x *TableAwsElasticbeanstalkApplicationVersionsGenerator) GetTableName() string {
	return "aws_elasticbeanstalk_application_versions"
}

func (x *TableAwsElasticbeanstalkApplicationVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticbeanstalkApplicationVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticbeanstalkApplicationVersionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsElasticbeanstalkApplicationVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config elasticbeanstalk.DescribeApplicationVersionsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().ElasticBeanstalk

			for {
				output, err := svc.DescribeApplicationVersions(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- output.ApplicationVersions

				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}

			return nil
		},
	}
}

func (x *TableAwsElasticbeanstalkApplicationVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticbeanstalk")
}

func (x *TableAwsElasticbeanstalkApplicationVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_created").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_updated").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_label").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApplicationVersionArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("application_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_build_information").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_bundle").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsElasticbeanstalkApplicationVersionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
