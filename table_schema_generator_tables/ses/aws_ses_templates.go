package ses

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsSesTemplatesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsSesTemplatesGenerator{}

func (x *TableAwsSesTemplatesGenerator) GetTableName() string {
	return "aws_ses_templates"
}

func (x *TableAwsSesTemplatesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsSesTemplatesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsSesTemplatesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsSesTemplatesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*aws_client.Client)
			svc := c.AwsServices().SES

			listInput := new(sesv2.ListEmailTemplatesInput)
			for {
				output, err := svc.ListEmailTemplates(ctx, listInput)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				aws_client.SendResults(resultChannel, output.TemplatesMetadata, func(result any) (any, error) {
					c := client.(*aws_client.Client)
					svc := c.AwsServices().SES
					templateMeta := result.(types.EmailTemplateMetadata)

					getOutput, err := svc.GetEmailTemplate(ctx, &sesv2.GetEmailTemplateInput{TemplateName: templateMeta.TemplateName})
					if err != nil {
						return nil, err
					}
					return &Template{
						TemplateName:		getOutput.TemplateName,
						Text:			getOutput.TemplateContent.Text,
						Html:			getOutput.TemplateContent.Html,
						Subject:		getOutput.TemplateContent.Subject,
						CreatedTimestamp:	templateMeta.CreatedTimestamp,
					}, nil

				})
				if aws.ToString(output.NextToken) == "" {
					break
				}
				listInput.NextToken = output.NextToken
			}

			return nil
		},
	}
}

type Template struct {
	TemplateName	*string

	Html	*string

	Subject	*string

	Text	*string

	CreatedTimestamp	*time.Time
}

func (x *TableAwsSesTemplatesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("email")
}

func (x *TableAwsSesTemplatesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				cl := client.(*aws_client.Client)
				idParts := []string{"template", *result.(*Template).TemplateName}
				return arn.ARN{
					Partition:	cl.Partition,
					Service:	"ses",
					Region:		cl.Region,
					AccountID:	cl.AccountID,
					Resource:	strings.Join(idParts, "/"),
				}.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("text").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("html").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subject").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_timestamp").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsSesTemplatesGenerator) GetSubTables() []*schema.Table {
	return nil
}
