package mocks

import (
	context "context"
	reflect "reflect"

	codebuild "github.com/aws/aws-sdk-go-v2/service/codebuild"
	gomock "github.com/golang/mock/gomock"
)

type MockCodebuildClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCodebuildClientMockRecorder
}

type MockCodebuildClientMockRecorder struct {
	mock *MockCodebuildClient
}

func NewMockCodebuildClient(ctrl *gomock.Controller) *MockCodebuildClient {
	mock := &MockCodebuildClient{ctrl: ctrl}
	mock.recorder = &MockCodebuildClientMockRecorder{mock}
	return mock
}

func (m *MockCodebuildClient) EXPECT() *MockCodebuildClientMockRecorder {
	return m.recorder
}

func (m *MockCodebuildClient) BatchGetProjects(arg0 context.Context, arg1 *codebuild.BatchGetProjectsInput, arg2 ...func(*codebuild.Options)) (*codebuild.BatchGetProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetProjects", varargs...)
	ret0, _ := ret[0].(*codebuild.BatchGetProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) BatchGetProjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetProjects", reflect.TypeOf((*MockCodebuildClient)(nil).BatchGetProjects), varargs...)
}

func (m *MockCodebuildClient) ListProjects(arg0 context.Context, arg1 *codebuild.ListProjectsInput, arg2 ...func(*codebuild.Options)) (*codebuild.ListProjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjects", varargs...)
	ret0, _ := ret[0].(*codebuild.ListProjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodebuildClientMockRecorder) ListProjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockCodebuildClient)(nil).ListProjects), varargs...)
}
