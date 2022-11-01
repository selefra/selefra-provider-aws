package mocks

import (
	context "context"
	reflect "reflect"

	iam "github.com/aws/aws-sdk-go-v2/service/iam"
	gomock "github.com/golang/mock/gomock"
)

type MockIamClient struct {
	ctrl		*gomock.Controller
	recorder	*MockIamClientMockRecorder
}

type MockIamClientMockRecorder struct {
	mock *MockIamClient
}

func NewMockIamClient(ctrl *gomock.Controller) *MockIamClient {
	mock := &MockIamClient{ctrl: ctrl}
	mock.recorder = &MockIamClientMockRecorder{mock}
	return mock
}

func (m *MockIamClient) EXPECT() *MockIamClientMockRecorder {
	return m.recorder
}

func (m *MockIamClient) GenerateCredentialReport(arg0 context.Context, arg1 *iam.GenerateCredentialReportInput, arg2 ...func(*iam.Options)) (*iam.GenerateCredentialReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GenerateCredentialReport", varargs...)
	ret0, _ := ret[0].(*iam.GenerateCredentialReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GenerateCredentialReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCredentialReport", reflect.TypeOf((*MockIamClient)(nil).GenerateCredentialReport), varargs...)
}

func (m *MockIamClient) GetAccessKeyLastUsed(arg0 context.Context, arg1 *iam.GetAccessKeyLastUsedInput, arg2 ...func(*iam.Options)) (*iam.GetAccessKeyLastUsedOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessKeyLastUsed", varargs...)
	ret0, _ := ret[0].(*iam.GetAccessKeyLastUsedOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetAccessKeyLastUsed(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKeyLastUsed", reflect.TypeOf((*MockIamClient)(nil).GetAccessKeyLastUsed), varargs...)
}

func (m *MockIamClient) GetAccountAuthorizationDetails(arg0 context.Context, arg1 *iam.GetAccountAuthorizationDetailsInput, arg2 ...func(*iam.Options)) (*iam.GetAccountAuthorizationDetailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountAuthorizationDetails", varargs...)
	ret0, _ := ret[0].(*iam.GetAccountAuthorizationDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetAccountAuthorizationDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountAuthorizationDetails", reflect.TypeOf((*MockIamClient)(nil).GetAccountAuthorizationDetails), varargs...)
}

func (m *MockIamClient) GetAccountPasswordPolicy(arg0 context.Context, arg1 *iam.GetAccountPasswordPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetAccountPasswordPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountPasswordPolicy", varargs...)
	ret0, _ := ret[0].(*iam.GetAccountPasswordPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetAccountPasswordPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountPasswordPolicy", reflect.TypeOf((*MockIamClient)(nil).GetAccountPasswordPolicy), varargs...)
}

func (m *MockIamClient) GetAccountSummary(arg0 context.Context, arg1 *iam.GetAccountSummaryInput, arg2 ...func(*iam.Options)) (*iam.GetAccountSummaryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccountSummary", varargs...)
	ret0, _ := ret[0].(*iam.GetAccountSummaryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetAccountSummary(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountSummary", reflect.TypeOf((*MockIamClient)(nil).GetAccountSummary), varargs...)
}

func (m *MockIamClient) GetCredentialReport(arg0 context.Context, arg1 *iam.GetCredentialReportInput, arg2 ...func(*iam.Options)) (*iam.GetCredentialReportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCredentialReport", varargs...)
	ret0, _ := ret[0].(*iam.GetCredentialReportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetCredentialReport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialReport", reflect.TypeOf((*MockIamClient)(nil).GetCredentialReport), varargs...)
}

func (m *MockIamClient) GetGroupPolicy(arg0 context.Context, arg1 *iam.GetGroupPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetGroupPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupPolicy", varargs...)
	ret0, _ := ret[0].(*iam.GetGroupPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetGroupPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupPolicy", reflect.TypeOf((*MockIamClient)(nil).GetGroupPolicy), varargs...)
}

func (m *MockIamClient) GetOpenIDConnectProvider(arg0 context.Context, arg1 *iam.GetOpenIDConnectProviderInput, arg2 ...func(*iam.Options)) (*iam.GetOpenIDConnectProviderOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOpenIDConnectProvider", varargs...)
	ret0, _ := ret[0].(*iam.GetOpenIDConnectProviderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetOpenIDConnectProvider(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenIDConnectProvider", reflect.TypeOf((*MockIamClient)(nil).GetOpenIDConnectProvider), varargs...)
}

func (m *MockIamClient) GetRole(arg0 context.Context, arg1 *iam.GetRoleInput, arg2 ...func(*iam.Options)) (*iam.GetRoleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRole", varargs...)
	ret0, _ := ret[0].(*iam.GetRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetRole(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockIamClient)(nil).GetRole), varargs...)
}

func (m *MockIamClient) GetRolePolicy(arg0 context.Context, arg1 *iam.GetRolePolicyInput, arg2 ...func(*iam.Options)) (*iam.GetRolePolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRolePolicy", varargs...)
	ret0, _ := ret[0].(*iam.GetRolePolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetRolePolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRolePolicy", reflect.TypeOf((*MockIamClient)(nil).GetRolePolicy), varargs...)
}

func (m *MockIamClient) GetSAMLProvider(arg0 context.Context, arg1 *iam.GetSAMLProviderInput, arg2 ...func(*iam.Options)) (*iam.GetSAMLProviderOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSAMLProvider", varargs...)
	ret0, _ := ret[0].(*iam.GetSAMLProviderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetSAMLProvider(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSAMLProvider", reflect.TypeOf((*MockIamClient)(nil).GetSAMLProvider), varargs...)
}

func (m *MockIamClient) GetUser(arg0 context.Context, arg1 *iam.GetUserInput, arg2 ...func(*iam.Options)) (*iam.GetUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUser", varargs...)
	ret0, _ := ret[0].(*iam.GetUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIamClient)(nil).GetUser), varargs...)
}

func (m *MockIamClient) GetUserPolicy(arg0 context.Context, arg1 *iam.GetUserPolicyInput, arg2 ...func(*iam.Options)) (*iam.GetUserPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserPolicy", varargs...)
	ret0, _ := ret[0].(*iam.GetUserPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) GetUserPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPolicy", reflect.TypeOf((*MockIamClient)(nil).GetUserPolicy), varargs...)
}

func (m *MockIamClient) ListAccessKeys(arg0 context.Context, arg1 *iam.ListAccessKeysInput, arg2 ...func(*iam.Options)) (*iam.ListAccessKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccessKeys", varargs...)
	ret0, _ := ret[0].(*iam.ListAccessKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListAccessKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessKeys", reflect.TypeOf((*MockIamClient)(nil).ListAccessKeys), varargs...)
}

func (m *MockIamClient) ListAccountAliases(arg0 context.Context, arg1 *iam.ListAccountAliasesInput, arg2 ...func(*iam.Options)) (*iam.ListAccountAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccountAliases", varargs...)
	ret0, _ := ret[0].(*iam.ListAccountAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListAccountAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountAliases", reflect.TypeOf((*MockIamClient)(nil).ListAccountAliases), varargs...)
}

func (m *MockIamClient) ListAttachedGroupPolicies(arg0 context.Context, arg1 *iam.ListAttachedGroupPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListAttachedGroupPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttachedGroupPolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListAttachedGroupPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListAttachedGroupPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedGroupPolicies", reflect.TypeOf((*MockIamClient)(nil).ListAttachedGroupPolicies), varargs...)
}

func (m *MockIamClient) ListAttachedRolePolicies(arg0 context.Context, arg1 *iam.ListAttachedRolePoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListAttachedRolePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttachedRolePolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListAttachedRolePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListAttachedRolePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedRolePolicies", reflect.TypeOf((*MockIamClient)(nil).ListAttachedRolePolicies), varargs...)
}

func (m *MockIamClient) ListAttachedUserPolicies(arg0 context.Context, arg1 *iam.ListAttachedUserPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListAttachedUserPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttachedUserPolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListAttachedUserPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListAttachedUserPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedUserPolicies", reflect.TypeOf((*MockIamClient)(nil).ListAttachedUserPolicies), varargs...)
}

func (m *MockIamClient) ListGroupPolicies(arg0 context.Context, arg1 *iam.ListGroupPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListGroupPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupPolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListGroupPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListGroupPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupPolicies", reflect.TypeOf((*MockIamClient)(nil).ListGroupPolicies), varargs...)
}

func (m *MockIamClient) ListGroups(arg0 context.Context, arg1 *iam.ListGroupsInput, arg2 ...func(*iam.Options)) (*iam.ListGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroups", varargs...)
	ret0, _ := ret[0].(*iam.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockIamClient)(nil).ListGroups), varargs...)
}

func (m *MockIamClient) ListGroupsForUser(arg0 context.Context, arg1 *iam.ListGroupsForUserInput, arg2 ...func(*iam.Options)) (*iam.ListGroupsForUserOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroupsForUser", varargs...)
	ret0, _ := ret[0].(*iam.ListGroupsForUserOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListGroupsForUser(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupsForUser", reflect.TypeOf((*MockIamClient)(nil).ListGroupsForUser), varargs...)
}

func (m *MockIamClient) ListOpenIDConnectProviders(arg0 context.Context, arg1 *iam.ListOpenIDConnectProvidersInput, arg2 ...func(*iam.Options)) (*iam.ListOpenIDConnectProvidersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOpenIDConnectProviders", varargs...)
	ret0, _ := ret[0].(*iam.ListOpenIDConnectProvidersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListOpenIDConnectProviders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOpenIDConnectProviders", reflect.TypeOf((*MockIamClient)(nil).ListOpenIDConnectProviders), varargs...)
}

func (m *MockIamClient) ListPolicyTags(arg0 context.Context, arg1 *iam.ListPolicyTagsInput, arg2 ...func(*iam.Options)) (*iam.ListPolicyTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPolicyTags", varargs...)
	ret0, _ := ret[0].(*iam.ListPolicyTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListPolicyTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicyTags", reflect.TypeOf((*MockIamClient)(nil).ListPolicyTags), varargs...)
}

func (m *MockIamClient) ListRolePolicies(arg0 context.Context, arg1 *iam.ListRolePoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListRolePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRolePolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListRolePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListRolePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRolePolicies", reflect.TypeOf((*MockIamClient)(nil).ListRolePolicies), varargs...)
}

func (m *MockIamClient) ListRoles(arg0 context.Context, arg1 *iam.ListRolesInput, arg2 ...func(*iam.Options)) (*iam.ListRolesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRoles", varargs...)
	ret0, _ := ret[0].(*iam.ListRolesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListRoles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoles", reflect.TypeOf((*MockIamClient)(nil).ListRoles), varargs...)
}

func (m *MockIamClient) ListSAMLProviders(arg0 context.Context, arg1 *iam.ListSAMLProvidersInput, arg2 ...func(*iam.Options)) (*iam.ListSAMLProvidersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSAMLProviders", varargs...)
	ret0, _ := ret[0].(*iam.ListSAMLProvidersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListSAMLProviders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSAMLProviders", reflect.TypeOf((*MockIamClient)(nil).ListSAMLProviders), varargs...)
}

func (m *MockIamClient) ListServerCertificates(arg0 context.Context, arg1 *iam.ListServerCertificatesInput, arg2 ...func(*iam.Options)) (*iam.ListServerCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServerCertificates", varargs...)
	ret0, _ := ret[0].(*iam.ListServerCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListServerCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServerCertificates", reflect.TypeOf((*MockIamClient)(nil).ListServerCertificates), varargs...)
}

func (m *MockIamClient) ListUserPolicies(arg0 context.Context, arg1 *iam.ListUserPoliciesInput, arg2 ...func(*iam.Options)) (*iam.ListUserPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUserPolicies", varargs...)
	ret0, _ := ret[0].(*iam.ListUserPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListUserPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserPolicies", reflect.TypeOf((*MockIamClient)(nil).ListUserPolicies), varargs...)
}

func (m *MockIamClient) ListUsers(arg0 context.Context, arg1 *iam.ListUsersInput, arg2 ...func(*iam.Options)) (*iam.ListUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUsers", varargs...)
	ret0, _ := ret[0].(*iam.ListUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockIamClient)(nil).ListUsers), varargs...)
}

func (m *MockIamClient) ListVirtualMFADevices(arg0 context.Context, arg1 *iam.ListVirtualMFADevicesInput, arg2 ...func(*iam.Options)) (*iam.ListVirtualMFADevicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVirtualMFADevices", varargs...)
	ret0, _ := ret[0].(*iam.ListVirtualMFADevicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIamClientMockRecorder) ListVirtualMFADevices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualMFADevices", reflect.TypeOf((*MockIamClient)(nil).ListVirtualMFADevices), varargs...)
}
