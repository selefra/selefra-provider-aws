package mocks

import (
	context "context"
	reflect "reflect"

	eventbridge "github.com/aws/aws-sdk-go-v2/service/eventbridge"
	gomock "github.com/golang/mock/gomock"
)

type MockEventBridgeClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEventBridgeClientMockRecorder
}

type MockEventBridgeClientMockRecorder struct {
	mock *MockEventBridgeClient
}

func NewMockEventBridgeClient(ctrl *gomock.Controller) *MockEventBridgeClient {
	mock := &MockEventBridgeClient{ctrl: ctrl}
	mock.recorder = &MockEventBridgeClientMockRecorder{mock}
	return mock
}

func (m *MockEventBridgeClient) EXPECT() *MockEventBridgeClientMockRecorder {
	return m.recorder
}

func (m *MockEventBridgeClient) ListEventBuses(arg0 context.Context, arg1 *eventbridge.ListEventBusesInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListEventBusesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEventBuses", varargs...)
	ret0, _ := ret[0].(*eventbridge.ListEventBusesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventBridgeClientMockRecorder) ListEventBuses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventBuses", reflect.TypeOf((*MockEventBridgeClient)(nil).ListEventBuses), varargs...)
}

func (m *MockEventBridgeClient) ListRules(arg0 context.Context, arg1 *eventbridge.ListRulesInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRules", varargs...)
	ret0, _ := ret[0].(*eventbridge.ListRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventBridgeClientMockRecorder) ListRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRules", reflect.TypeOf((*MockEventBridgeClient)(nil).ListRules), varargs...)
}

func (m *MockEventBridgeClient) ListTagsForResource(arg0 context.Context, arg1 *eventbridge.ListTagsForResourceInput, arg2 ...func(*eventbridge.Options)) (*eventbridge.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*eventbridge.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventBridgeClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockEventBridgeClient)(nil).ListTagsForResource), varargs...)
}
