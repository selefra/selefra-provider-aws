package iam

import (
	"context"
	"errors"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/smithy-go"
	"github.com/gocarina/gocsv"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIamCredentialReportsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamCredentialReportsGenerator{}

func (x *TableAwsIamCredentialReportsGenerator) GetTableName() string {
	return "aws_iam_credential_reports"
}

func (x *TableAwsIamCredentialReportsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamCredentialReportsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamCredentialReportsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"arn",
			"user_creation_time",
		},
	}
}

func (x *TableAwsIamCredentialReportsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			var err error
			var apiErr smithy.APIError
			var reportOutput *iam.GetCredentialReportOutput
			svc := client.(*aws_client.Client).AwsServices().IAM
			for {
				reportOutput, err = svc.GetCredentialReport(ctx, &iam.GetCredentialReportInput{})
				if err == nil && reportOutput != nil {
					var users []*CredentialReportEntry
					err = gocsv.UnmarshalBytes(reportOutput.Content, &users)
					if err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
					resultChannel <- users
				}
				if !errors.As(err, &apiErr) {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				switch apiErr.ErrorCode() {
				case "ReportNotPresent", "ReportExpired":
					_, err := svc.GenerateCredentialReport(ctx, &iam.GenerateCredentialReportInput{})
					if err != nil {
						var serviceError smithy.APIError
						if !errors.As(err, &serviceError) {
							return schema.NewDiagnosticsErrorPullTable(task.Table, err)

						}

						if serviceError.ErrorCode() != "LimitExceeded" {
							return schema.NewDiagnosticsErrorPullTable(task.Table, err)

						}
						if err := aws_client.Sleep(ctx, 5*time.Second); err != nil {
							return schema.NewDiagnosticsErrorPullTable(task.Table, err)

						}
					}
				case "ReportInProgress":

					if err := aws_client.Sleep(ctx, 5*time.Second); err != nil {
						return schema.NewDiagnosticsErrorPullTable(task.Table, err)

					}
				default:
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
			}
		},
	}
}

type CredentialReportEntry struct {
	User				string		`csv:"user"`
	Arn				string		`csv:"arn"`
	UserCreationTime		DateTime	`csv:"user_creation_time"`
	PasswordStatus			string		`csv:"password_enabled"`
	PasswordLastChanged		DateTime	`csv:"password_last_changed"`
	PasswordNextRotation		DateTime	`csv:"password_next_rotation"`
	MfaActive			bool		`csv:"mfa_active"`
	AccessKey1Active		bool		`csv:"access_key_1_active"`
	AccessKey2Active		bool		`csv:"access_key_2_active"`
	AccessKey1LastRotated		DateTime	`csv:"access_key_1_last_rotated"`
	AccessKey2LastRotated		DateTime	`csv:"access_key_2_last_rotated"`
	Cert1Active			bool		`csv:"cert_1_active"`
	Cert2Active			bool		`csv:"cert_2_active"`
	Cert1LastRotated		DateTime	`csv:"cert_1_last_rotated"`
	Cert2LastRotated		DateTime	`csv:"cert_2_last_rotated"`
	AccessKey1LastUsedDate		DateTime	`csv:"access_key_1_last_used_date"`
	AccessKey1LastUsedRegion	string		`csv:"access_key_1_last_used_region"`
	AccessKey1LastUsedService	string		`csv:"access_key_1_last_used_service"`
	AccessKey2LastUsedDate		DateTime	`csv:"access_key_2_last_used_date"`
	AccessKey2LastUsedRegion	string		`csv:"access_key_2_last_used_region"`
	AccessKey2LastUsedService	string		`csv:"access_key_2_last_used_service"`
	PasswordLastUsed		DateTime	`csv:"password_last_used"`
}

type DateTime struct {
	*time.Time
}

func (d *DateTime) UnmarshalCSV(val string) (err error) {
	switch val {
	case "N/A", "not_supported":
		d.Time = nil
		return nil
	}
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return err
	}
	d.Time = &t
	return nil
}

func (x *TableAwsIamCredentialReportsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamCredentialReportsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_1_last_used_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				r := result.(*CredentialReportEntry)
				return r.AccessKey1LastUsedDate.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_last_used").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				r := result.(*CredentialReportEntry)
				return r.PasswordLastUsed.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cert2_active").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key1_last_used_region").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_creation_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				r := result.(*CredentialReportEntry)
				return r.UserCreationTime.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_last_changed").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				r := result.(*CredentialReportEntry)
				return r.PasswordLastChanged.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_2_last_rotated").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				r := result.(*CredentialReportEntry)
				return r.AccessKey2LastRotated.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cert_1_last_rotated").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				r := result.(*CredentialReportEntry)
				return r.Cert1LastRotated.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key2_last_used_region").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key2_last_used_service").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("arn").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cert_2_last_rotated").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				r := result.(*CredentialReportEntry)
				return r.Cert2LastRotated.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key1_active").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key2_active").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key1_last_used_service").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_next_rotation").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				r := result.(*CredentialReportEntry)
				return r.PasswordNextRotation.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_1_last_rotated").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				r := result.(*CredentialReportEntry)
				return r.AccessKey1LastRotated.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_key_2_last_used_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				r := result.(*CredentialReportEntry)
				return r.AccessKey2LastUsedDate.String(), nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mfa_active").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cert1_active").ColumnType(schema.ColumnTypeBool).Build(),
	}
}

func (x *TableAwsIamCredentialReportsGenerator) GetSubTables() []*schema.Table {
	return nil
}
