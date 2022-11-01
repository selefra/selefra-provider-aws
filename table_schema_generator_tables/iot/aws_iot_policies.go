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

type TableAwsIotPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIotPoliciesGenerator{}

func (x *TableAwsIotPoliciesGenerator) GetTableName() string {
	return "aws_iot_policies"
}

func (x *TableAwsIotPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIotPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIotPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsIotPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().IOT
			input := iot.ListPoliciesInput{
				PageSize: aws.Int32(250),
			}

			for {
				response, err := svc.ListPolicies(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				for _, s := range response.Policies {
					profile, err := svc.GetPolicy(ctx, &iot.GetPolicyInput{
						PolicyName: s.PolicyName,
					}, func(options *iot.Options) {
						options.Region = cl.Region
					})
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- profile
				}

				if aws.ToString(response.NextMarker) == "" {
					break
				}
				input.Marker = response.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsIotPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("iot")
}

func (x *TableAwsIotPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PolicyArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsIotPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
