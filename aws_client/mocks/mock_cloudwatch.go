package mocks

import (
	context "context"
	reflect "reflect"

	cloudwatch "github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudwatchClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCloudwatchClientMockRecorder
}

type MockCloudwatchClientMockRecorder struct {
	mock *MockCloudwatchClient
}

func NewMockCloudwatchClient(ctrl *gomock.Controller) *MockCloudwatchClient {
	mock := &MockCloudwatchClient{ctrl: ctrl}
	mock.recorder = &MockCloudwatchClientMockRecorder{mock}
	return mock
}

func (m *MockCloudwatchClient) EXPECT() *MockCloudwatchClientMockRecorder {
	return m.recorder
}

func (m *MockCloudwatchClient) DescribeAlarms(arg0 context.Context, arg1 *cloudwatch.DescribeAlarmsInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.DescribeAlarmsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAlarms", varargs...)
	ret0, _ := ret[0].(*cloudwatch.DescribeAlarmsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) DescribeAlarms(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAlarms", reflect.TypeOf((*MockCloudwatchClient)(nil).DescribeAlarms), varargs...)
}

func (m *MockCloudwatchClient) ListTagsForResource(arg0 context.Context, arg1 *cloudwatch.ListTagsForResourceInput, arg2 ...func(*cloudwatch.Options)) (*cloudwatch.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*cloudwatch.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudwatchClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockCloudwatchClient)(nil).ListTagsForResource), varargs...)
}
