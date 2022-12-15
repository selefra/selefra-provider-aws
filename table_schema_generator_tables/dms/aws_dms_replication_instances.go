package dms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsDmsReplicationInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsDmsReplicationInstancesGenerator{}

func (x *TableAwsDmsReplicationInstancesGenerator) GetTableName() string {
	return "aws_dms_replication_instances"
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().DMS

			var describeReplicationInstancesInput *databasemigrationservice.DescribeReplicationInstancesInput
			describeReplicationInstancesOutput, err := svc.DescribeReplicationInstances(ctx, describeReplicationInstancesInput)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if len(describeReplicationInstancesOutput.ReplicationInstances) == 0 {
				return nil
			}

			listTagsForResourceInput := databasemigrationservice.ListTagsForResourceInput{}
			for _, replicationInstance := range describeReplicationInstancesOutput.ReplicationInstances {
				listTagsForResourceInput.ResourceArnList = append(listTagsForResourceInput.ResourceArnList, *replicationInstance.ReplicationInstanceArn)
			}
			var listTagsForResourceOutput *databasemigrationservice.ListTagsForResourceOutput
			listTagsForResourceOutput, err = svc.ListTagsForResource(ctx, &listTagsForResourceInput)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			replicationInstanceTags := make(map[string]map[string]interface{})
			for _, tag := range listTagsForResourceOutput.TagList {
				if replicationInstanceTags[*tag.ResourceArn] == nil {
					replicationInstanceTags[*tag.ResourceArn] = make(map[string]interface{})
				}
				replicationInstanceTags[*tag.ResourceArn][*tag.Key] = *tag.Value
			}

			for _, replicationInstance := range describeReplicationInstancesOutput.ReplicationInstances {
				wrapper := ReplicationInstanceWrapper{
					ReplicationInstance: replicationInstance,
					Tags:                replicationInstanceTags[*replicationInstance.ReplicationInstanceArn],
				}
				resultChannel <- wrapper
			}
			return nil
		},
	}
}

type ReplicationInstanceWrapper struct {
	types.ReplicationInstance
	Tags map[string]interface{}
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("dms")
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allocated_storage").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("engine_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_modified_values").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preferred_maintenance_window").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_instance_private_ip_addresses").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_name_servers").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_create_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_instance_public_ip_addresses").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReplicationInstanceArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_availability_zone").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_instance_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("multi_az").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("MultiAZ")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("publicly_accessible").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_instance_identifier").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_instance_private_ip_address").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_instance_public_ip_address").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("free_until").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_instance_class").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_zone").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_security_groups").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_subnet_group").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_minor_version_upgrade").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsDmsReplicationInstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
