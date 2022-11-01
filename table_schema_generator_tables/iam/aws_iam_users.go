package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIamUsersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamUsersGenerator{}

func (x *TableAwsIamUsersGenerator) GetTableName() string {
	return "aws_iam_users"
}

func (x *TableAwsIamUsersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamUsersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamUsersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
			"account_id",
		},
	}
}

func (x *TableAwsIamUsersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			config := iam.ListUsersInput{}
			c := client.(*aws_client.Client)
			svc := c.AwsServices().IAM
			p := iam.NewListUsersPaginator(svc, &config)
			for p.HasMorePages() {
				response, err := p.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.Users, func(result any) (any, error) {
					listUser := result.(types.User)
					svc := client.(*aws_client.Client).AwsServices().IAM
					userDetail, err := svc.GetUser(ctx, &iam.GetUserInput{
						UserName: aws.String(*listUser.UserName),
					})
					if err != nil {
						return nil, err
					}
					return userDetail.User, nil

				})
			}
			return nil
		},
	}
}

func (x *TableAwsIamUsersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamUsersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("path").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_last_used").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissions_boundary").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UserId")).Build(),
	}
}

func (x *TableAwsIamUsersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAwsIamUserAccessKeysGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsIamUserGroupsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsIamUserAttachedPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAwsIamUserPoliciesGenerator{}),
	}
}
