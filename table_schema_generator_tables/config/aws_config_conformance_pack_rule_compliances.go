package config

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsConfigConformancePackRuleCompliancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsConfigConformancePackRuleCompliancesGenerator{}

func (x *TableAwsConfigConformancePackRuleCompliancesGenerator) GetTableName() string {
	return "aws_config_conformance_pack_rule_compliances"
}

func (x *TableAwsConfigConformancePackRuleCompliancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsConfigConformancePackRuleCompliancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsConfigConformancePackRuleCompliancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAwsConfigConformancePackRuleCompliancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			conformancePackDetail := task.ParentRawResult.(types.ConformancePackDetail)
			c := client.(*aws_client.Client)
			cs := c.AwsServices().ConfigService
			params := configservice.DescribeConformancePackComplianceInput{
				ConformancePackName: conformancePackDetail.ConformancePackName,
			}
			for {
				resp, err := cs.DescribeConformancePackCompliance(ctx, &params)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, conformancePackRuleCompliance := range resp.ConformancePackRuleComplianceList {
					detailParams := &configservice.GetConformancePackComplianceDetailsInput{
						ConformancePackName:	conformancePackDetail.ConformancePackName,
						Filters: &types.ConformancePackEvaluationFilters{
							ConfigRuleNames: []string{*conformancePackRuleCompliance.ConfigRuleName},
						},
					}
					for {
						output, err := cs.GetConformancePackComplianceDetails(ctx, detailParams)
						if err != nil {
							return schema.NewDiagnosticsErrorPullTable(task.Table, err)

						}
						for _, conformancePackComplianceDetail := range output.ConformancePackRuleEvaluationResults {
							resultChannel <- ConformancePackComplianceWrapper{
								ComplianceType:			conformancePackRuleCompliance.ComplianceType,
								ConfigRuleName:			conformancePackRuleCompliance.ConfigRuleName,
								Controls:			conformancePackRuleCompliance.Controls,
								ConfigRuleInvokedTime:		conformancePackComplianceDetail.ConfigRuleInvokedTime,
								EvaluationResultIdentifier:	conformancePackComplianceDetail.EvaluationResultIdentifier,
								ResultRecordedTime:		conformancePackComplianceDetail.ResultRecordedTime,
								Annotation:			conformancePackComplianceDetail.Annotation,
							}
						}
						if output.NextToken == nil {
							break
						}
						detailParams.NextToken = output.NextToken
					}
				}
				if resp.NextToken == nil {
					break
				}
				params.NextToken = resp.NextToken
			}
			return nil
		},
	}
}

type ConformancePackComplianceWrapper struct {
	ComplianceType	types.ConformancePackComplianceType

	ConfigRuleName	*string

	Controls	[]string

	ConfigRuleInvokedTime	*time.Time

	EvaluationResultIdentifier	*types.EvaluationResultIdentifier

	ResultRecordedTime	*time.Time

	Annotation	*string
}

func (x *TableAwsConfigConformancePackRuleCompliancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("config")
}

func (x *TableAwsConfigConformancePackRuleCompliancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("annotation").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_config_conformance_packs_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to aws_config_conformance_packs.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conformance_pack_arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("arn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("config_rule_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("controls").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_recorded_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compliance_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("config_rule_invoked_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("evaluation_result_identifier").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAwsConfigConformancePackRuleCompliancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
