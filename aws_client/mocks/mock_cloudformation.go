package mocks

import (
	context "context"
	reflect "reflect"

	cloudformation "github.com/aws/aws-sdk-go-v2/service/cloudformation"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudFormationClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCloudFormationClientMockRecorder
}

type MockCloudFormationClientMockRecorder struct {
	mock *MockCloudFormationClient
}

func NewMockCloudFormationClient(ctrl *gomock.Controller) *MockCloudFormationClient {
	mock := &MockCloudFormationClient{ctrl: ctrl}
	mock.recorder = &MockCloudFormationClientMockRecorder{mock}
	return mock
}

func (m *MockCloudFormationClient) EXPECT() *MockCloudFormationClientMockRecorder {
	return m.recorder
}

func (m *MockCloudFormationClient) DescribeStacks(arg0 context.Context, arg1 *cloudformation.DescribeStacksInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.DescribeStacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStacks", varargs...)
	ret0, _ := ret[0].(*cloudformation.DescribeStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudFormationClientMockRecorder) DescribeStacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStacks", reflect.TypeOf((*MockCloudFormationClient)(nil).DescribeStacks), varargs...)
}

func (m *MockCloudFormationClient) ListStackResources(arg0 context.Context, arg1 *cloudformation.ListStackResourcesInput, arg2 ...func(*cloudformation.Options)) (*cloudformation.ListStackResourcesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStackResources", varargs...)
	ret0, _ := ret[0].(*cloudformation.ListStackResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudFormationClientMockRecorder) ListStackResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStackResources", reflect.TypeOf((*MockCloudFormationClient)(nil).ListStackResources), varargs...)
}
