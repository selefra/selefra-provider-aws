package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsS3BucketEncryptionRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsS3BucketEncryptionRulesGenerator{}

func (x *TableAwsS3BucketEncryptionRulesGenerator) GetTableName() string {
	return "aws_s3_bucket_encryption_rules"
}

func (x *TableAwsS3BucketEncryptionRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsS3BucketEncryptionRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsS3BucketEncryptionRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsS3BucketEncryptionRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(*WrappedBucket)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().S3
			if task.ParentRow.GetOrDefault("region", nil).(string) == "" {
				return nil
			}
			aclOutput, err := svc.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{Bucket: r.Name}, func(options *s3.Options) {
				options.Region = task.ParentRow.GetOrDefault("region", nil).(string)
			})
			if err != nil {
				if aws_client.IsAWSError(err, "ServerSideEncryptionConfigurationNotFoundError") {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- aclOutput.ServerSideEncryptionConfiguration.Rules
			return nil
		},
	}
}

func (x *TableAwsS3BucketEncryptionRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsS3BucketEncryptionRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("apply_server_side_encryption_by_default").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bucket_key_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_s3_buckets_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_s3_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bucket_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
	}
}

func (x *TableAwsS3BucketEncryptionRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}
