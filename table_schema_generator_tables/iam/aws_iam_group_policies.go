package iam

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIamGroupPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamGroupPoliciesGenerator{}

func (x *TableAwsIamGroupPoliciesGenerator) GetTableName() string {
	return "aws_iam_group_policies"
}

func (x *TableAwsIamGroupPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamGroupPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamGroupPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsIamGroupPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().IAM
			group := task.ParentRawResult.(types.Group)
			config := iam.ListGroupPoliciesInput{
				GroupName: group.GroupName,
			}
			for {
				output, err := svc.ListGroupPolicies(ctx, &config)
				if err != nil {
					if c.IsNotFoundError(err) {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.PolicyNames, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().IAM
					p := result.(string)
					group := task.ParentRawResult.(types.Group)

					policyResult, err := svc.GetGroupPolicy(ctx, &iam.GetGroupPolicyInput{PolicyName: &p, GroupName: group.GroupName})
					if err != nil {
						return nil, err
					}
					return policyResult, nil

				})
				if aws.ToString(output.Marker) == "" {
					break
				}
				config.Marker = output.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsIamGroupPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamGroupPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("policy_document").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					r := result.(*iam.GetGroupPolicyOutput)

					decodedDocument, err := url.QueryUnescape(*r.PolicyDocument)
					if err != nil {
						return nil, err
					}

					var document map[string]interface{}
					err = json.Unmarshal([]byte(decodedDocument), &document)
					if err != nil {
						return nil, err
					}
					return document, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_iam_groups_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_iam_groups.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
	}
}

func (x *TableAwsIamGroupPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
