package mocks

import (
	context "context"
	reflect "reflect"

	sagemaker "github.com/aws/aws-sdk-go-v2/service/sagemaker"
	gomock "github.com/golang/mock/gomock"
)

type MockSageMakerClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSageMakerClientMockRecorder
}

type MockSageMakerClientMockRecorder struct {
	mock *MockSageMakerClient
}

func NewMockSageMakerClient(ctrl *gomock.Controller) *MockSageMakerClient {
	mock := &MockSageMakerClient{ctrl: ctrl}
	mock.recorder = &MockSageMakerClientMockRecorder{mock}
	return mock
}

func (m *MockSageMakerClient) EXPECT() *MockSageMakerClientMockRecorder {
	return m.recorder
}

func (m *MockSageMakerClient) DescribeEndpointConfig(arg0 context.Context, arg1 *sagemaker.DescribeEndpointConfigInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeEndpointConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEndpointConfig", varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeEndpointConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSageMakerClientMockRecorder) DescribeEndpointConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEndpointConfig", reflect.TypeOf((*MockSageMakerClient)(nil).DescribeEndpointConfig), varargs...)
}

func (m *MockSageMakerClient) DescribeModel(arg0 context.Context, arg1 *sagemaker.DescribeModelInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeModelOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeModel", varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeModelOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSageMakerClientMockRecorder) DescribeModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeModel", reflect.TypeOf((*MockSageMakerClient)(nil).DescribeModel), varargs...)
}

func (m *MockSageMakerClient) DescribeNotebookInstance(arg0 context.Context, arg1 *sagemaker.DescribeNotebookInstanceInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeNotebookInstanceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNotebookInstance", varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeNotebookInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSageMakerClientMockRecorder) DescribeNotebookInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNotebookInstance", reflect.TypeOf((*MockSageMakerClient)(nil).DescribeNotebookInstance), varargs...)
}

func (m *MockSageMakerClient) DescribeTrainingJob(arg0 context.Context, arg1 *sagemaker.DescribeTrainingJobInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.DescribeTrainingJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTrainingJob", varargs...)
	ret0, _ := ret[0].(*sagemaker.DescribeTrainingJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSageMakerClientMockRecorder) DescribeTrainingJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTrainingJob", reflect.TypeOf((*MockSageMakerClient)(nil).DescribeTrainingJob), varargs...)
}

func (m *MockSageMakerClient) ListEndpointConfigs(arg0 context.Context, arg1 *sagemaker.ListEndpointConfigsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListEndpointConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEndpointConfigs", varargs...)
	ret0, _ := ret[0].(*sagemaker.ListEndpointConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSageMakerClientMockRecorder) ListEndpointConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEndpointConfigs", reflect.TypeOf((*MockSageMakerClient)(nil).ListEndpointConfigs), varargs...)
}

func (m *MockSageMakerClient) ListModels(arg0 context.Context, arg1 *sagemaker.ListModelsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListModelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListModels", varargs...)
	ret0, _ := ret[0].(*sagemaker.ListModelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSageMakerClientMockRecorder) ListModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListModels", reflect.TypeOf((*MockSageMakerClient)(nil).ListModels), varargs...)
}

func (m *MockSageMakerClient) ListNotebookInstances(arg0 context.Context, arg1 *sagemaker.ListNotebookInstancesInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListNotebookInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNotebookInstances", varargs...)
	ret0, _ := ret[0].(*sagemaker.ListNotebookInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSageMakerClientMockRecorder) ListNotebookInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNotebookInstances", reflect.TypeOf((*MockSageMakerClient)(nil).ListNotebookInstances), varargs...)
}

func (m *MockSageMakerClient) ListTags(arg0 context.Context, arg1 *sagemaker.ListTagsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTags", varargs...)
	ret0, _ := ret[0].(*sagemaker.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSageMakerClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockSageMakerClient)(nil).ListTags), varargs...)
}

func (m *MockSageMakerClient) ListTrainingJobs(arg0 context.Context, arg1 *sagemaker.ListTrainingJobsInput, arg2 ...func(*sagemaker.Options)) (*sagemaker.ListTrainingJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTrainingJobs", varargs...)
	ret0, _ := ret[0].(*sagemaker.ListTrainingJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSageMakerClientMockRecorder) ListTrainingJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrainingJobs", reflect.TypeOf((*MockSageMakerClient)(nil).ListTrainingJobs), varargs...)
}
