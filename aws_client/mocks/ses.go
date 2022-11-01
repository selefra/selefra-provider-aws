package mocks

import (
	context "context"
	reflect "reflect"

	sesv2 "github.com/aws/aws-sdk-go-v2/service/sesv2"
	gomock "github.com/golang/mock/gomock"
)

type MockSESClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSESClientMockRecorder
}

type MockSESClientMockRecorder struct {
	mock *MockSESClient
}

func NewMockSESClient(ctrl *gomock.Controller) *MockSESClient {
	mock := &MockSESClient{ctrl: ctrl}
	mock.recorder = &MockSESClientMockRecorder{mock}
	return mock
}

func (m *MockSESClient) EXPECT() *MockSESClientMockRecorder {
	return m.recorder
}

func (m *MockSESClient) GetEmailTemplate(arg0 context.Context, arg1 *sesv2.GetEmailTemplateInput, arg2 ...func(*sesv2.Options)) (*sesv2.GetEmailTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEmailTemplate", varargs...)
	ret0, _ := ret[0].(*sesv2.GetEmailTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSESClientMockRecorder) GetEmailTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmailTemplate", reflect.TypeOf((*MockSESClient)(nil).GetEmailTemplate), varargs...)
}

func (m *MockSESClient) ListEmailTemplates(arg0 context.Context, arg1 *sesv2.ListEmailTemplatesInput, arg2 ...func(*sesv2.Options)) (*sesv2.ListEmailTemplatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEmailTemplates", varargs...)
	ret0, _ := ret[0].(*sesv2.ListEmailTemplatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSESClientMockRecorder) ListEmailTemplates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEmailTemplates", reflect.TypeOf((*MockSESClient)(nil).ListEmailTemplates), varargs...)
}
