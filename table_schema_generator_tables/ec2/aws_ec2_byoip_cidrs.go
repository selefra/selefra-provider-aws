package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEc2ByoipCidrsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEc2ByoipCidrsGenerator{}

func (x *TableAwsEc2ByoipCidrsGenerator) GetTableName() string {
	return "aws_ec2_byoip_cidrs"
}

func (x *TableAwsEc2ByoipCidrsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEc2ByoipCidrsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEc2ByoipCidrsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"cidr",
		},
	}
}

func (x *TableAwsEc2ByoipCidrsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			config := ec2.DescribeByoipCidrsInput{
				MaxResults: aws.Int32(100),
			}

			c := client.(*aws_client.Client)

			if _, ok := map[string]struct{}{
				"cn-north-1":		{},
				"cn-northwest-1":	{},
			}[c.Region]; ok {
				return nil
			}
			svc := c.AwsServices().EC2
			for {
				response, err := svc.DescribeByoipCidrs(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- response.ByoipCidrs
				if aws.ToString(response.NextToken) == "" {
					break
				}
				config.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEc2ByoipCidrsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ec2")
}

func (x *TableAwsEc2ByoipCidrsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsEc2ByoipCidrsGenerator) GetSubTables() []*schema.Table {
	return nil
}
