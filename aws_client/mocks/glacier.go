package mocks

import (
	context "context"
	reflect "reflect"

	glacier "github.com/aws/aws-sdk-go-v2/service/glacier"
	gomock "github.com/golang/mock/gomock"
)

type MockGlacierClient struct {
	ctrl		*gomock.Controller
	recorder	*MockGlacierClientMockRecorder
}

type MockGlacierClientMockRecorder struct {
	mock *MockGlacierClient
}

func NewMockGlacierClient(ctrl *gomock.Controller) *MockGlacierClient {
	mock := &MockGlacierClient{ctrl: ctrl}
	mock.recorder = &MockGlacierClientMockRecorder{mock}
	return mock
}

func (m *MockGlacierClient) EXPECT() *MockGlacierClientMockRecorder {
	return m.recorder
}

func (m *MockGlacierClient) GetDataRetrievalPolicy(arg0 context.Context, arg1 *glacier.GetDataRetrievalPolicyInput, arg2 ...func(*glacier.Options)) (*glacier.GetDataRetrievalPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDataRetrievalPolicy", varargs...)
	ret0, _ := ret[0].(*glacier.GetDataRetrievalPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) GetDataRetrievalPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataRetrievalPolicy", reflect.TypeOf((*MockGlacierClient)(nil).GetDataRetrievalPolicy), varargs...)
}

func (m *MockGlacierClient) GetVaultAccessPolicy(arg0 context.Context, arg1 *glacier.GetVaultAccessPolicyInput, arg2 ...func(*glacier.Options)) (*glacier.GetVaultAccessPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVaultAccessPolicy", varargs...)
	ret0, _ := ret[0].(*glacier.GetVaultAccessPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) GetVaultAccessPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVaultAccessPolicy", reflect.TypeOf((*MockGlacierClient)(nil).GetVaultAccessPolicy), varargs...)
}

func (m *MockGlacierClient) GetVaultLock(arg0 context.Context, arg1 *glacier.GetVaultLockInput, arg2 ...func(*glacier.Options)) (*glacier.GetVaultLockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVaultLock", varargs...)
	ret0, _ := ret[0].(*glacier.GetVaultLockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) GetVaultLock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVaultLock", reflect.TypeOf((*MockGlacierClient)(nil).GetVaultLock), varargs...)
}

func (m *MockGlacierClient) GetVaultNotifications(arg0 context.Context, arg1 *glacier.GetVaultNotificationsInput, arg2 ...func(*glacier.Options)) (*glacier.GetVaultNotificationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVaultNotifications", varargs...)
	ret0, _ := ret[0].(*glacier.GetVaultNotificationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) GetVaultNotifications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVaultNotifications", reflect.TypeOf((*MockGlacierClient)(nil).GetVaultNotifications), varargs...)
}

func (m *MockGlacierClient) ListTagsForVault(arg0 context.Context, arg1 *glacier.ListTagsForVaultInput, arg2 ...func(*glacier.Options)) (*glacier.ListTagsForVaultOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForVault", varargs...)
	ret0, _ := ret[0].(*glacier.ListTagsForVaultOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) ListTagsForVault(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForVault", reflect.TypeOf((*MockGlacierClient)(nil).ListTagsForVault), varargs...)
}

func (m *MockGlacierClient) ListVaults(arg0 context.Context, arg1 *glacier.ListVaultsInput, arg2 ...func(*glacier.Options)) (*glacier.ListVaultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVaults", varargs...)
	ret0, _ := ret[0].(*glacier.ListVaultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlacierClientMockRecorder) ListVaults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVaults", reflect.TypeOf((*MockGlacierClient)(nil).ListVaults), varargs...)
}
