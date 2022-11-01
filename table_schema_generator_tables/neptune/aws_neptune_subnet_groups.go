package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsNeptuneSubnetGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsNeptuneSubnetGroupsGenerator{}

func (x *TableAwsNeptuneSubnetGroupsGenerator) GetTableName() string {
	return "aws_neptune_subnet_groups"
}

func (x *TableAwsNeptuneSubnetGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsNeptuneSubnetGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsNeptuneSubnetGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsNeptuneSubnetGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config neptune.DescribeDBSubnetGroupsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Neptune
			for {
				response, err := svc.DescribeDBSubnetGroups(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.DBSubnetGroups
				if aws.ToString(response.Marker) == "" {
					break
				}
				config.Marker = response.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsNeptuneSubnetGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("neptune")
}

func (x *TableAwsNeptuneSubnetGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_subnet_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBSubnetGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnets").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBSubnetGroupArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_subnet_group_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBSubnetGroupDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_group_status").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsNeptuneSubnetGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
