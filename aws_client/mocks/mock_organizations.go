package mocks

import (
	context "context"
	reflect "reflect"

	organizations "github.com/aws/aws-sdk-go-v2/service/organizations"
	gomock "github.com/golang/mock/gomock"
)

type MockOrganizationsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockOrganizationsClientMockRecorder
}

type MockOrganizationsClientMockRecorder struct {
	mock *MockOrganizationsClient
}

func NewMockOrganizationsClient(ctrl *gomock.Controller) *MockOrganizationsClient {
	mock := &MockOrganizationsClient{ctrl: ctrl}
	mock.recorder = &MockOrganizationsClientMockRecorder{mock}
	return mock
}

func (m *MockOrganizationsClient) EXPECT() *MockOrganizationsClientMockRecorder {
	return m.recorder
}

func (m *MockOrganizationsClient) ListAccounts(arg0 context.Context, arg1 *organizations.ListAccountsInput, arg2 ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccounts", varargs...)
	ret0, _ := ret[0].(*organizations.ListAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockOrganizationsClient)(nil).ListAccounts), varargs...)
}

func (m *MockOrganizationsClient) ListAccountsForParent(arg0 context.Context, arg1 *organizations.ListAccountsForParentInput, arg2 ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAccountsForParent", varargs...)
	ret0, _ := ret[0].(*organizations.ListAccountsForParentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListAccountsForParent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountsForParent", reflect.TypeOf((*MockOrganizationsClient)(nil).ListAccountsForParent), varargs...)
}

func (m *MockOrganizationsClient) ListTagsForResource(arg0 context.Context, arg1 *organizations.ListTagsForResourceInput, arg2 ...func(*organizations.Options)) (*organizations.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*organizations.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockOrganizationsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockOrganizationsClient)(nil).ListTagsForResource), varargs...)
}
