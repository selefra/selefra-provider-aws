package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRdsDbSecurityGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRdsDbSecurityGroupsGenerator{}

func (x *TableAwsRdsDbSecurityGroupsGenerator) GetTableName() string {
	return "aws_rds_db_security_groups"
}

func (x *TableAwsRdsDbSecurityGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRdsDbSecurityGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRdsDbSecurityGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRdsDbSecurityGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().RDS
			var input rds.DescribeDBSecurityGroupsInput
			for {
				output, err := svc.DescribeDBSecurityGroups(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.DBSecurityGroups
				if aws.ToString(output.Marker) == "" {
					break
				}
				input.Marker = output.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsRdsDbSecurityGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("rds")
}

func (x *TableAwsRdsDbSecurityGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_security_group_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBSecurityGroupDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_ranges").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IPRanges")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBSecurityGroupArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_security_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DBSecurityGroupName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ec2_security_groups").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EC2SecurityGroups")).Build(),
	}
}

func (x *TableAwsRdsDbSecurityGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
