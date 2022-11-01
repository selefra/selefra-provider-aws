package cloudtrail

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCloudtrailTrailsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCloudtrailTrailsGenerator{}

func (x *TableAwsCloudtrailTrailsGenerator) GetTableName() string {
	return "aws_cloudtrail_trails"
}

func (x *TableAwsCloudtrailTrailsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCloudtrailTrailsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCloudtrailTrailsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsCloudtrailTrailsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Cloudtrail

			response, err := svc.DescribeTrails(ctx, nil)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			getBundledTrailsWithTags := func(trails []types.Trail, region string) ([]*CloudTrailWrapper, error) {
				processed := make([]*CloudTrailWrapper, len(trails))

				input := cloudtrail.ListTagsInput{
					ResourceIdList: make([]string, 0, len(trails)),
				}

				for i, h := range trails {
					processed[i] = &CloudTrailWrapper{
						Trail:	h,
						Tags:	make(map[string]string),
					}

					arnParts, err := arn.Parse(*h.TrailARN)
					if err != nil {

						continue
					}
					if aws.ToBool(h.IsOrganizationTrail) && c.AccountID != arnParts.AccountID {

						continue
					}

					input.ResourceIdList = append(input.ResourceIdList, *h.TrailARN)
				}

				if len(input.ResourceIdList) == 0 {
					return processed, nil
				}

				for {
					response, err := svc.ListTags(ctx, &input, func(options *cloudtrail.Options) {
						options.Region = region
					})
					if err != nil {
						return nil, err
					}
					for i, tr := range processed {
						aws_client.TagsIntoMap(getCloudTrailTagsByResourceID(*tr.TrailARN, response.ResourceTagList), processed[i].Tags)
					}
					if aws.ToString(response.NextToken) == "" {
						break
					}
					input.NextToken = response.NextToken
				}

				return processed, nil
			}

			aggregatedTrails, err := aggregateCloudTrails(response.TrailList)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			for region, trails := range aggregatedTrails {
				for i := 0; i < len(trails); i += 20 {
					end := i + 20

					if end > len(trails) {
						end = len(trails)
					}
					t := trails[i:end]
					processed, err := getBundledTrailsWithTags(t, region)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- processed
				}
			}

			return nil
		},
	}
}

func aggregateCloudTrails(trails []types.Trail) (map[string][]types.Trail, error) {
	resp := make(map[string][]types.Trail)
	for _, t := range trails {
		if t.HomeRegion == nil {
			return nil, fmt.Errorf("got cloudtrail with HomeRegion == nil")
		}
		resp[*t.HomeRegion] = append(resp[*t.HomeRegion], t)
	}
	return resp, nil
}
func getCloudTrailTagsByResourceID(id string, set []types.ResourceTag) []types.Tag {
	for _, s := range set {
		if *s.ResourceId == id {
			return s.TagsList
		}
	}
	return nil
}

func (x *TableAwsCloudtrailTrailsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("cloudtrail")
}

func (x *TableAwsCloudtrailTrailsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("s3_key_prefix").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sns_topic_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("include_global_service_events").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloudwatch_logs_log_group_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					groupName := ""
					r := result.(*CloudTrailWrapper)
					if r.CloudWatchLogsLogGroupArn != nil {
						matches := regexp.MustCompile("arn:[a-zA-Z0-9-]+:logs:[a-z0-9-]+:[0-9]+:log-group:([a-zA-Z0-9-/]+):").
							FindStringSubmatch(*r.CloudWatchLogsLogGroupArn)
						if len(matches) < 2 {

						} else {
							groupName = matches[1]
						}
					}

					return groupName, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Cloudtrail
					r := result.(*CloudTrailWrapper)
					response, err := svc.GetTrailStatus(ctx,
						&cloudtrail.GetTrailStatusInput{Name: r.TrailARN}, func(o *cloudtrail.Options) {
							o.Region = *r.HomeRegion
						})
					if err != nil {
						return nil, err
					}
					return response, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloud_watch_logs_role_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_custom_event_selectors").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_insight_selectors").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_file_validation_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("s3_bucket_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloud_watch_logs_log_group_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("home_region").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_organization_trail").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sns_topic_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SnsTopicARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TrailARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_multi_region_trail").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsCloudtrailTrailsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsCloudtrailTrailEventSelectorsGenerator{}),
	}
}
