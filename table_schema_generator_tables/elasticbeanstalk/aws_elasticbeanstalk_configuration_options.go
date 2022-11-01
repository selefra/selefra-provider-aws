package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticbeanstalkConfigurationOptionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticbeanstalkConfigurationOptionsGenerator{}

func (x *TableAwsElasticbeanstalkConfigurationOptionsGenerator) GetTableName() string {
	return "aws_elasticbeanstalk_configuration_options"
}

func (x *TableAwsElasticbeanstalkConfigurationOptionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticbeanstalkConfigurationOptionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticbeanstalkConfigurationOptionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsElasticbeanstalkConfigurationOptionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(types.EnvironmentDescription)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().ElasticBeanstalk
			configOptionsIn := elasticbeanstalk.DescribeConfigurationOptionsInput{
				ApplicationName:	p.ApplicationName,
				EnvironmentName:	p.EnvironmentName,
			}
			output, err := svc.DescribeConfigurationOptions(ctx, &configOptionsIn)
			if err != nil {

				if aws_client.IsInvalidParameterValueError(err) {

					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			for _, option := range output.Options {
				resultChannel <- ConfigurationOptionDescriptionWrapper{
					ConfigurationOptionDescription:	option, ApplicationArn: c.ARN("elasticbeanstalk", "application", *p.ApplicationName),
				}
			}

			return nil
		},
	}
}

type ConfigurationOptionDescriptionWrapper struct {
	types.ConfigurationOptionDescription
	ApplicationArn	string
}

func (x *TableAwsElasticbeanstalkConfigurationOptionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticbeanstalk")
}

func (x *TableAwsElasticbeanstalkConfigurationOptionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_value").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_defined").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_elasticbeanstalk_environments_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_elasticbeanstalk_environments.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_value").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_value").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("regex").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("application_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("change_severity").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_length").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value_options").ColumnType(schema.ColumnTypeStringArray).Build(),
	}
}

func (x *TableAwsElasticbeanstalkConfigurationOptionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
