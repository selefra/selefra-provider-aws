package s3

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsS3BucketsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsS3BucketsGenerator{}

func (x *TableAwsS3BucketsGenerator) GetTableName() string {
	return "aws_s3_buckets"
}

func (x *TableAwsS3BucketsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsS3BucketsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsS3BucketsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsS3BucketsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().S3
			response, err := svc.ListBuckets(ctx, nil, func(options *s3.Options) {
				options.Region = listBucketRegion(cl)
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			var wg sync.WaitGroup
			buckets := make(chan types.Bucket)
			errs := make(chan error)
			for i := 0; i < fetchS3BucketsPoolSize; i++ {
				wg.Add(1)
				go fetchS3BucketsWorker(ctx, client, buckets, errs, resultChannel, &wg)
			}
			go func() {
				defer close(buckets)
				for _, bucket := range response.Buckets {
					select {
					case <-ctx.Done():
						return
					case buckets <- bucket:
					}
				}
			}()
			done := make(chan struct{})
			go func() {
				for err = range errs {

				}
				close(done)
			}()
			wg.Wait()
			close(errs)
			<-done

			return nil
		},
	}
}

var fetchS3BucketsPoolSize = 10

func fetchS3BucketsWorker(ctx context.Context, client any, buckets <-chan types.Bucket, errs chan<- error, res chan<- interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	cl := client.(*aws_client.Client)
	for bucket := range buckets {
		wb := &WrappedBucket{Name: bucket.Name, CreationDate: bucket.CreationDate}
		err := resolveS3BucketsAttributes(ctx, client, wb)
		if err != nil {
			if !isBucketNotFoundError(cl, err) {
				errs <- err
			}
			continue
		}
		res <- wb
	}
}
func isBucketNotFoundError(cl *aws_client.Client, err error) bool {
	if cl.IsNotFoundError(err) {
		return true
	}
	if err.Error() == "bucket not found" {
		return true
	}
	return false
}

func listBucketRegion(cl *aws_client.Client) string {
	switch cl.Partition {
	case "aws-cn":
		return "cn-north-1"
	case "aws-us-gov":
		return "us-gov-west-1"
	default:
		return "us-east-1"
	}
}
func resolveBucketLogging(ctx context.Context, client any, resource *WrappedBucket, bucketRegion string) error {
	svc := client.(*aws_client.Client).AwsServices().S3
	loggingOutput, err := svc.GetBucketLogging(ctx, &s3.GetBucketLoggingInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		if aws_client.IgnoreAccessDeniedServiceDisabled(err) {

			return nil
		}
		return err
	}
	if loggingOutput.LoggingEnabled == nil {
		return nil
	}
	resource.LoggingTargetBucket = loggingOutput.LoggingEnabled.TargetBucket
	resource.LoggingTargetPrefix = loggingOutput.LoggingEnabled.TargetPrefix
	return nil
}
func resolveBucketOwnershipControls(ctx context.Context, client any, resource *WrappedBucket, bucketRegion string) error {
	c := client.(*aws_client.Client)
	svc := c.AwsServices().S3

	getBucketOwnershipControlOutput, err := svc.GetBucketOwnershipControls(ctx, &s3.GetBucketOwnershipControlsInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})

	if err != nil {

		if aws_client.IsAWSError(err, "OwnershipControlsNotFoundError") {
			return nil
		}

		if aws_client.IgnoreAccessDeniedServiceDisabled(err) {

			return nil
		}

		return err
	}

	if getBucketOwnershipControlOutput == nil {
		return nil
	}

	ownershipControlRules := getBucketOwnershipControlOutput.OwnershipControls.Rules

	if len(ownershipControlRules) == 0 {
		return nil
	}

	stringArray := make([]string, 0, len(ownershipControlRules))

	for _, ownershipControlRule := range ownershipControlRules {
		stringArray = append(stringArray, string(ownershipControlRule.ObjectOwnership))
	}

	resource.OwnershipControls = stringArray
	return nil
}
func resolveBucketPolicy(ctx context.Context, client any, resource *WrappedBucket, bucketRegion string) error {
	c := client.(*aws_client.Client)
	svc := c.AwsServices().S3
	policyOutput, err := svc.GetBucketPolicy(ctx, &s3.GetBucketPolicyInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})

	if err != nil {

		if aws_client.IsAWSError(err, "NoSuchBucketPolicy") {
			return nil
		}
		if aws_client.IgnoreAccessDeniedServiceDisabled(err) {

			return nil
		}
		return err
	}
	if policyOutput == nil || policyOutput.Policy == nil {
		return nil
	}
	var p map[string]interface{}
	err = json.Unmarshal([]byte(*policyOutput.Policy), &p)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON policy: %v", err)
	}
	resource.Policy = p
	return nil
}
func resolveBucketPublicAccessBlock(ctx context.Context, client any, resource *WrappedBucket, bucketRegion string) error {
	c := client.(*aws_client.Client)
	svc := c.AwsServices().S3
	publicAccessOutput, err := svc.GetPublicAccessBlock(ctx, &s3.GetPublicAccessBlockInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {

		if isBucketNotFoundError(c, err) {
			return nil
		}
		if aws_client.IgnoreAccessDeniedServiceDisabled(err) {

			return nil
		}
		return err
	}
	resource.BlockPublicAcls = publicAccessOutput.PublicAccessBlockConfiguration.BlockPublicAcls
	resource.BlockPublicPolicy = publicAccessOutput.PublicAccessBlockConfiguration.BlockPublicPolicy
	resource.IgnorePublicAcls = publicAccessOutput.PublicAccessBlockConfiguration.IgnorePublicAcls
	resource.RestrictPublicBuckets = publicAccessOutput.PublicAccessBlockConfiguration.RestrictPublicBuckets
	return nil
}
func resolveBucketReplication(ctx context.Context, client any, resource *WrappedBucket, bucketRegion string) error {
	c := client.(*aws_client.Client)
	svc := c.AwsServices().S3
	replicationOutput, err := svc.GetBucketReplication(ctx, &s3.GetBucketReplicationInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})

	if err != nil {

		if aws_client.IsAWSError(err, "ReplicationConfigurationNotFoundError") {
			return nil
		}
		if aws_client.IgnoreAccessDeniedServiceDisabled(err) {

			return nil
		}
		return err
	}
	if replicationOutput.ReplicationConfiguration == nil {
		return nil
	}
	resource.ReplicationRole = replicationOutput.ReplicationConfiguration.Role
	resource.ReplicationRules = replicationOutput.ReplicationConfiguration.Rules
	return nil
}
func resolveBucketTagging(ctx context.Context, client any, resource *WrappedBucket, bucketRegion string) error {
	c := client.(*aws_client.Client)
	svc := c.AwsServices().S3
	taggingOutput, err := svc.GetBucketTagging(ctx, &s3.GetBucketTaggingInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {

		if aws_client.IsAWSError(err, "NoSuchTagSet") {
			return nil
		}
		if aws_client.IgnoreAccessDeniedServiceDisabled(err) {

			return nil
		}
		return err
	}
	if taggingOutput == nil {
		return nil
	}
	tags := make(map[string]*string, len(taggingOutput.TagSet))
	for _, t := range taggingOutput.TagSet {
		tags[*t.Key] = t.Value
	}
	resource.Tags = tags
	return nil
}
func resolveBucketVersioning(ctx context.Context, client any, resource *WrappedBucket, bucketRegion string) error {
	c := client.(*aws_client.Client)
	svc := c.AwsServices().S3
	versioningOutput, err := svc.GetBucketVersioning(ctx, &s3.GetBucketVersioningInput{Bucket: resource.Name}, func(options *s3.Options) {
		options.Region = bucketRegion
	})
	if err != nil {
		if aws_client.IgnoreAccessDeniedServiceDisabled(err) {

			return nil
		}
		return err
	}
	resource.VersioningStatus = versioningOutput.Status
	resource.VersioningMfaDelete = versioningOutput.MFADelete
	return nil
}
func resolveS3BucketsAttributes(ctx context.Context, client any, resource *WrappedBucket) error {
	c := client.(*aws_client.Client)
	mgr := c.AwsServices().S3Manager

	output, err := mgr.GetBucketRegion(ctx, *resource.Name)
	if err != nil {
		if isBucketNotFoundError(c, err) {
			return nil
		}
		return err
	}

	resource.Region = "us-east-1"
	if output != "" {
		resource.Region = output
	}
	if err = resolveBucketLogging(ctx, client, resource, resource.Region); err != nil {
		if isBucketNotFoundError(c, err) {
			return nil
		}
		return err
	}

	if err = resolveBucketPolicy(ctx, client, resource, resource.Region); err != nil {
		return err
	}

	if err = resolveBucketVersioning(ctx, client, resource, resource.Region); err != nil {
		return err
	}

	if err = resolveBucketPublicAccessBlock(ctx, client, resource, resource.Region); err != nil {
		return err
	}

	if err = resolveBucketReplication(ctx, client, resource, resource.Region); err != nil {
		return err
	}

	if err = resolveBucketTagging(ctx, client, resource, resource.Region); err != nil {
		return err
	}

	return resolveBucketOwnershipControls(ctx, client, resource, resource.Region)
}

func (x *TableAwsS3BucketsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsS3BucketsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("block_public_policy").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ownership_controls").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logging_target_bucket").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("versioning_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				idsComputer := func() ([]string, error) {
					return []string{*result.(*WrappedBucket).Name}, nil
				}

				ids, err := idsComputer()
				if err != nil {
					return nil, diagnostics.AddErrorColumnValueExtractor(task.Table, column, err)
				}

				cl := client.(*aws_client.Client)
				return arn.ARN{
					Partition:	cl.Partition,
					Service:	"s3",
					Region:		"",
					AccountID:	"",
					Resource:	strings.Join(ids, "/"),
				}.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_role").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logging_target_prefix").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("block_public_acls").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_rules").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("versioning_mfa_delete").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ignore_public_acls").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restrict_public_buckets").ColumnType(schema.ColumnTypeBool).Build(),
	}
}

func (x *TableAwsS3BucketsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsS3BucketGrantsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsS3BucketCorsRulesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsS3BucketEncryptionRulesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsS3BucketLifecyclesGenerator{}),
	}
}
