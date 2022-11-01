package route53

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsRoute53HostedZonesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsRoute53HostedZonesGenerator{}

func (x *TableAwsRoute53HostedZonesGenerator) GetTableName() string {
	return "aws_route53_hosted_zones"
}

func (x *TableAwsRoute53HostedZonesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsRoute53HostedZonesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsRoute53HostedZonesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsRoute53HostedZonesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config route53.ListHostedZonesInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Route53

			processHostedZonesBundle := func(hostedZones []types.HostedZone) error {
				tagsCfg := &route53.ListTagsForResourcesInput{ResourceType: types.TagResourceTypeHostedzone, ResourceIds: make([]string, 0, len(hostedZones))}
				for i := range hostedZones {
					parsedId := strings.Replace(*hostedZones[i].Id, fmt.Sprintf("/%s/", types.TagResourceTypeHostedzone), "", 1)
					hostedZones[i].Id = &parsedId
					tagsCfg.ResourceIds = append(tagsCfg.ResourceIds, parsedId)
				}
				tagsResponse, err := svc.ListTagsForResources(ctx, tagsCfg)
				if err != nil {
					return err
				}
				for _, h := range hostedZones {
					gotHostedZone, err := svc.GetHostedZone(ctx, &route53.GetHostedZoneInput{Id: h.Id})
					if err != nil {
						return err
					}
					var delegationSetId *string
					if gotHostedZone.DelegationSet != nil {
						delegationSetId = gotHostedZone.DelegationSet.Id
					}
					resultChannel <- &Route53HostedZoneWrapper{
						HostedZone:		h,
						Tags:			aws_client.TagsToMap(getRoute53tagsByResourceID(*h.Id, tagsResponse.ResourceTagSets)),
						DelegationSetId:	delegationSetId,
						VPCs:			gotHostedZone.VPCs,
					}
				}
				return nil
			}

			for {
				response, err := svc.ListHostedZones(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				for i := 0; i < len(response.HostedZones); i += 10 {
					end := i + 10

					if end > len(response.HostedZones) {
						end = len(response.HostedZones)
					}
					zones := response.HostedZones[i:end]
					err := processHostedZonesBundle(zones)
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

func (x *TableAwsRoute53HostedZonesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsRoute53HostedZonesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("config").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("linked_service").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delegation_set_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					cl := client.(*aws_client.Client)
					hz := result.(*Route53HostedZoneWrapper)
					return cl.PartitionGlobalARN("route53", "hostedzone", *hz.Id), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("caller_reference").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_record_set_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpcs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VPCs")).Build(),
	}
}

func (x *TableAwsRoute53HostedZonesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsRoute53HostedZoneQueryLoggingConfigsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsRoute53HostedZoneResourceRecordSetsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsRoute53HostedZoneTrafficPolicyInstancesGenerator{}),
	}
}
