package mocks

import (
	context "context"
	reflect "reflect"

	elasticbeanstalk "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	gomock "github.com/golang/mock/gomock"
)

type MockElasticbeanstalkClient struct {
	ctrl		*gomock.Controller
	recorder	*MockElasticbeanstalkClientMockRecorder
}

type MockElasticbeanstalkClientMockRecorder struct {
	mock *MockElasticbeanstalkClient
}

func NewMockElasticbeanstalkClient(ctrl *gomock.Controller) *MockElasticbeanstalkClient {
	mock := &MockElasticbeanstalkClient{ctrl: ctrl}
	mock.recorder = &MockElasticbeanstalkClientMockRecorder{mock}
	return mock
}

func (m *MockElasticbeanstalkClient) EXPECT() *MockElasticbeanstalkClientMockRecorder {
	return m.recorder
}

func (m *MockElasticbeanstalkClient) DescribeApplicationVersions(arg0 context.Context, arg1 *elasticbeanstalk.DescribeApplicationVersionsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeApplicationVersions", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeApplicationVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeApplicationVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeApplicationVersions", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeApplicationVersions), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeApplications(arg0 context.Context, arg1 *elasticbeanstalk.DescribeApplicationsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeApplicationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeApplications", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeApplications", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeApplications), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeConfigurationOptions(arg0 context.Context, arg1 *elasticbeanstalk.DescribeConfigurationOptionsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConfigurationOptions", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeConfigurationOptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeConfigurationOptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConfigurationOptions", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeConfigurationOptions), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeConfigurationSettings(arg0 context.Context, arg1 *elasticbeanstalk.DescribeConfigurationSettingsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeConfigurationSettings", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeConfigurationSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeConfigurationSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeConfigurationSettings", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeConfigurationSettings), varargs...)
}

func (m *MockElasticbeanstalkClient) DescribeEnvironments(arg0 context.Context, arg1 *elasticbeanstalk.DescribeEnvironmentsInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.DescribeEnvironmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEnvironments", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.DescribeEnvironmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) DescribeEnvironments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEnvironments", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).DescribeEnvironments), varargs...)
}

func (m *MockElasticbeanstalkClient) ListTagsForResource(arg0 context.Context, arg1 *elasticbeanstalk.ListTagsForResourceInput, arg2 ...func(*elasticbeanstalk.Options)) (*elasticbeanstalk.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*elasticbeanstalk.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticbeanstalkClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockElasticbeanstalkClient)(nil).ListTagsForResource), varargs...)
}
