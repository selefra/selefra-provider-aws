package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEcsClusterContainerInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEcsClusterContainerInstancesGenerator{}

func (x *TableAwsEcsClusterContainerInstancesGenerator) GetTableName() string {
	return "aws_ecs_cluster_container_instances"
}

func (x *TableAwsEcsClusterContainerInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEcsClusterContainerInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEcsClusterContainerInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsEcsClusterContainerInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cluster := task.ParentRawResult.(types.Cluster)
			region := client.(*aws_client.Client).Region
			svc := client.(*aws_client.Client).AwsServices().ECS
			config := ecs.ListContainerInstancesInput{
				Cluster: cluster.ClusterArn,
			}
			for {
				listContainerInstances, err := svc.ListContainerInstances(ctx, &config, func(o *ecs.Options) {
					o.Region = region
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				if len(listContainerInstances.ContainerInstanceArns) == 0 {
					return nil
				}
				describeServicesInput := ecs.DescribeContainerInstancesInput{
					Cluster:            cluster.ClusterArn,
					ContainerInstances: listContainerInstances.ContainerInstanceArns,
					Include:            []types.ContainerInstanceField{types.ContainerInstanceFieldTags},
				}
				describeContainerInstances, err := svc.DescribeContainerInstances(ctx, &describeServicesInput, func(o *ecs.Options) {
					o.Region = region
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- describeContainerInstances.ContainerInstances

				if listContainerInstances.NextToken == nil {
					break
				}
				config.NextToken = listContainerInstances.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsEcsClusterContainerInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ecs")
}

func (x *TableAwsEcsClusterContainerInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("agent_connected").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("container_instance_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("running_tasks_count").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("agent_update_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attachments").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity_provider_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_reason").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_status").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registered_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("remaining_resources").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_ecs_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_ecs_clusters.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attributes").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ec2_instance_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_tasks_count").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registered_resources").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_info").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsEcsClusterContainerInstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
