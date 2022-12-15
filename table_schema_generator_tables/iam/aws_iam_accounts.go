package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/mitchellh/mapstructure"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAwsIamAccountsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAwsIamAccountsGenerator{}

func (x *TableAwsIamAccountsGenerator) GetTableName() string {
	return "aws_iam_accounts"
}

func (x *TableAwsIamAccountsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAwsIamAccountsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAwsIamAccountsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"account_id",
		},
	}
}

func (x *TableAwsIamAccountsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*aws_client.Client).AwsServices().IAM

			summary, err := svc.GetAccountSummary(ctx, &iam.GetAccountSummaryInput{})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			var accSummary Account
			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", WeaklyTypedInput: true, Result: &accSummary})
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if err := decoder.Decode(summary.SummaryMap); err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			config := iam.ListAccountAliasesInput{}
			for {
				response, err := svc.ListAccountAliases(ctx, &config)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				accSummary.Aliases = append(accSummary.Aliases, response.AccountAliases...)

				if aws.ToString(response.Marker) == "" {
					break
				}
				config.Marker = response.Marker
			}
			resultChannel <- accSummary
			return nil
		},
	}
}

type Account struct {
	Users                             int32    `json:"users,omitempty"`
	UsersQuota                        int32    `json:"users_quota,omitempty"`
	Groups                            int32    `json:"groups,omitempty"`
	GroupsQuota                       int32    `json:"groups_quota,omitempty"`
	ServerCertificates                int32    `json:"server_certificates,omitempty"`
	ServerCertificatesQuota           int32    `json:"server_certificates_quota,omitempty"`
	UserPolicySizeQuota               int32    `json:"user_policy_size_quota,omitempty"`
	GroupPolicySizeQuota              int32    `json:"group_policy_size_quota,omitempty"`
	GroupsPerUserQuota                int32    `json:"groups_per_user_quota,omitempty"`
	SigningCertificatesPerUserQuota   int32    `json:"signing_certificates_per_user_quota,omitempty"`
	AccessKeysPerUserQuota            int32    `json:"access_keys_per_user_quota,omitempty"`
	MFADevices                        int32    `json:"mfa_devices,omitempty"`
	MFADevicesInUse                   int32    `json:"mfa_devices_in_use,omitempty"`
	AccountMFAEnabled                 bool     `json:"account_mfa_enabled,omitempty"`
	AccountAccessKeysPresent          bool     `json:"account_access_keys_present,omitempty"`
	AccountSigningCertificatesPresent bool     `json:"account_signing_certificates_present,omitempty"`
	AttachedPoliciesPerGroupQuota     int32    `json:"attached_policies_per_group_quota,omitempty"`
	AttachedPoliciesPerRoleQuota      int32    `json:"attached_policies_per_role_quota,omitempty"`
	AttachedPoliciesPerUserQuota      int32    `json:"attached_policies_per_user_quota,omitempty"`
	Policies                          int32    `json:"policies,omitempty"`
	PoliciesQuota                     int32    `json:"policies_quota,omitempty"`
	PolicySizeQuota                   int32    `json:"policy_size_quota,omitempty"`
	PolicyVersionsInUse               int32    `json:"policy_versions_in_use,omitempty"`
	PolicyVersionsInUseQuota          int32    `json:"policy_versions_in_use_quota,omitempty"`
	VersionsPerPolicyQuota            int32    `json:"versions_per_policy_quota,omitempty"`
	GlobalEndpointTokenVersion        int32    `json:"global_endpoint_token_version,omitempty"`
	Aliases                           []string `json:"aliases,omitempty"`
}

func (x *TableAwsIamAccountsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return aws_client.ExpandByPartition()
}

func (x *TableAwsIamAccountsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("server_certificates").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("server_certificates_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_mfa_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AccountMFAEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_signing_certificates_present").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_versions_in_use_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_id").ColumnType(schema.ColumnTypeString).
			Extractor(aws_client.AwsAccountIDExtractor()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("groups_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("signing_certificates_per_user_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mfa_devices_in_use").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MFADevicesInUse")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aliases").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_policy_size_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access_keys_per_user_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("users").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mfa_devices").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MFADevices")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attached_policies_per_group_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attached_policies_per_user_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policies").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_versions_in_use").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("groups").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attached_policies_per_role_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("users_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_policy_size_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("groups_per_user_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("account_access_keys_present").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policies_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("versions_per_policy_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy_size_quota").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("global_endpoint_token_version").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAwsIamAccountsGenerator) GetSubTables() []*schema.Table {
	return nil
}
