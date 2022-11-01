package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsLightsailContainerServiceImagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailContainerServiceImagesGenerator{}

func (x *TableAwsLightsailContainerServiceImagesGenerator) GetTableName() string {
	return "aws_lightsail_container_service_images"
}

func (x *TableAwsLightsailContainerServiceImagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailContainerServiceImagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailContainerServiceImagesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsLightsailContainerServiceImagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(types.ContainerService)
			input := lightsail.GetContainerImagesInput{
				ServiceName: r.ContainerServiceName,
			}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			deployments, err := svc.GetContainerImages(ctx, &input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- deployments.ContainerImages
			return nil
		},
	}
}

func (x *TableAwsLightsailContainerServiceImagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailContainerServiceImagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("digest").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_lightsail_container_services_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_lightsail_container_services.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("container_service_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
	}
}

func (x *TableAwsLightsailContainerServiceImagesGenerator) GetSubTables() []*schema.Table {
	return nil
}
