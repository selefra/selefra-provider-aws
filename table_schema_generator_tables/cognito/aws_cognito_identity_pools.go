package cognito

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsCognitoIdentityPoolsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsCognitoIdentityPoolsGenerator{}

func (x *TableAwsCognitoIdentityPoolsGenerator) GetTableName() string {
	return "aws_cognito_identity_pools"
}

func (x *TableAwsCognitoIdentityPoolsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsCognitoIdentityPoolsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsCognitoIdentityPoolsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"id",
		},
	}
}

func (x *TableAwsCognitoIdentityPoolsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().CognitoIdentityPools
			params := cognitoidentity.ListIdentityPoolsInput{

				MaxResults: 60,
			}
			for {
				out, err := svc.ListIdentityPools(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, out.IdentityPools, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().CognitoIdentityPools
					item := result.(types.IdentityPoolShortDescription)

					ipo, err := svc.DescribeIdentityPool(ctx, &cognitoidentity.DescribeIdentityPoolInput{IdentityPoolId: item.IdentityPoolId})
					if err != nil {
						return nil, err
					}
					return ipo, nil

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

func (x *TableAwsCognitoIdentityPoolsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("cognito-identity")
}

func (x *TableAwsCognitoIdentityPoolsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IdentityPoolId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_classic_flow").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cognito_identity_providers").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("developer_provider_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_unauthenticated_identities").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_metadata").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				diagnostics := schema.NewDiagnostics()

				idsComputer := func() ([]string, error) {
					return []string{"identitypool", *result.(*cognitoidentity.DescribeIdentityPoolOutput).IdentityPoolId}, nil
				}

				ids, err := idsComputer()
				if err != nil {
					return nil, diagnostics.AddErrorColumnValueExtractor(task.Table, column, err)
				}

				cl := client.(*aws_client.Client)
				return arn.ARN{
					Partition:	cl.Partition,
					Service:	"cognito-identity",
					Region:		cl.Region,
					AccountID:	cl.AccountID,
					Resource:	strings.Join(ids, "/"),
				}.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_pool_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_pool_tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("open_id_connect_provider_ar_ns").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("OpenIdConnectProviderARNs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("saml_provider_ar_ns").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SamlProviderARNs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_login_providers").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsCognitoIdentityPoolsGenerator) GetSubTables() []*schema.Table {
	return nil
}
