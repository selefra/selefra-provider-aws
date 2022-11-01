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

type TableAwsElasticbeanstalkConfigurationSettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticbeanstalkConfigurationSettingsGenerator{}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetTableName() string {
	return "aws_elasticbeanstalk_configuration_settings"
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			p := task.ParentRawResult.(types.EnvironmentDescription)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().ElasticBeanstalk

			configOptionsIn := elasticbeanstalk.DescribeConfigurationSettingsInput{
				ApplicationName:	p.ApplicationName,
				EnvironmentName:	p.EnvironmentName,
			}
			output, err := svc.DescribeConfigurationSettings(ctx, &configOptionsIn)
			if err != nil {

				if aws_client.IsInvalidParameterValueError(err) {

					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			for _, option := range output.ConfigurationSettings {
				resultChannel <- ConfigurationSettingsDescriptionWrapper{
					ConfigurationSettingsDescription:	option, ApplicationArn: c.ARN("elasticbeanstalk", "application", *p.ApplicationName),
				}
			}

			return nil
		},
	}
}

type ConfigurationSettingsDescriptionWrapper struct {
	types.ConfigurationSettingsDescription
	ApplicationArn	string
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticbeanstalk")
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("option_settings").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("solution_stack_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("application_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_elasticbeanstalk_environments_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_elasticbeanstalk_environments.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_updated").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("application_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_created").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment_name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsElasticbeanstalkConfigurationSettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
