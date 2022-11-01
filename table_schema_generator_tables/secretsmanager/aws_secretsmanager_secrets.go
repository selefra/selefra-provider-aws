package secretsmanager

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSecretsmanagerSecretsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSecretsmanagerSecretsGenerator{}

func (x *TableAwsSecretsmanagerSecretsGenerator) GetTableName() string {
	return "aws_secretsmanager_secrets"
}

func (x *TableAwsSecretsmanagerSecretsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSecretsmanagerSecretsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSecretsmanagerSecretsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSecretsmanagerSecretsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().SecretsManager
			cfg := secretsmanager.ListSecretsInput{}
			for {
				response, err := svc.ListSecrets(ctx, &cfg)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, response.SecretList, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().SecretsManager
					n := result.(types.SecretListEntry)

					resp, err := svc.DescribeSecret(ctx, &secretsmanager.DescribeSecretInput{
						SecretId: n.ARN,
					})
					if err != nil {
						return nil, err
					}
					return resp, nil

				})
				if aws.ToString(response.NextToken) == "" {
					break
				}
				cfg.NextToken = response.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsSecretsmanagerSecretsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("secretsmanager")
}

func (x *TableAwsSecretsmanagerSecretsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("policy").ColumnType(schema.ColumnTypeJSON).Description("`A JSON-formatted string that describes the permissions that are associated with the attached secret.`").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					r := result.(*secretsmanager.DescribeSecretOutput)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().SecretsManager
					cfg := secretsmanager.GetResourcePolicyInput{
						SecretId: r.ARN,
					}
					response, err := svc.GetResourcePolicy(ctx, &cfg)
					if err != nil {
						return nil, err
					}

					if response.ResourcePolicy == nil {
						return nil, nil
					}

					v := map[string]interface{}{}
					err = json.Unmarshal([]byte(*response.ResourcePolicy), &v)
					if err != nil {
						return nil, err
					}
					return v, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owning_service").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("`The list of user-defined tags associated with the secret`").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_accessed_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_changed_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_rotated_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_ids_to_stages").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rotation_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rotation_lambda_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RotationLambdaARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rotation_rules").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("primary_region").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_status").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ARN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
	}
}

func (x *TableAwsSecretsmanagerSecretsGenerator) GetSubTables() []*schema.Table {
	return nil
}
