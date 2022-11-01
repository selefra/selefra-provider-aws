package mocks

import (
	context "context"
	reflect "reflect"

	inspector "github.com/aws/aws-sdk-go-v2/service/inspector"
	gomock "github.com/golang/mock/gomock"
)

type MockInspectorClient struct {
	ctrl		*gomock.Controller
	recorder	*MockInspectorClientMockRecorder
}

type MockInspectorClientMockRecorder struct {
	mock *MockInspectorClient
}

func NewMockInspectorClient(ctrl *gomock.Controller) *MockInspectorClient {
	mock := &MockInspectorClient{ctrl: ctrl}
	mock.recorder = &MockInspectorClientMockRecorder{mock}
	return mock
}

func (m *MockInspectorClient) EXPECT() *MockInspectorClientMockRecorder {
	return m.recorder
}

func (m *MockInspectorClient) DescribeFindings(arg0 context.Context, arg1 *inspector.DescribeFindingsInput, arg2 ...func(*inspector.Options)) (*inspector.DescribeFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFindings", varargs...)
	ret0, _ := ret[0].(*inspector.DescribeFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) DescribeFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFindings", reflect.TypeOf((*MockInspectorClient)(nil).DescribeFindings), varargs...)
}

func (m *MockInspectorClient) ListFindings(arg0 context.Context, arg1 *inspector.ListFindingsInput, arg2 ...func(*inspector.Options)) (*inspector.ListFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFindings", varargs...)
	ret0, _ := ret[0].(*inspector.ListFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorClientMockRecorder) ListFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFindings", reflect.TypeOf((*MockInspectorClient)(nil).ListFindings), varargs...)
}
