package mocks

import (
	context "context"
	reflect "reflect"

	workspaces "github.com/aws/aws-sdk-go-v2/service/workspaces"
	gomock "github.com/golang/mock/gomock"
)

type MockWorkspacesClient struct {
	ctrl		*gomock.Controller
	recorder	*MockWorkspacesClientMockRecorder
}

type MockWorkspacesClientMockRecorder struct {
	mock *MockWorkspacesClient
}

func NewMockWorkspacesClient(ctrl *gomock.Controller) *MockWorkspacesClient {
	mock := &MockWorkspacesClient{ctrl: ctrl}
	mock.recorder = &MockWorkspacesClientMockRecorder{mock}
	return mock
}

func (m *MockWorkspacesClient) EXPECT() *MockWorkspacesClientMockRecorder {
	return m.recorder
}

func (m *MockWorkspacesClient) DescribeWorkspaceDirectories(arg0 context.Context, arg1 *workspaces.DescribeWorkspaceDirectoriesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspaceDirectoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkspaceDirectories", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspaceDirectoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaceDirectories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkspaceDirectories", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaceDirectories), varargs...)
}

func (m *MockWorkspacesClient) DescribeWorkspaces(arg0 context.Context, arg1 *workspaces.DescribeWorkspacesInput, arg2 ...func(*workspaces.Options)) (*workspaces.DescribeWorkspacesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkspaces", varargs...)
	ret0, _ := ret[0].(*workspaces.DescribeWorkspacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWorkspacesClientMockRecorder) DescribeWorkspaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkspaces", reflect.TypeOf((*MockWorkspacesClient)(nil).DescribeWorkspaces), varargs...)
}
