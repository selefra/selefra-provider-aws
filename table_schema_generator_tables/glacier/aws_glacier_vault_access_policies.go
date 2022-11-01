package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsGlacierVaultAccessPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlacierVaultAccessPoliciesGenerator{}

func (x *TableAwsGlacierVaultAccessPoliciesGenerator) GetTableName() string {
	return "aws_glacier_vault_access_policies"
}

func (x *TableAwsGlacierVaultAccessPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlacierVaultAccessPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlacierVaultAccessPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"vault_arn",
		},
	}
}

func (x *TableAwsGlacierVaultAccessPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Glacier
			p := task.ParentRawResult.(types.DescribeVaultOutput)

			response, err := svc.GetVaultAccessPolicy(ctx, &glacier.GetVaultAccessPolicyInput{
				VaultName: p.VaultName,
			})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response.Policy
			return nil
		},
	}
}

func (x *TableAwsGlacierVaultAccessPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glacier")
}

func (x *TableAwsGlacierVaultAccessPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vault_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(aws_client.MarshaledJsonExtractor("Policy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_glacier_vaults_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_glacier_vaults.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAwsGlacierVaultAccessPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
