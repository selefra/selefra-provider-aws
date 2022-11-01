package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	s3controlTypes "github.com/aws/aws-sdk-go-v2/service/s3control/types"
	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsS3AccountsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsS3AccountsGenerator{}

func (x *TableAwsS3AccountsGenerator) GetTableName() string {
	return "aws_s3_accounts"
}

func (x *TableAwsS3AccountsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsS3AccountsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsS3AccountsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
		},
	}
}

func (x *TableAwsS3AccountsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)

			svc := c.AwsServices().S3Control
			var accountConfig s3control.GetPublicAccessBlockInput
			accountConfig.AccountId = aws.String(c.AccountID)
			resp, err := svc.GetPublicAccessBlock(ctx, &accountConfig)

			if err != nil {

				var nspabc *s3controlTypes.NoSuchPublicAccessBlockConfiguration
				if !errors.As(err, &nspabc) {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- PublicAccessBlockConfigurationWrapper{ConfigExists: false}
			} else {
				resultChannel <- PublicAccessBlockConfigurationWrapper{PublicAccessBlockConfiguration: *resp.PublicAccessBlockConfiguration, ConfigExists: true}
			}

			return nil
		},
	}
}

type PublicAccessBlockConfigurationWrapper struct {
	s3controlTypes.PublicAccessBlockConfiguration
	ConfigExists	bool
}

func (x *TableAwsS3AccountsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsS3AccountsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("block_public_policy").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ignore_public_acls").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restrict_public_buckets").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("config_exists").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("block_public_acls").ColumnType(schema.ColumnTypeBool).Build(),
	}
}

func (x *TableAwsS3AccountsGenerator) GetSubTables() []*schema.Table {
	return nil
}
