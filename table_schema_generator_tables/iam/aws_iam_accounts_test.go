package iam

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildAccount(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIamClient(ctrl)

	acc := struct {
		Users					int32	`json:"users,omitempty"`
		UsersQuota				int32	`json:"users_quota,omitempty"`
		Groups					int32	`json:"groups,omitempty"`
		GroupsQuota				int32	`json:"groups_quota,omitempty"`
		ServerCertificates			int32	`json:"server_certificates,omitempty"`
		ServerCertificatesQuota			int32	`json:"server_certificates_quota,omitempty"`
		UserPolicySizeQuota			int32	`json:"user_policy_size_quota,omitempty"`
		GroupPolicySizeQuota			int32	`json:"group_policy_size_quota,omitempty"`
		GroupsPerUserQuota			int32	`json:"groups_per_user_quota,omitempty"`
		SigningCertificatesPerUserQuota		int32	`json:"signing_certificates_per_user_quota,omitempty"`
		AccessKeysPerUserQuota			int32	`json:"access_keys_per_user_quota,omitempty"`
		MFADevices				int32	`json:"mfa_devices"`
		MFADevicesInUse				int32	`json:"mfa_devices_in_use"`
		AccountMFAEnabled			int32	`json:"account_mfa_enabled,omitempty"`
		AccountAccessKeysPresent		int32	`json:"account_access_keys_present,omitempty"`
		AccountSigningCertificatesPresent	int32	`json:"account_signing_certificates_present,omitempty"`
		AttachedPoliciesPerGroupQuota		int32	`json:"attached_policies_per_group_quota,omitempty"`
		AttachedPoliciesPerRoleQuota		int32	`json:"attached_policies_per_role_quota,omitempty"`
		AttachedPoliciesPerUserQuota		int32	`json:"attached_policies_per_user_quota,omitempty"`
		Policies				int32	`json:"policies,omitempty"`
		PoliciesQuota				int32	`json:"policies_quota,omitempty"`
		PolicySizeQuota				int32	`json:"policy_size_quota,omitempty"`
		PolicyVersionsInUse			int32	`json:"policy_versions_in_use,omitempty"`
		PolicyVersionsInUseQuota		int32	`json:"policy_versions_in_use_quota,omitempty"`
		VersionsPerPolicyQuota			int32	`json:"versions_per_policy_quota,omitempty"`
		GlobalEndpointTokenVersion		int32	`json:"global_endpoint_token_version,omitempty"`
	}{}

	if err := faker.FakeObject(&acc); err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(acc)
	if err != nil {
		t.Fatal(err)
	}
	summaryData := make(map[string]int32)
	if err := json.Unmarshal(data, &summaryData); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetAccountSummary(gomock.Any(), gomock.Any()).AnyTimes().Return(&iam.GetAccountSummaryOutput{SummaryMap: summaryData}, nil)
	m.EXPECT().ListAccountAliases(gomock.Any(), gomock.Any()).AnyTimes().Return(&iam.ListAccountAliasesOutput{AccountAliases: []string{"testAccount"}}, nil)

	return aws_client.AwsServices{
		IAM: m,
	}
}

func TestAccounts(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamAccountsGenerator{}), buildAccount, aws_client.TestOptions{})
}
