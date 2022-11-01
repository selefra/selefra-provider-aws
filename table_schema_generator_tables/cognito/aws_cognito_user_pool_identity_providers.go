package cognito

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCognitoUserPoolIdentityProvidersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCognitoUserPoolIdentityProvidersGenerator{}

func (x *TableAwsCognitoUserPoolIdentityProvidersGenerator) GetTableName() string {
	return "aws_cognito_user_pool_identity_providers"
}

func (x *TableAwsCognitoUserPoolIdentityProvidersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCognitoUserPoolIdentityProvidersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCognitoUserPoolIdentityProvidersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsCognitoUserPoolIdentityProvidersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			pool := task.ParentRawResult.(*types.UserPoolType)
			c := client.(*aws_client.Client)
			svc := c.AwsServices().CognitoUserPools

			params := cognitoidentityprovider.ListIdentityProvidersInput{UserPoolId: pool.Id}
			for {
				out, err := svc.ListIdentityProviders(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, out.Providers, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().CognitoUserPools
					item := result.(types.ProviderDescription)
					pool := task.ParentRawResult.(*types.UserPoolType)

					pd, err := svc.DescribeIdentityProvider(ctx, &cognitoidentityprovider.DescribeIdentityProviderInput{
						ProviderName:	item.ProviderName,
						UserPoolId:	pool.Id,
					})
					if err != nil {
						return nil, err
					}
					return pd.IdentityProvider, nil

				})
				if aws.ToString(out.NextToken) == "" {
					break
				}
				params.NextToken = out.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsCognitoUserPoolIdentityProvidersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("cognito-identity")
}

func (x *TableAwsCognitoUserPoolIdentityProvidersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("provider_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_pool_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_pool_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("idp_identifiers").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provider_details").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_cognito_user_pools_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_cognito_user_pools.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attribute_mapping").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provider_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableAwsCognitoUserPoolIdentityProvidersGenerator) GetSubTables() []*schema.Table {
	return nil
}
