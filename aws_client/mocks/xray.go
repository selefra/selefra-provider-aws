package mocks

import (
	context "context"
	reflect "reflect"

	xray "github.com/aws/aws-sdk-go-v2/service/xray"
	gomock "github.com/golang/mock/gomock"
)

type MockXrayClient struct {
	ctrl		*gomock.Controller
	recorder	*MockXrayClientMockRecorder
}

type MockXrayClientMockRecorder struct {
	mock *MockXrayClient
}

func NewMockXrayClient(ctrl *gomock.Controller) *MockXrayClient {
	mock := &MockXrayClient{ctrl: ctrl}
	mock.recorder = &MockXrayClientMockRecorder{mock}
	return mock
}

func (m *MockXrayClient) EXPECT() *MockXrayClientMockRecorder {
	return m.recorder
}

func (m *MockXrayClient) GetEncryptionConfig(arg0 context.Context, arg1 *xray.GetEncryptionConfigInput, arg2 ...func(*xray.Options)) (*xray.GetEncryptionConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEncryptionConfig", varargs...)
	ret0, _ := ret[0].(*xray.GetEncryptionConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetEncryptionConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEncryptionConfig", reflect.TypeOf((*MockXrayClient)(nil).GetEncryptionConfig), varargs...)
}

func (m *MockXrayClient) GetGroups(arg0 context.Context, arg1 *xray.GetGroupsInput, arg2 ...func(*xray.Options)) (*xray.GetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroups", varargs...)
	ret0, _ := ret[0].(*xray.GetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*MockXrayClient)(nil).GetGroups), varargs...)
}

func (m *MockXrayClient) GetSamplingRules(arg0 context.Context, arg1 *xray.GetSamplingRulesInput, arg2 ...func(*xray.Options)) (*xray.GetSamplingRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSamplingRules", varargs...)
	ret0, _ := ret[0].(*xray.GetSamplingRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) GetSamplingRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamplingRules", reflect.TypeOf((*MockXrayClient)(nil).GetSamplingRules), varargs...)
}

func (m *MockXrayClient) ListTagsForResource(arg0 context.Context, arg1 *xray.ListTagsForResourceInput, arg2 ...func(*xray.Options)) (*xray.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*xray.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockXrayClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockXrayClient)(nil).ListTagsForResource), varargs...)
}
