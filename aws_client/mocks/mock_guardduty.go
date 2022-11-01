package mocks

import (
	context "context"
	reflect "reflect"

	guardduty "github.com/aws/aws-sdk-go-v2/service/guardduty"
	gomock "github.com/golang/mock/gomock"
)

type MockGuardDutyClient struct {
	ctrl		*gomock.Controller
	recorder	*MockGuardDutyClientMockRecorder
}

type MockGuardDutyClientMockRecorder struct {
	mock *MockGuardDutyClient
}

func NewMockGuardDutyClient(ctrl *gomock.Controller) *MockGuardDutyClient {
	mock := &MockGuardDutyClient{ctrl: ctrl}
	mock.recorder = &MockGuardDutyClientMockRecorder{mock}
	return mock
}

func (m *MockGuardDutyClient) EXPECT() *MockGuardDutyClientMockRecorder {
	return m.recorder
}

func (m *MockGuardDutyClient) GetDetector(arg0 context.Context, arg1 *guardduty.GetDetectorInput, arg2 ...func(*guardduty.Options)) (*guardduty.GetDetectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDetector", varargs...)
	ret0, _ := ret[0].(*guardduty.GetDetectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuardDutyClientMockRecorder) GetDetector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetector", reflect.TypeOf((*MockGuardDutyClient)(nil).GetDetector), varargs...)
}

func (m *MockGuardDutyClient) ListDetectors(arg0 context.Context, arg1 *guardduty.ListDetectorsInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListDetectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDetectors", varargs...)
	ret0, _ := ret[0].(*guardduty.ListDetectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuardDutyClientMockRecorder) ListDetectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDetectors", reflect.TypeOf((*MockGuardDutyClient)(nil).ListDetectors), varargs...)
}

func (m *MockGuardDutyClient) ListMembers(arg0 context.Context, arg1 *guardduty.ListMembersInput, arg2 ...func(*guardduty.Options)) (*guardduty.ListMembersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMembers", varargs...)
	ret0, _ := ret[0].(*guardduty.ListMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGuardDutyClientMockRecorder) ListMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockGuardDutyClient)(nil).ListMembers), varargs...)
}
