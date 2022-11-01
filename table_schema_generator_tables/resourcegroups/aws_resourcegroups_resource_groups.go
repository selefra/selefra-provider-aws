package resourcegroups

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsResourcegroupsResourceGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsResourcegroupsResourceGroupsGenerator{}

func (x *TableAwsResourcegroupsResourceGroupsGenerator) GetTableName() string {
	return "aws_resourcegroups_resource_groups"
}

func (x *TableAwsResourcegroupsResourceGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsResourcegroupsResourceGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsResourcegroupsResourceGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsResourcegroupsResourceGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config resourcegroups.ListGroupsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().ResourceGroups
			for {
				output, err := svc.ListGroups(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.GroupIdentifiers, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					group := result.(types.GroupIdentifier)
					svc := c.AwsServices().ResourceGroups
					groupResponse, err := svc.GetGroup(ctx, &resourcegroups.GetGroupInput{
						Group: group.GroupArn,
					})
					if err != nil {
						return nil, err
					}

					input := resourcegroups.GetGroupQueryInput{
						Group: groupResponse.Group.GroupArn,
					}
					output, err := svc.GetGroupQuery(ctx, &input)
					if err != nil {
						return nil, err
					}
					return ResourceGroupWrapper{
						Group:		groupResponse.Group,
						ResourceQuery:	output.GroupQuery.ResourceQuery,
					}, nil

				})
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

type ResourceGroupWrapper struct {
	*types.Group
	*types.ResourceQuery
}

func (x *TableAwsResourcegroupsResourceGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("resource-groups")
}

func (x *TableAwsResourcegroupsResourceGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("query").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GroupArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsResourcegroupsResourceGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
