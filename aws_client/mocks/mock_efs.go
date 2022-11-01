package mocks

import (
	context "context"
	reflect "reflect"

	efs "github.com/aws/aws-sdk-go-v2/service/efs"
	gomock "github.com/golang/mock/gomock"
)

type MockEfsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEfsClientMockRecorder
}

type MockEfsClientMockRecorder struct {
	mock *MockEfsClient
}

func NewMockEfsClient(ctrl *gomock.Controller) *MockEfsClient {
	mock := &MockEfsClient{ctrl: ctrl}
	mock.recorder = &MockEfsClientMockRecorder{mock}
	return mock
}

func (m *MockEfsClient) EXPECT() *MockEfsClientMockRecorder {
	return m.recorder
}

func (m *MockEfsClient) DescribeBackupPolicy(arg0 context.Context, arg1 *efs.DescribeBackupPolicyInput, arg2 ...func(*efs.Options)) (*efs.DescribeBackupPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBackupPolicy", varargs...)
	ret0, _ := ret[0].(*efs.DescribeBackupPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeBackupPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackupPolicy", reflect.TypeOf((*MockEfsClient)(nil).DescribeBackupPolicy), varargs...)
}

func (m *MockEfsClient) DescribeFileSystems(arg0 context.Context, arg1 *efs.DescribeFileSystemsInput, arg2 ...func(*efs.Options)) (*efs.DescribeFileSystemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFileSystems", varargs...)
	ret0, _ := ret[0].(*efs.DescribeFileSystemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEfsClientMockRecorder) DescribeFileSystems(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystems", reflect.TypeOf((*MockEfsClient)(nil).DescribeFileSystems), varargs...)
}
