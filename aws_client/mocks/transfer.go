package mocks

import (
	context "context"
	reflect "reflect"

	transfer "github.com/aws/aws-sdk-go-v2/service/transfer"
	gomock "github.com/golang/mock/gomock"
)

type MockTransferClient struct {
	ctrl		*gomock.Controller
	recorder	*MockTransferClientMockRecorder
}

type MockTransferClientMockRecorder struct {
	mock *MockTransferClient
}

func NewMockTransferClient(ctrl *gomock.Controller) *MockTransferClient {
	mock := &MockTransferClient{ctrl: ctrl}
	mock.recorder = &MockTransferClientMockRecorder{mock}
	return mock
}

func (m *MockTransferClient) EXPECT() *MockTransferClientMockRecorder {
	return m.recorder
}

func (m *MockTransferClient) DescribeServer(arg0 context.Context, arg1 *transfer.DescribeServerInput, arg2 ...func(*transfer.Options)) (*transfer.DescribeServerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeServer", varargs...)
	ret0, _ := ret[0].(*transfer.DescribeServerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) DescribeServer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServer", reflect.TypeOf((*MockTransferClient)(nil).DescribeServer), varargs...)
}

func (m *MockTransferClient) ListServers(arg0 context.Context, arg1 *transfer.ListServersInput, arg2 ...func(*transfer.Options)) (*transfer.ListServersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServers", varargs...)
	ret0, _ := ret[0].(*transfer.ListServersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListServers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServers", reflect.TypeOf((*MockTransferClient)(nil).ListServers), varargs...)
}

func (m *MockTransferClient) ListTagsForResource(arg0 context.Context, arg1 *transfer.ListTagsForResourceInput, arg2 ...func(*transfer.Options)) (*transfer.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*transfer.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockTransferClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockTransferClient)(nil).ListTagsForResource), varargs...)
}
