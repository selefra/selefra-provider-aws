package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsGlueClassifiersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsGlueClassifiersGenerator{}

func (x *TableAwsGlueClassifiersGenerator) GetTableName() string {
	return "aws_glue_classifiers"
}

func (x *TableAwsGlueClassifiersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsGlueClassifiersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsGlueClassifiersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
			"region",
			"name",
		},
	}
}

func (x *TableAwsGlueClassifiersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().Glue
			input := glue.GetClassifiersInput{}
			for {
				output, err := svc.GetClassifiers(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Classifiers

				if aws.ToString(output.NextToken) == "" {
					break
				}
				input.NextToken = output.NextToken
			}
			return nil
		},
	}
}

func (x *TableAwsGlueClassifiersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("glue")
}

func (x *TableAwsGlueClassifiersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("csv_classifier").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("grok_classifier").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("json_classifier").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("xml_classifier").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("XMLClassifier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					r := result.(types.Classifier)
					if r.CsvClassifier != nil {
						return r.CsvClassifier.Name, nil
					}
					if r.JsonClassifier != nil {
						return r.JsonClassifier.Name, nil
					}
					if r.GrokClassifier != nil {
						return r.GrokClassifier.Name, nil
					}
					if r.XMLClassifier != nil {
						return r.XMLClassifier.Name, nil
					}
					return nil, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
	}
}

func (x *TableAwsGlueClassifiersGenerator) GetSubTables() []*schema.Table {
	return nil
}
