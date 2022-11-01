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

type TableAwsIamSamlIdentityProvidersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamSamlIdentityProvidersGenerator{}

func (x *TableAwsIamSamlIdentityProvidersGenerator) GetTableName() string {
	return "aws_iam_saml_identity_providers"
}

func (x *TableAwsIamSamlIdentityProvidersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamSamlIdentityProvidersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamSamlIdentityProvidersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsIamSamlIdentityProvidersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().IAM
			response, err := svc.ListSAMLProviders(ctx, &iam.ListSAMLProvidersInput{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			aws_client.SendResults(resultChannel, response.SAMLProviderList, func(result any) (any, error) {
				svc := client.(*aws_client.Client).AwsServices().IAM
				p := result.(types.SAMLProviderListEntry)

				providerResponse, err := svc.GetSAMLProvider(ctx, &iam.GetSAMLProviderInput{SAMLProviderArn: p.Arn})
				if err != nil {
					return nil, err
				}
				return IAMSAMLIdentityProviderWrapper{GetSAMLProviderOutput: providerResponse, Arn: *p.Arn}, nil

			})
			return nil
		},
	}
}

type IAMSAMLIdentityProviderWrapper struct {
	*iam.GetSAMLProviderOutput
	Arn	string
}

func (x *TableAwsIamSamlIdentityProvidersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamSamlIdentityProvidersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_until").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsIamSamlIdentityProvidersGenerator) GetSubTables() []*schema.Table {
	return nil
}
