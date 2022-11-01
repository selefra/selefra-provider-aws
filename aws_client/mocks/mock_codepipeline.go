package mocks

import (
	context "context"
	reflect "reflect"

	codepipeline "github.com/aws/aws-sdk-go-v2/service/codepipeline"
	gomock "github.com/golang/mock/gomock"
)

type MockCodePipelineClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCodePipelineClientMockRecorder
}

type MockCodePipelineClientMockRecorder struct {
	mock *MockCodePipelineClient
}

func NewMockCodePipelineClient(ctrl *gomock.Controller) *MockCodePipelineClient {
	mock := &MockCodePipelineClient{ctrl: ctrl}
	mock.recorder = &MockCodePipelineClientMockRecorder{mock}
	return mock
}

func (m *MockCodePipelineClient) EXPECT() *MockCodePipelineClientMockRecorder {
	return m.recorder
}

func (m *MockCodePipelineClient) GetPipeline(arg0 context.Context, arg1 *codepipeline.GetPipelineInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.GetPipelineOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPipeline", varargs...)
	ret0, _ := ret[0].(*codepipeline.GetPipelineOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodePipelineClientMockRecorder) GetPipeline(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipeline", reflect.TypeOf((*MockCodePipelineClient)(nil).GetPipeline), varargs...)
}

func (m *MockCodePipelineClient) ListPipelines(arg0 context.Context, arg1 *codepipeline.ListPipelinesInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.ListPipelinesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPipelines", varargs...)
	ret0, _ := ret[0].(*codepipeline.ListPipelinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodePipelineClientMockRecorder) ListPipelines(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPipelines", reflect.TypeOf((*MockCodePipelineClient)(nil).ListPipelines), varargs...)
}

func (m *MockCodePipelineClient) ListTagsForResource(arg0 context.Context, arg1 *codepipeline.ListTagsForResourceInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*codepipeline.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodePipelineClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockCodePipelineClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockCodePipelineClient) ListWebhooks(arg0 context.Context, arg1 *codepipeline.ListWebhooksInput, arg2 ...func(*codepipeline.Options)) (*codepipeline.ListWebhooksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWebhooks", varargs...)
	ret0, _ := ret[0].(*codepipeline.ListWebhooksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCodePipelineClientMockRecorder) ListWebhooks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWebhooks", reflect.TypeOf((*MockCodePipelineClient)(nil).ListWebhooks), varargs...)
}
