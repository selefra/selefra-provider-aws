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

type TableAwsIamRolePoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamRolePoliciesGenerator{}

func (x *TableAwsIamRolePoliciesGenerator) GetTableName() string {
	return "aws_iam_role_policies"
}

func (x *TableAwsIamRolePoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamRolePoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamRolePoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsIamRolePoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().IAM
			role := task.ParentRawResult.(*types.Role)
			config := iam.ListRolePoliciesInput{
				RoleName: role.RoleName,
			}
			for {
				output, err := svc.ListRolePolicies(ctx, &config)
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
					role := task.ParentRawResult.(*types.Role)

					policyResult, err := svc.GetRolePolicy(ctx, &iam.GetRolePolicyInput{PolicyName: &p, RoleName: role.RoleName})
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

func (x *TableAwsIamRolePoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamRolePoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("policy_document").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					r := result.(*iam.GetRolePolicyOutput)

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
		table_schema_generator.NewColumnBuilder().ColumnName("policy_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_iam_roles_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_iam_roles.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
	}
}

func (x *TableAwsIamRolePoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
