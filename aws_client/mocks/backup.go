package mocks

import (
	context "context"
	reflect "reflect"

	backup "github.com/aws/aws-sdk-go-v2/service/backup"
	gomock "github.com/golang/mock/gomock"
)

type MockBackupClient struct {
	ctrl		*gomock.Controller
	recorder	*MockBackupClientMockRecorder
}

type MockBackupClientMockRecorder struct {
	mock *MockBackupClient
}

func NewMockBackupClient(ctrl *gomock.Controller) *MockBackupClient {
	mock := &MockBackupClient{ctrl: ctrl}
	mock.recorder = &MockBackupClientMockRecorder{mock}
	return mock
}

func (m *MockBackupClient) EXPECT() *MockBackupClientMockRecorder {
	return m.recorder
}

func (m *MockBackupClient) DescribeGlobalSettings(arg0 context.Context, arg1 *backup.DescribeGlobalSettingsInput, arg2 ...func(*backup.Options)) (*backup.DescribeGlobalSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeGlobalSettings", varargs...)
	ret0, _ := ret[0].(*backup.DescribeGlobalSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeGlobalSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeGlobalSettings", reflect.TypeOf((*MockBackupClient)(nil).DescribeGlobalSettings), varargs...)
}

func (m *MockBackupClient) DescribeRegionSettings(arg0 context.Context, arg1 *backup.DescribeRegionSettingsInput, arg2 ...func(*backup.Options)) (*backup.DescribeRegionSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRegionSettings", varargs...)
	ret0, _ := ret[0].(*backup.DescribeRegionSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) DescribeRegionSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRegionSettings", reflect.TypeOf((*MockBackupClient)(nil).DescribeRegionSettings), varargs...)
}

func (m *MockBackupClient) GetBackupPlan(arg0 context.Context, arg1 *backup.GetBackupPlanInput, arg2 ...func(*backup.Options)) (*backup.GetBackupPlanOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackupPlan", varargs...)
	ret0, _ := ret[0].(*backup.GetBackupPlanOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetBackupPlan(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupPlan", reflect.TypeOf((*MockBackupClient)(nil).GetBackupPlan), varargs...)
}

func (m *MockBackupClient) GetBackupSelection(arg0 context.Context, arg1 *backup.GetBackupSelectionInput, arg2 ...func(*backup.Options)) (*backup.GetBackupSelectionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackupSelection", varargs...)
	ret0, _ := ret[0].(*backup.GetBackupSelectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetBackupSelection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupSelection", reflect.TypeOf((*MockBackupClient)(nil).GetBackupSelection), varargs...)
}

func (m *MockBackupClient) GetBackupVaultAccessPolicy(arg0 context.Context, arg1 *backup.GetBackupVaultAccessPolicyInput, arg2 ...func(*backup.Options)) (*backup.GetBackupVaultAccessPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackupVaultAccessPolicy", varargs...)
	ret0, _ := ret[0].(*backup.GetBackupVaultAccessPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetBackupVaultAccessPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupVaultAccessPolicy", reflect.TypeOf((*MockBackupClient)(nil).GetBackupVaultAccessPolicy), varargs...)
}

func (m *MockBackupClient) GetBackupVaultNotifications(arg0 context.Context, arg1 *backup.GetBackupVaultNotificationsInput, arg2 ...func(*backup.Options)) (*backup.GetBackupVaultNotificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBackupVaultNotifications", varargs...)
	ret0, _ := ret[0].(*backup.GetBackupVaultNotificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) GetBackupVaultNotifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupVaultNotifications", reflect.TypeOf((*MockBackupClient)(nil).GetBackupVaultNotifications), varargs...)
}

func (m *MockBackupClient) ListBackupPlans(arg0 context.Context, arg1 *backup.ListBackupPlansInput, arg2 ...func(*backup.Options)) (*backup.ListBackupPlansOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBackupPlans", varargs...)
	ret0, _ := ret[0].(*backup.ListBackupPlansOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListBackupPlans(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackupPlans", reflect.TypeOf((*MockBackupClient)(nil).ListBackupPlans), varargs...)
}

func (m *MockBackupClient) ListBackupSelections(arg0 context.Context, arg1 *backup.ListBackupSelectionsInput, arg2 ...func(*backup.Options)) (*backup.ListBackupSelectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBackupSelections", varargs...)
	ret0, _ := ret[0].(*backup.ListBackupSelectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListBackupSelections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackupSelections", reflect.TypeOf((*MockBackupClient)(nil).ListBackupSelections), varargs...)
}

func (m *MockBackupClient) ListBackupVaults(arg0 context.Context, arg1 *backup.ListBackupVaultsInput, arg2 ...func(*backup.Options)) (*backup.ListBackupVaultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBackupVaults", varargs...)
	ret0, _ := ret[0].(*backup.ListBackupVaultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListBackupVaults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackupVaults", reflect.TypeOf((*MockBackupClient)(nil).ListBackupVaults), varargs...)
}

func (m *MockBackupClient) ListRecoveryPointsByBackupVault(arg0 context.Context, arg1 *backup.ListRecoveryPointsByBackupVaultInput, arg2 ...func(*backup.Options)) (*backup.ListRecoveryPointsByBackupVaultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRecoveryPointsByBackupVault", varargs...)
	ret0, _ := ret[0].(*backup.ListRecoveryPointsByBackupVaultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListRecoveryPointsByBackupVault(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecoveryPointsByBackupVault", reflect.TypeOf((*MockBackupClient)(nil).ListRecoveryPointsByBackupVault), varargs...)
}

func (m *MockBackupClient) ListTags(arg0 context.Context, arg1 *backup.ListTagsInput, arg2 ...func(*backup.Options)) (*backup.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTags", varargs...)
	ret0, _ := ret[0].(*backup.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockBackupClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockBackupClient)(nil).ListTags), varargs...)
}
