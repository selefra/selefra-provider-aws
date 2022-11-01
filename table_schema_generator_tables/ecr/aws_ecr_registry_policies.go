package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsEcrRegistryPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsEcrRegistryPoliciesGenerator{}

func (x *TableAwsEcrRegistryPoliciesGenerator) GetTableName() string {
	return "aws_ecr_registry_policies"
}

func (x *TableAwsEcrRegistryPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsEcrRegistryPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsEcrRegistryPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"registry_id",
		},
	}
}

func (x *TableAwsEcrRegistryPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().ECR
			output, err := svc.GetRegistryPolicy(ctx, &ecr.GetRegistryPolicyInput{})
			if err != nil {
				if aws_client.IsAWSError(err, "RegistryPolicyNotFoundException") {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output
			return nil
		},
	}
}

func (x *TableAwsEcrRegistryPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("api.ecr")
}

func (x *TableAwsEcrRegistryPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("policy_text").ColumnType(schema.ColumnTypeJSON).
			Extractor(aws_client.MarshaledJsonExtractor("PolicyText")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registry_id").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAwsEcrRegistryPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
