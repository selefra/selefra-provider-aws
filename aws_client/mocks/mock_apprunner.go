package mocks

import (
	context "context"
	reflect "reflect"

	apprunner "github.com/aws/aws-sdk-go-v2/service/apprunner"
	gomock "github.com/golang/mock/gomock"
)

type MockAppRunnerClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAppRunnerClientMockRecorder
}

type MockAppRunnerClientMockRecorder struct {
	mock *MockAppRunnerClient
}

func NewMockAppRunnerClient(ctrl *gomock.Controller) *MockAppRunnerClient {
	mock := &MockAppRunnerClient{ctrl: ctrl}
	mock.recorder = &MockAppRunnerClientMockRecorder{mock}
	return mock
}

func (m *MockAppRunnerClient) EXPECT() *MockAppRunnerClientMockRecorder {
	return m.recorder
}

func (m *MockAppRunnerClient) DescribeAutoScalingConfiguration(arg0 context.Context, arg1 *apprunner.DescribeAutoScalingConfigurationInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeAutoScalingConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAutoScalingConfiguration", varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeAutoScalingConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) DescribeAutoScalingConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAutoScalingConfiguration", reflect.TypeOf((*MockAppRunnerClient)(nil).DescribeAutoScalingConfiguration), varargs...)
}

func (m *MockAppRunnerClient) DescribeCustomDomains(arg0 context.Context, arg1 *apprunner.DescribeCustomDomainsInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeCustomDomainsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCustomDomains", varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeCustomDomainsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) DescribeCustomDomains(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCustomDomains", reflect.TypeOf((*MockAppRunnerClient)(nil).DescribeCustomDomains), varargs...)
}

func (m *MockAppRunnerClient) DescribeObservabilityConfiguration(arg0 context.Context, arg1 *apprunner.DescribeObservabilityConfigurationInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeObservabilityConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeObservabilityConfiguration", varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeObservabilityConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) DescribeObservabilityConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeObservabilityConfiguration", reflect.TypeOf((*MockAppRunnerClient)(nil).DescribeObservabilityConfiguration), varargs...)
}

func (m *MockAppRunnerClient) DescribeService(arg0 context.Context, arg1 *apprunner.DescribeServiceInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeServiceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeService", varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) DescribeService(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeService", reflect.TypeOf((*MockAppRunnerClient)(nil).DescribeService), varargs...)
}

func (m *MockAppRunnerClient) DescribeVpcConnector(arg0 context.Context, arg1 *apprunner.DescribeVpcConnectorInput, arg2 ...func(*apprunner.Options)) (*apprunner.DescribeVpcConnectorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVpcConnector", varargs...)
	ret0, _ := ret[0].(*apprunner.DescribeVpcConnectorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) DescribeVpcConnector(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVpcConnector", reflect.TypeOf((*MockAppRunnerClient)(nil).DescribeVpcConnector), varargs...)
}

func (m *MockAppRunnerClient) ListAutoScalingConfigurations(arg0 context.Context, arg1 *apprunner.ListAutoScalingConfigurationsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListAutoScalingConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAutoScalingConfigurations", varargs...)
	ret0, _ := ret[0].(*apprunner.ListAutoScalingConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) ListAutoScalingConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAutoScalingConfigurations", reflect.TypeOf((*MockAppRunnerClient)(nil).ListAutoScalingConfigurations), varargs...)
}

func (m *MockAppRunnerClient) ListConnections(arg0 context.Context, arg1 *apprunner.ListConnectionsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConnections", varargs...)
	ret0, _ := ret[0].(*apprunner.ListConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) ListConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnections", reflect.TypeOf((*MockAppRunnerClient)(nil).ListConnections), varargs...)
}

func (m *MockAppRunnerClient) ListObservabilityConfigurations(arg0 context.Context, arg1 *apprunner.ListObservabilityConfigurationsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListObservabilityConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListObservabilityConfigurations", varargs...)
	ret0, _ := ret[0].(*apprunner.ListObservabilityConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) ListObservabilityConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObservabilityConfigurations", reflect.TypeOf((*MockAppRunnerClient)(nil).ListObservabilityConfigurations), varargs...)
}

func (m *MockAppRunnerClient) ListOperations(arg0 context.Context, arg1 *apprunner.ListOperationsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListOperationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOperations", varargs...)
	ret0, _ := ret[0].(*apprunner.ListOperationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) ListOperations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockAppRunnerClient)(nil).ListOperations), varargs...)
}

func (m *MockAppRunnerClient) ListServices(arg0 context.Context, arg1 *apprunner.ListServicesInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*apprunner.ListServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) ListServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockAppRunnerClient)(nil).ListServices), varargs...)
}

func (m *MockAppRunnerClient) ListTagsForResource(arg0 context.Context, arg1 *apprunner.ListTagsForResourceInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*apprunner.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAppRunnerClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockAppRunnerClient) ListVpcConnectors(arg0 context.Context, arg1 *apprunner.ListVpcConnectorsInput, arg2 ...func(*apprunner.Options)) (*apprunner.ListVpcConnectorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVpcConnectors", varargs...)
	ret0, _ := ret[0].(*apprunner.ListVpcConnectorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppRunnerClientMockRecorder) ListVpcConnectors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVpcConnectors", reflect.TypeOf((*MockAppRunnerClient)(nil).ListVpcConnectors), varargs...)
}
