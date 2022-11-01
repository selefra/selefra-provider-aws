package mocks

import (
	context "context"
	reflect "reflect"

	ecs "github.com/aws/aws-sdk-go-v2/service/ecs"
	gomock "github.com/golang/mock/gomock"
)

type MockEcsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEcsClientMockRecorder
}

type MockEcsClientMockRecorder struct {
	mock *MockEcsClient
}

func NewMockEcsClient(ctrl *gomock.Controller) *MockEcsClient {
	mock := &MockEcsClient{ctrl: ctrl}
	mock.recorder = &MockEcsClientMockRecorder{mock}
	return mock
}

func (m *MockEcsClient) EXPECT() *MockEcsClientMockRecorder {
	return m.recorder
}

func (m *MockEcsClient) DescribeClusters(arg0 context.Context, arg1 *ecs.DescribeClustersInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusters", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusters", reflect.TypeOf((*MockEcsClient)(nil).DescribeClusters), varargs...)
}

func (m *MockEcsClient) DescribeContainerInstances(arg0 context.Context, arg1 *ecs.DescribeContainerInstancesInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeContainerInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeContainerInstances", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeContainerInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeContainerInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeContainerInstances", reflect.TypeOf((*MockEcsClient)(nil).DescribeContainerInstances), varargs...)
}

func (m *MockEcsClient) DescribeServices(arg0 context.Context, arg1 *ecs.DescribeServicesInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeServices", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServices", reflect.TypeOf((*MockEcsClient)(nil).DescribeServices), varargs...)
}

func (m *MockEcsClient) DescribeTaskDefinition(arg0 context.Context, arg1 *ecs.DescribeTaskDefinitionInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeTaskDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTaskDefinition", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeTaskDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeTaskDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTaskDefinition", reflect.TypeOf((*MockEcsClient)(nil).DescribeTaskDefinition), varargs...)
}

func (m *MockEcsClient) DescribeTasks(arg0 context.Context, arg1 *ecs.DescribeTasksInput, arg2 ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTasks", varargs...)
	ret0, _ := ret[0].(*ecs.DescribeTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) DescribeTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTasks", reflect.TypeOf((*MockEcsClient)(nil).DescribeTasks), varargs...)
}

func (m *MockEcsClient) ListClusters(arg0 context.Context, arg1 *ecs.ListClustersInput, arg2 ...func(*ecs.Options)) (*ecs.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].(*ecs.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockEcsClient)(nil).ListClusters), varargs...)
}

func (m *MockEcsClient) ListContainerInstances(arg0 context.Context, arg1 *ecs.ListContainerInstancesInput, arg2 ...func(*ecs.Options)) (*ecs.ListContainerInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListContainerInstances", varargs...)
	ret0, _ := ret[0].(*ecs.ListContainerInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListContainerInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainerInstances", reflect.TypeOf((*MockEcsClient)(nil).ListContainerInstances), varargs...)
}

func (m *MockEcsClient) ListServices(arg0 context.Context, arg1 *ecs.ListServicesInput, arg2 ...func(*ecs.Options)) (*ecs.ListServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*ecs.ListServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockEcsClient)(nil).ListServices), varargs...)
}

func (m *MockEcsClient) ListTaskDefinitions(arg0 context.Context, arg1 *ecs.ListTaskDefinitionsInput, arg2 ...func(*ecs.Options)) (*ecs.ListTaskDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTaskDefinitions", varargs...)
	ret0, _ := ret[0].(*ecs.ListTaskDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListTaskDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTaskDefinitions", reflect.TypeOf((*MockEcsClient)(nil).ListTaskDefinitions), varargs...)
}

func (m *MockEcsClient) ListTasks(arg0 context.Context, arg1 *ecs.ListTasksInput, arg2 ...func(*ecs.Options)) (*ecs.ListTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTasks", varargs...)
	ret0, _ := ret[0].(*ecs.ListTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcsClientMockRecorder) ListTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockEcsClient)(nil).ListTasks), varargs...)
}
