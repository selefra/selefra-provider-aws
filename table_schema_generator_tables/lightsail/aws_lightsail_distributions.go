package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"golang.org/x/sync/errgroup"
)

type TableAwsLightsailDistributionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsLightsailDistributionsGenerator{}

func (x *TableAwsLightsailDistributionsGenerator) GetTableName() string {
	return "aws_lightsail_distributions"
}

func (x *TableAwsLightsailDistributionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsLightsailDistributionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsLightsailDistributionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsLightsailDistributionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var input lightsail.GetDistributionsInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Lightsail
			for {
				response, err := svc.GetDistributions(ctx, &input, func(options *lightsail.Options) {

					options.Region = "us-east-1"
				})
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				errs, ctx := errgroup.WithContext(ctx)
				errs.SetLimit(MaxGoroutines)
				for _, d := range response.Distributions {
					func(d types.LightsailDistribution) {
						errs.Go(func() error {
							return fetchCacheReset(ctx, resultChannel, c, d)
						})
					}(d)
				}
				err = errs.Wait()
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				if aws.ToString(response.NextPageToken) == "" {
					break
				}
				input.PageToken = response.NextPageToken
			}
			return nil
		},
	}
}

type DistributionWrapper struct {
	*types.LightsailDistribution
	LatestCacheReset	*lightsail.GetDistributionLatestCacheResetOutput
}

func fetchCacheReset(ctx context.Context, res chan<- interface{}, c *aws_client.Client, d types.LightsailDistribution) error {
	svc := c.AwsServices().Lightsail
	resetInput := lightsail.GetDistributionLatestCacheResetInput{
		DistributionName: d.Name,
	}
	resetResp, err := svc.GetDistributionLatestCacheReset(ctx, &resetInput, func(options *lightsail.Options) {

		options.Region = "us-east-1"
	})
	if err != nil && !c.IsNotFoundError(err) {
		return err
	}
	res <- DistributionWrapper{LightsailDistribution: &d, LatestCacheReset: resetResp}
	return nil
}

func (x *TableAwsLightsailDistributionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("lightsail")
}

func (x *TableAwsLightsailDistributionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("support_code").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_behavior_settings").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("origin").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_behaviors").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_cache_reset").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bundle_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_cache_behavior").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("able_to_update_bundle").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alternative_domain_names").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("origin_public_dns").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OriginPublicDNS")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsLightsailDistributionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
