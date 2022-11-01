package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsElasticbeanstalkEnvironmentsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsElasticbeanstalkEnvironmentsGenerator{}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetTableName() string {
	return "aws_elasticbeanstalk_environments"
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"id",
		},
	}
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config elasticbeanstalk.DescribeEnvironmentsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().ElasticBeanstalk
			for {
				response, err := svc.DescribeEnvironments(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.Environments
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("elasticbeanstalk")
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("abortable_operation_in_progress").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("application_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cname").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CNAME")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_updated").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operations_role").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("listeners").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					p := result.(types.EnvironmentDescription)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().ElasticBeanstalk
					tagsOutput, err := svc.ListTagsForResource(ctx, &elasticbeanstalk.ListTagsForResourceInput{
						ResourceArn: p.EnvironmentArn,
					}, func(o *elasticbeanstalk.Options) {})
					if err != nil {

						if cl.IsNotFoundError(err) {
							return nil, nil
						}
						return nil, err
					}
					if len(tagsOutput.ResourceTags) == 0 {
						return nil, nil
					}
					tags := make(map[string]*string)
					for _, s := range tagsOutput.ResourceTags {
						tags[*s.Key] = s.Value
					}
					return tags, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tier").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_label").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("date_created").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EnvironmentId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EndpointURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment_links").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("solution_stack_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EnvironmentArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsElasticbeanstalkEnvironmentsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsElasticbeanstalkConfigurationSettingsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsElasticbeanstalkConfigurationOptionsGenerator{}),
	}
}
