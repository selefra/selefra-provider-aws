package route53

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRoute53HealthChecksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRoute53HealthChecksGenerator{}

func (x *TableAwsRoute53HealthChecksGenerator) GetTableName() string {
	return "aws_route53_health_checks"
}

func (x *TableAwsRoute53HealthChecksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRoute53HealthChecksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRoute53HealthChecksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRoute53HealthChecksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config route53.ListHealthChecksInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Route53

			processHealthChecksBundle := func(healthChecks []types.HealthCheck) error {
				tagsCfg := &route53.ListTagsForResourcesInput{ResourceType: types.TagResourceTypeHealthcheck, ResourceIds: make([]string, 0, len(healthChecks))}
				for _, h := range healthChecks {
					tagsCfg.ResourceIds = append(tagsCfg.ResourceIds, *h.Id)
				}
				tagsResponse, err := svc.ListTagsForResources(ctx, tagsCfg)
				if err != nil {
					return err
				}
				for _, h := range healthChecks {
					wrapper := Route53HealthCheckWrapper{
						HealthCheck: h,
						Tags:        aws_client.TagsToMap(getRoute53tagsByResourceID(*h.Id, tagsResponse.ResourceTagSets)),
					}
					resultChannel <- wrapper
				}
				return nil
			}

			for {
				response, err := svc.ListHealthChecks(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				for i := 0; i < len(response.HealthChecks); i += 10 {
					end := i + 10

					if end > len(response.HealthChecks) {
						end = len(response.HealthChecks)
					}
					zones := response.HealthChecks[i:end]
					err := processHealthChecksBundle(zones)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
				}

				if aws.ToString(response.Marker) == "" {
					break
				}
				config.Marker = response.Marker
			}
			return nil
		},
	}
}

type Route53HealthCheckWrapper struct {
	types.HealthCheck
	Tags map[string]string
}

func getRoute53tagsByResourceID(id string, set []types.ResourceTagSet) []types.Tag {
	for _, s := range set {
		if *s.ResourceId == id {
			return s.Tags
		}
	}
	return nil
}

func (x *TableAwsRoute53HealthChecksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsRoute53HealthChecksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				idsComputer := func() ([]string, error) {
					return []string{"healthcheck", *result.(Route53HealthCheckWrapper).Id}, nil
				}

				ids, err := idsComputer()
				if err != nil {
					return nil, diagnostics.AddErrorColumnValueExtractor(task.Table, column, err)
				}

				cl := client.(*aws_client.Client)
				return arn.ARN{
					Partition: cl.Partition,
					Service:   "route53",
					Region:    "",
					AccountID: "",
					Resource:  strings.Join(ids, "/"),
				}.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("`The tags associated with the health check.`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("caller_reference").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("linked_service").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloud_watch_alarm_configuration_dimensions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					r := result.(Route53HealthCheckWrapper)

					if r.CloudWatchAlarmConfiguration == nil {
						return nil, nil
					}
					tags := map[string]*string{}
					for _, t := range r.CloudWatchAlarmConfiguration.Dimensions {
						tags[*t.Name] = t.Value
					}
					return tags, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_check_config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_check_version").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloud_watch_alarm_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsRoute53HealthChecksGenerator) GetSubTables() []*schema.Table {
	return nil
}
