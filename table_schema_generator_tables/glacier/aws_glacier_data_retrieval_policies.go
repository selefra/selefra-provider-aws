package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsGlacierDataRetrievalPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlacierDataRetrievalPoliciesGenerator{}

func (x *TableAwsGlacierDataRetrievalPoliciesGenerator) GetTableName() string {
	return "aws_glacier_data_retrieval_policies"
}

func (x *TableAwsGlacierDataRetrievalPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlacierDataRetrievalPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlacierDataRetrievalPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
		},
	}
}

func (x *TableAwsGlacierDataRetrievalPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Glacier

			response, err := svc.GetDataRetrievalPolicy(ctx, &glacier.GetDataRetrievalPolicyInput{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response.Policy
			return nil
		},
	}
}

func (x *TableAwsGlacierDataRetrievalPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glacier")
}

func (x *TableAwsGlacierDataRetrievalPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsGlacierDataRetrievalPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
