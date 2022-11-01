package shield

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsShieldAttacksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsShieldAttacksGenerator{}

func (x *TableAwsShieldAttacksGenerator) GetTableName() string {
	return "aws_shield_attacks"
}

func (x *TableAwsShieldAttacksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsShieldAttacksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsShieldAttacksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAwsShieldAttacksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Shield
			end := time.Now()
			start := end.Add(-time.Hour * 24)
			config := shield.ListAttacksInput{
				EndTime:	&types.TimeRange{ToExclusive: &end},
				StartTime:	&types.TimeRange{FromInclusive: &start},
			}
			for {
				output, err := svc.ListAttacks(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.AttackSummaries, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().Shield
					a := result.(types.AttackSummary)

					attack, err := svc.DescribeAttack(ctx, &shield.DescribeAttackInput{AttackId: a.AttackId})
					if err != nil {
						return nil, err
					}
					return attack.Attack, nil

				})
				if aws.ToString(output.NextToken) == "" {
					break
				}
				config.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsShieldAttacksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsShieldAttacksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("attack_counters").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attack_properties").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mitigations").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sub_resources").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("`The unique identifier (ID) of the attack`").
			Extractor(column_value_extractor.StructSelector("AttackId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsShieldAttacksGenerator) GetSubTables() []*schema.Table {
	return nil
}
