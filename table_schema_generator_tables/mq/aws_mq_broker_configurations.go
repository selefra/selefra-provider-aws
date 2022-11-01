package mq

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsMqBrokerConfigurationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsMqBrokerConfigurationsGenerator{}

func (x *TableAwsMqBrokerConfigurationsGenerator) GetTableName() string {
	return "aws_mq_broker_configurations"
}

func (x *TableAwsMqBrokerConfigurationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsMqBrokerConfigurationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsMqBrokerConfigurationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsMqBrokerConfigurationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			broker := task.ParentRawResult.(*mq.DescribeBrokerOutput)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().MQ

			if broker.Configurations == nil {
				return nil
			}

			list := broker.Configurations.History
			if broker.Configurations.Current != nil {
				list = append(list, *broker.Configurations.Current)
			}

			dupes := make(map[string]struct{}, len(list))
			configurations := make([]mq.DescribeConfigurationOutput, 0, len(list))
			for _, cfg := range list {
				if cfg.Id == nil {
					continue
				}

				if _, ok := dupes[*cfg.Id]; ok {
					continue
				}
				dupes[*cfg.Id] = struct{}{}

				input := mq.DescribeConfigurationInput{ConfigurationId: cfg.Id}
				output, err := svc.DescribeConfiguration(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				configurations = append(configurations, *output)
			}
			resultChannel <- configurations
			return nil
		},
	}
}

func (x *TableAwsMqBrokerConfigurationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("mq")
}

func (x *TableAwsMqBrokerConfigurationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_mq_brokers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_mq_brokers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authentication_strategy").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("broker_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_revision").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_type").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsMqBrokerConfigurationsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsMqBrokerConfigurationRevisionsGenerator{}),
	}
}
