package iam

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIamUserAccessKeysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamUserAccessKeysGenerator{}

func (x *TableAwsIamUserAccessKeysGenerator) GetTableName() string {
	return "aws_iam_user_access_keys"
}

func (x *TableAwsIamUserAccessKeysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamUserAccessKeysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamUserAccessKeysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsIamUserAccessKeysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var config iam.ListAccessKeysInput
			p := task.ParentRawResult.(*types.User)
			svc := client.(*aws_client.Client).AwsServices().IAM
			config.UserName = p.UserName
			for {
				output, err := svc.ListAccessKeys(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				keys := make([]AccessKeyWrapper, len(output.AccessKeyMetadata))
				for i, key := range output.AccessKeyMetadata {
					switch i {
					case 0:
						rotated := task.ParentRow.GetOrDefault("access_key_1_last_rotated", nil)
						if rotated != nil {
							keys[i] = AccessKeyWrapper{AccessKeyMetadata: key, LastRotated: rotated.(time.Time)}
						} else {
							keys[i] = AccessKeyWrapper{AccessKeyMetadata: key, LastRotated: *key.CreateDate}
						}
					case 1:
						rotated := task.ParentRow.GetOrDefault("access_key_2_last_rotated", nil)
						if rotated != nil {
							keys[i] = AccessKeyWrapper{AccessKeyMetadata: key, LastRotated: rotated.(time.Time)}
						} else {
							keys[i] = AccessKeyWrapper{AccessKeyMetadata: key, LastRotated: *key.CreateDate}
						}
					default:
						keys[i] = AccessKeyWrapper{AccessKeyMetadata: key}
					}
				}
				resultChannel <- keys
				if output.Marker == nil {
					break
				}
				config.Marker = output.Marker
			}
			return nil
		},
	}
}

type AccessKeyWrapper struct {
	types.AccessKeyMetadata
	LastRotated	time.Time
}

func (x *TableAwsIamUserAccessKeysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamUserAccessKeysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("last_rotated").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_used_service_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_used").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_iam_users_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_iam_users.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsIamUserAccessKeysGenerator) GetSubTables() []*schema.Table {
	return nil
}
