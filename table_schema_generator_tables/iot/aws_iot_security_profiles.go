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

type TableAwsIotSecurityProfilesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIotSecurityProfilesGenerator{}

func (x *TableAwsIotSecurityProfilesGenerator) GetTableName() string {
	return "aws_iot_security_profiles"
}

func (x *TableAwsIotSecurityProfilesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIotSecurityProfilesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIotSecurityProfilesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsIotSecurityProfilesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().IOT
			input := iot.ListSecurityProfilesInput{
				MaxResults: aws.Int32(250),
			}

			for {
				response, err := svc.ListSecurityProfiles(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				for _, s := range response.SecurityProfileIdentifiers {
					profile, err := svc.DescribeSecurityProfile(ctx, &iot.DescribeSecurityProfileInput{
						SecurityProfileName: s.Name,
					}, func(options *iot.Options) {
						options.Region = cl.Region
					})
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- profile
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

func (x *TableAwsIotSecurityProfilesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("iot")
}

func (x *TableAwsIotSecurityProfilesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("targets").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					i := result.(*iot.DescribeSecurityProfileOutput)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().IOT
					input := iot.ListTargetsForSecurityProfileInput{
						SecurityProfileName: i.SecurityProfileName,
						MaxResults:          aws.Int32(250),
					}

					var targets []string
					for {
						response, err := svc.ListTargetsForSecurityProfile(ctx, &input)
						if err != nil {
							return nil, err
						}

						for _, t := range response.SecurityProfileTargets {
							targets = append(targets, *t.Arn)
						}

						if aws.ToString(response.NextToken) == "" {
							break
						}
						input.NextToken = response.NextToken
					}
					return targets, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_metrics_to_retain").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_profile_description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_profile_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SecurityProfileArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("additional_metrics_to_retain_v2").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alert_targets").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("behaviors").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsIotSecurityProfilesGenerator) GetSubTables() []*schema.Table {
	return nil
}
