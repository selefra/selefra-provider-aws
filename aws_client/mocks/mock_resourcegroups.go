package mocks

import (
	context "context"
	reflect "reflect"

	resourcegroups "github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	gomock "github.com/golang/mock/gomock"
)

type MockResourceGroupsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockResourceGroupsClientMockRecorder
}

type MockResourceGroupsClientMockRecorder struct {
	mock *MockResourceGroupsClient
}

func NewMockResourceGroupsClient(ctrl *gomock.Controller) *MockResourceGroupsClient {
	mock := &MockResourceGroupsClient{ctrl: ctrl}
	mock.recorder = &MockResourceGroupsClientMockRecorder{mock}
	return mock
}

func (m *MockResourceGroupsClient) EXPECT() *MockResourceGroupsClientMockRecorder {
	return m.recorder
}

func (m *MockResourceGroupsClient) GetGroup(arg0 context.Context, arg1 *resourcegroups.GetGroupInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroup", varargs...)
	ret0, _ := ret[0].(*resourcegroups.GetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceGroupsClientMockRecorder) GetGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockResourceGroupsClient)(nil).GetGroup), varargs...)
}

func (m *MockResourceGroupsClient) GetGroupQuery(arg0 context.Context, arg1 *resourcegroups.GetGroupQueryInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.GetGroupQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroupQuery", varargs...)
	ret0, _ := ret[0].(*resourcegroups.GetGroupQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceGroupsClientMockRecorder) GetGroupQuery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupQuery", reflect.TypeOf((*MockResourceGroupsClient)(nil).GetGroupQuery), varargs...)
}

func (m *MockResourceGroupsClient) GetTags(arg0 context.Context, arg1 *resourcegroups.GetTagsInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.GetTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTags", varargs...)
	ret0, _ := ret[0].(*resourcegroups.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceGroupsClientMockRecorder) GetTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockResourceGroupsClient)(nil).GetTags), varargs...)
}

func (m *MockResourceGroupsClient) ListGroups(arg0 context.Context, arg1 *resourcegroups.ListGroupsInput, arg2 ...func(*resourcegroups.Options)) (*resourcegroups.ListGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGroups", varargs...)
	ret0, _ := ret[0].(*resourcegroups.ListGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourceGroupsClientMockRecorder) ListGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroups", reflect.TypeOf((*MockResourceGroupsClient)(nil).ListGroups), varargs...)
}
