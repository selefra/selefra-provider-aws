package mq

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsMqBrokerUsersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsMqBrokerUsersGenerator{}

func (x *TableAwsMqBrokerUsersGenerator) GetTableName() string {
	return "aws_mq_broker_users"
}

func (x *TableAwsMqBrokerUsersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsMqBrokerUsersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsMqBrokerUsersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsMqBrokerUsersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			broker := task.ParentRawResult.(*mq.DescribeBrokerOutput)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().MQ
			for _, us := range broker.Users {
				input := mq.DescribeUserInput{
					BrokerId:	broker.BrokerId,
					Username:	us.Username,
				}
				output, err := svc.DescribeUser(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output
			}
			return nil
		},
	}
}

func (x *TableAwsMqBrokerUsersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("mq")
}

func (x *TableAwsMqBrokerUsersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("pending").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("username").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("broker_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("broker_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("groups").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("console_access").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_mq_brokers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_mq_brokers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsMqBrokerUsersGenerator) GetSubTables() []*schema.Table {
	return nil
}
