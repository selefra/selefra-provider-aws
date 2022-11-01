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

type TableAwsIamUserPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamUserPoliciesGenerator{}

func (x *TableAwsIamUserPoliciesGenerator) GetTableName() string {
	return "aws_iam_user_policies"
}

func (x *TableAwsIamUserPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamUserPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamUserPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsIamUserPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().IAM
			user := task.ParentRawResult.(*types.User)
			config := iam.ListUserPoliciesInput{UserName: user.UserName}
			for {
				output, err := svc.ListUserPolicies(ctx, &config)
				if err != nil {
					if c.IsNotFoundError(err) {
						return nil
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.PolicyNames, func(result any) (any, error) {
					svc := client.(*aws_client.Client).AwsServices().IAM
					p := result.(string)
					user := task.ParentRawResult.(*types.User)

					policyResult, err := svc.GetUserPolicy(ctx, &iam.GetUserPolicyInput{PolicyName: &p, UserName: user.UserName})
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

func (x *TableAwsIamUserPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamUserPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("policy_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_iam_users_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_iam_users.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_document").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					r := result.(*iam.GetUserPolicyOutput)

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
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsIamUserPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
