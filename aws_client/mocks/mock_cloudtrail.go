package mocks

import (
	context "context"
	reflect "reflect"

	cloudtrail "github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudtrailClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCloudtrailClientMockRecorder
}

type MockCloudtrailClientMockRecorder struct {
	mock *MockCloudtrailClient
}

func NewMockCloudtrailClient(ctrl *gomock.Controller) *MockCloudtrailClient {
	mock := &MockCloudtrailClient{ctrl: ctrl}
	mock.recorder = &MockCloudtrailClientMockRecorder{mock}
	return mock
}

func (m *MockCloudtrailClient) EXPECT() *MockCloudtrailClientMockRecorder {
	return m.recorder
}

func (m *MockCloudtrailClient) DescribeTrails(arg0 context.Context, arg1 *cloudtrail.DescribeTrailsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.DescribeTrailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTrails", varargs...)
	ret0, _ := ret[0].(*cloudtrail.DescribeTrailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) DescribeTrails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTrails", reflect.TypeOf((*MockCloudtrailClient)(nil).DescribeTrails), varargs...)
}

func (m *MockCloudtrailClient) GetEventSelectors(arg0 context.Context, arg1 *cloudtrail.GetEventSelectorsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetEventSelectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEventSelectors", varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetEventSelectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) GetEventSelectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventSelectors", reflect.TypeOf((*MockCloudtrailClient)(nil).GetEventSelectors), varargs...)
}

func (m *MockCloudtrailClient) GetTrailStatus(arg0 context.Context, arg1 *cloudtrail.GetTrailStatusInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.GetTrailStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTrailStatus", varargs...)
	ret0, _ := ret[0].(*cloudtrail.GetTrailStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) GetTrailStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrailStatus", reflect.TypeOf((*MockCloudtrailClient)(nil).GetTrailStatus), varargs...)
}

func (m *MockCloudtrailClient) ListTags(arg0 context.Context, arg1 *cloudtrail.ListTagsInput, arg2 ...func(*cloudtrail.Options)) (*cloudtrail.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTags", varargs...)
	ret0, _ := ret[0].(*cloudtrail.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudtrailClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockCloudtrailClient)(nil).ListTags), varargs...)
}
