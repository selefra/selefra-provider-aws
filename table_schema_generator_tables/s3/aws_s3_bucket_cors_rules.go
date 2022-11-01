package s3

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsS3BucketCorsRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsS3BucketCorsRulesGenerator{}

func (x *TableAwsS3BucketCorsRulesGenerator) GetTableName() string {
	return "aws_s3_bucket_cors_rules"
}

func (x *TableAwsS3BucketCorsRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsS3BucketCorsRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsS3BucketCorsRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsS3BucketCorsRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			r := task.ParentRawResult.(*WrappedBucket)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().S3
			if task.ParentRow.GetOrDefault("region", nil).(string) == "" {
				return nil
			}
			corsOutput, err := svc.GetBucketCors(ctx, &s3.GetBucketCorsInput{Bucket: r.Name}, func(options *s3.Options) {
				options.Region = task.ParentRow.GetOrDefault("region", nil).(string)
			})
			if err != nil {
				if aws_client.IsAWSError(err, "NoSuchCORSConfiguration", "NoSuchBucket") {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if corsOutput != nil {
				resultChannel <- corsOutput.CORSRules
			}
			return nil
		},
	}
}

type WrappedBucket struct {
	CreationDate	*time.Time

	Name	*string

	ReplicationRole		*string
	ReplicationRules	[]types.ReplicationRule
	Region			string
	LoggingTargetBucket	*string
	LoggingTargetPrefix	*string
	Policy			map[string]interface{}
	VersioningStatus	types.BucketVersioningStatus
	VersioningMfaDelete	types.MFADeleteStatus
	BlockPublicAcls		bool
	BlockPublicPolicy	bool
	IgnorePublicAcls	bool
	RestrictPublicBuckets	bool
	Tags			map[string]*string
	OwnershipControls	[]string
}

func (x *TableAwsS3BucketCorsRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsS3BucketCorsRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_headers").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_age_seconds").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_s3_buckets_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_s3_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_methods").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_origins").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bucket_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expose_headers").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
	}
}

func (x *TableAwsS3BucketCorsRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}
