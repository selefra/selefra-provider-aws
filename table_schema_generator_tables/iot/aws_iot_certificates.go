package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIotCertificatesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIotCertificatesGenerator{}

func (x *TableAwsIotCertificatesGenerator) GetTableName() string {
	return "aws_iot_certificates"
}

func (x *TableAwsIotCertificatesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIotCertificatesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIotCertificatesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
		},
	}
}

func (x *TableAwsIotCertificatesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			cl := client.(*aws_client.Client)
			svc := cl.AwsServices().IOT
			input := iot.ListCertificatesInput{
				PageSize: aws.Int32(250),
			}

			for {
				response, err := svc.ListCertificates(ctx, &input)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				for _, ct := range response.Certificates {
					cert, err := svc.DescribeCertificate(ctx, &iot.DescribeCertificateInput{
						CertificateId: ct.CertificateId,
					}, func(options *iot.Options) {
						options.Region = cl.Region
					})
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- cert.CertificateDescription
				}

				if aws.ToString(response.NextMarker) == "" {
					break
				}
				input.Marker = response.NextMarker
			}
			return nil
		},
	}
}

func (x *TableAwsIotCertificatesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartitionAndRegion("iot")
}

func (x *TableAwsIotCertificatesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("previous_owned_by").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsRegionIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policies").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					i := result.(*types.CertificateDescription)
					cl := client.(*aws_client.Client)
					svc := cl.AwsServices().IOT
					input := iot.ListAttachedPoliciesInput{
						Target:		i.CertificateArn,
						PageSize:	aws.Int32(250),
					}

					var policies []string
					for {
						response, err := svc.ListAttachedPolicies(ctx, &input)
						if err != nil {
							return nil, err
						}

						for _, p := range response.Policies {
							policies = append(policies, *p.PolicyArn)
						}

						if aws.ToString(response.NextMarker) == "" {
							break
						}
						input.Marker = response.NextMarker
					}
					return policies, nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ca_certificate_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owned_by").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_pem").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_version").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateArn")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_mode").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generation_id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_date").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transfer_data").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("validity").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsIotCertificatesGenerator) GetSubTables() []*schema.Table {
	return nil
}
