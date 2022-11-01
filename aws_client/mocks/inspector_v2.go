package mocks

import (
	context "context"
	reflect "reflect"

	inspector2 "github.com/aws/aws-sdk-go-v2/service/inspector2"
	gomock "github.com/golang/mock/gomock"
)

type MockInspectorV2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockInspectorV2ClientMockRecorder
}

type MockInspectorV2ClientMockRecorder struct {
	mock *MockInspectorV2Client
}

func NewMockInspectorV2Client(ctrl *gomock.Controller) *MockInspectorV2Client {
	mock := &MockInspectorV2Client{ctrl: ctrl}
	mock.recorder = &MockInspectorV2ClientMockRecorder{mock}
	return mock
}

func (m *MockInspectorV2Client) EXPECT() *MockInspectorV2ClientMockRecorder {
	return m.recorder
}

func (m *MockInspectorV2Client) ListFindings(arg0 context.Context, arg1 *inspector2.ListFindingsInput, arg2 ...func(*inspector2.Options)) (*inspector2.ListFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFindings", varargs...)
	ret0, _ := ret[0].(*inspector2.ListFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockInspectorV2ClientMockRecorder) ListFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFindings", reflect.TypeOf((*MockInspectorV2Client)(nil).ListFindings), varargs...)
}
