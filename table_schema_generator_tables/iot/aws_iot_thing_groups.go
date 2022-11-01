package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIotThingGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIotThingGroupsGenerator{}

func (x *TableAwsIotThingGroupsGenerator) GetTableName() string {
	return "aws_iot_thing_groups"
}

func (x *TableAwsIotThingGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIotThingGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIotThingGroupsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsIotThingGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			input := iot.ListThingGroupsInput{
				MaxResults: aws.Int32(250),
			}
			c := client.(*aws_client.Client)

			svc := c.AwsServices().IOT
			for {
				response, err := svc.ListThingGroups(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, g := range response.ThingGroups {
					group, err := svc.DescribeThingGroup(ctx, &iot.DescribeThingGroupInput{
						ThingGroupName: g.GroupName,
					}, func(options *iot.Options) {
						options.Region = c.Region
					})
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- group
				}

				if aws.ToString(response.NextToken) == "" {
					break
				}
				input.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsIotThingGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("iot")
}

func (x *TableAwsIotThingGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("thing_group_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("thing_group_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("things_in_group").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					i := result.(*iot.DescribeThingGroupOutput)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().IOT
					input := iot.ListThingsInThingGroupInput{
						ThingGroupName:	i.ThingGroupName,
						MaxResults:	aws.Int32(250),
					}

					var things []string
					for {
						response, err := svc.ListThingsInThingGroup(ctx, &input)
						if err != nil {
							return nil, err
						}

						things = append(things, response.Things...)

						if aws.ToString(response.NextToken) == "" {
							break
						}
						input.NextToken = response.NextToken
					}
					return things, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ThingGroupArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("index_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_string").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("thing_group_properties").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policies").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					i := result.(*iot.DescribeThingGroupOutput)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().IOT
					input := iot.ListAttachedPoliciesInput{
						Target:		i.ThingGroupArn,
						PageSize:	aws.Int32(250),
					}

					var policies []string
					for {
						response, err := svc.ListAttachedPolicies(ctx, &input)
						if err != nil {
							return nil, err
						}

						for _, p := range response.Policies {
							policies = append(policies, *p.PolicyArn)
						}

						if aws.ToString(response.NextMarker) == "" {
							break
						}
						input.Marker = response.NextMarker
					}
					return policies, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("query_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("thing_group_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsIotThingGroupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
