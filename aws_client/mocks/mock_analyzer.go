package mocks

import (
	context "context"
	reflect "reflect"

	accessanalyzer "github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	gomock "github.com/golang/mock/gomock"
)

type MockAnalyzerClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAnalyzerClientMockRecorder
}

type MockAnalyzerClientMockRecorder struct {
	mock *MockAnalyzerClient
}

func NewMockAnalyzerClient(ctrl *gomock.Controller) *MockAnalyzerClient {
	mock := &MockAnalyzerClient{ctrl: ctrl}
	mock.recorder = &MockAnalyzerClientMockRecorder{mock}
	return mock
}

func (m *MockAnalyzerClient) EXPECT() *MockAnalyzerClientMockRecorder {
	return m.recorder
}

func (m *MockAnalyzerClient) ListAnalyzers(arg0 context.Context, arg1 *accessanalyzer.ListAnalyzersInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListAnalyzersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAnalyzers", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListAnalyzersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAnalyzerClientMockRecorder) ListAnalyzers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAnalyzers", reflect.TypeOf((*MockAnalyzerClient)(nil).ListAnalyzers), varargs...)
}

func (m *MockAnalyzerClient) ListArchiveRules(arg0 context.Context, arg1 *accessanalyzer.ListArchiveRulesInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListArchiveRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListArchiveRules", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListArchiveRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAnalyzerClientMockRecorder) ListArchiveRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListArchiveRules", reflect.TypeOf((*MockAnalyzerClient)(nil).ListArchiveRules), varargs...)
}

func (m *MockAnalyzerClient) ListFindings(arg0 context.Context, arg1 *accessanalyzer.ListFindingsInput, arg2 ...func(*accessanalyzer.Options)) (*accessanalyzer.ListFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFindings", varargs...)
	ret0, _ := ret[0].(*accessanalyzer.ListFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAnalyzerClientMockRecorder) ListFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFindings", reflect.TypeOf((*MockAnalyzerClient)(nil).ListFindings), varargs...)
}
