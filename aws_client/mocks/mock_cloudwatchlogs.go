package mocks

import (
	context "context"
	reflect "reflect"

	cloudwatchlogs "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudwatchLogsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCloudwatchLogsClientMockRecorder
}

type MockCloudwatchLogsClientMockRecorder struct {
	mock *MockCloudwatchLogsClient
}

func NewMockCloudwatchLogsClient(ctrl *gomock.Controller) *MockCloudwatchLogsClient {
	mock := &MockCloudwatchLogsClient{ctrl: ctrl}
	mock.recorder = &MockCloudwatchLogsClientMockRecorder{mock}
	return mock
}

func (m *MockCloudwatchLogsClient) EXPECT() *MockCloudwatchLogsClientMockRecorder {
	return m.recorder
}

func (m *MockCloudwatchLogsClient) DescribeLogGroups(arg0 context.Context, arg1 *cloudwatchlogs.DescribeLogGroupsInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLogGroups", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeLogGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchLogsClientMockRecorder) DescribeLogGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLogGroups", reflect.TypeOf((*MockCloudwatchLogsClient)(nil).DescribeLogGroups), varargs...)
}

func (m *MockCloudwatchLogsClient) DescribeMetricFilters(arg0 context.Context, arg1 *cloudwatchlogs.DescribeMetricFiltersInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.DescribeMetricFiltersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMetricFilters", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.DescribeMetricFiltersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchLogsClientMockRecorder) DescribeMetricFilters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMetricFilters", reflect.TypeOf((*MockCloudwatchLogsClient)(nil).DescribeMetricFilters), varargs...)
}

func (m *MockCloudwatchLogsClient) ListTagsLogGroup(arg0 context.Context, arg1 *cloudwatchlogs.ListTagsLogGroupInput, arg2 ...func(*cloudwatchlogs.Options)) (*cloudwatchlogs.ListTagsLogGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsLogGroup", varargs...)
	ret0, _ := ret[0].(*cloudwatchlogs.ListTagsLogGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchLogsClientMockRecorder) ListTagsLogGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsLogGroup", reflect.TypeOf((*MockCloudwatchLogsClient)(nil).ListTagsLogGroup), varargs...)
}
