package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEc2RegionalConfigGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEc2RegionalConfigGenerator{}

func (x *TableAwsEc2RegionalConfigGenerator) GetTableName() string {
	return "aws_ec2_regional_config"
}

func (x *TableAwsEc2RegionalConfigGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEc2RegionalConfigGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEc2RegionalConfigGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
		},
	}
}

func (x *TableAwsEc2RegionalConfigGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)

			svc := c.AwsServices().EC2
			var regionalConfig RegionalConfig
			resp, err := svc.GetEbsDefaultKmsKeyId(ctx, &ec2.GetEbsDefaultKmsKeyIdInput{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			regionalConfig.EbsDefaultKmsKeyId = resp.KmsKeyId

			ebsResp, err := svc.GetEbsEncryptionByDefault(ctx, &ec2.GetEbsEncryptionByDefaultInput{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			if ebsResp.EbsEncryptionByDefault != nil {
				regionalConfig.EbsEncryptionEnabledByDefault = *ebsResp.EbsEncryptionByDefault
			}
			resultChannel <- regionalConfig
			return nil
		},
	}
}

type RegionalConfig struct {
	EbsEncryptionEnabledByDefault	bool
	EbsDefaultKmsKeyId		*string
}

func (x *TableAwsEc2RegionalConfigGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("ec2")
}

func (x *TableAwsEc2RegionalConfigGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ebs_encryption_enabled_by_default").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ebs_default_kms_key_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsEc2RegionalConfigGenerator) GetSubTables() []*schema.Table {
	return nil
}
