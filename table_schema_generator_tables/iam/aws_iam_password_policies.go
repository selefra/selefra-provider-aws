package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIamPasswordPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamPasswordPoliciesGenerator{}

func (x *TableAwsIamPasswordPoliciesGenerator) GetTableName() string {
	return "aws_iam_password_policies"
}

func (x *TableAwsIamPasswordPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamPasswordPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamPasswordPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
		},
	}
}

func (x *TableAwsIamPasswordPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config iam.GetAccountPasswordPolicyInput
			c := client.(*aws_client.Client)
			svc := c.AwsServices().IAM
			response, err := svc.GetAccountPasswordPolicy(ctx, &config)
			if err != nil {
				if c.IsNotFoundError(err) {
					resultChannel <- PasswordPolicyWrapper{PolicyExists: false}
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- PasswordPolicyWrapper{PasswordPolicy: *response.PasswordPolicy, PolicyExists: true}
			return nil
		},
	}
}

type PasswordPolicyWrapper struct {
	types.PasswordPolicy
	PolicyExists	bool
}

func (x *TableAwsIamPasswordPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamPasswordPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("require_lowercase_characters").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_numbers").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_uppercase_characters").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_users_to_change_password").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum_password_length").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_reuse_prevention").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_symbols").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_exists").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expire_passwords").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hard_expiry").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_password_age").ColumnType(schema.ColumnTypeInt).Build(),
	}
}

func (x *TableAwsIamPasswordPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
