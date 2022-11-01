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

type TableAwsIamUserAttachedPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamUserAttachedPoliciesGenerator{}

func (x *TableAwsIamUserAttachedPoliciesGenerator) GetTableName() string {
	return "aws_iam_user_attached_policies"
}

func (x *TableAwsIamUserAttachedPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamUserAttachedPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamUserAttachedPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsIamUserAttachedPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config iam.ListAttachedUserPoliciesInput
			p := task.ParentRawResult.(*types.User)
			svc := client.(*aws_client.Client).AwsServices().IAM
			config.UserName = p.UserName
			for {
				output, err := svc.ListAttachedUserPolicies(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.AttachedPolicies
				if output.Marker == nil {
					break
				}
				config.Marker = output.Marker
			}
			return nil
		},
	}
}

func (x *TableAwsIamUserAttachedPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamUserAttachedPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("policy_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_iam_users_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_iam_users.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_arn").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsIamUserAttachedPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
