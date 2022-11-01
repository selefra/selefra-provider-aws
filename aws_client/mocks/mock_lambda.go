package mocks

import (
	context "context"
	reflect "reflect"

	lambda "github.com/aws/aws-sdk-go-v2/service/lambda"
	gomock "github.com/golang/mock/gomock"
)

type MockLambdaClient struct {
	ctrl		*gomock.Controller
	recorder	*MockLambdaClientMockRecorder
}

type MockLambdaClientMockRecorder struct {
	mock *MockLambdaClient
}

func NewMockLambdaClient(ctrl *gomock.Controller) *MockLambdaClient {
	mock := &MockLambdaClient{ctrl: ctrl}
	mock.recorder = &MockLambdaClientMockRecorder{mock}
	return mock
}

func (m *MockLambdaClient) EXPECT() *MockLambdaClientMockRecorder {
	return m.recorder
}

func (m *MockLambdaClient) GetCodeSigningConfig(arg0 context.Context, arg1 *lambda.GetCodeSigningConfigInput, arg2 ...func(*lambda.Options)) (*lambda.GetCodeSigningConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCodeSigningConfig", varargs...)
	ret0, _ := ret[0].(*lambda.GetCodeSigningConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetCodeSigningConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCodeSigningConfig", reflect.TypeOf((*MockLambdaClient)(nil).GetCodeSigningConfig), varargs...)
}

func (m *MockLambdaClient) GetFunction(arg0 context.Context, arg1 *lambda.GetFunctionInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFunction", varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunction", reflect.TypeOf((*MockLambdaClient)(nil).GetFunction), varargs...)
}

func (m *MockLambdaClient) GetFunctionCodeSigningConfig(arg0 context.Context, arg1 *lambda.GetFunctionCodeSigningConfigInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionCodeSigningConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFunctionCodeSigningConfig", varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionCodeSigningConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetFunctionCodeSigningConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunctionCodeSigningConfig", reflect.TypeOf((*MockLambdaClient)(nil).GetFunctionCodeSigningConfig), varargs...)
}

func (m *MockLambdaClient) GetFunctionUrlConfig(arg0 context.Context, arg1 *lambda.GetFunctionUrlConfigInput, arg2 ...func(*lambda.Options)) (*lambda.GetFunctionUrlConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFunctionUrlConfig", varargs...)
	ret0, _ := ret[0].(*lambda.GetFunctionUrlConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetFunctionUrlConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunctionUrlConfig", reflect.TypeOf((*MockLambdaClient)(nil).GetFunctionUrlConfig), varargs...)
}

func (m *MockLambdaClient) GetLayerVersionPolicy(arg0 context.Context, arg1 *lambda.GetLayerVersionPolicyInput, arg2 ...func(*lambda.Options)) (*lambda.GetLayerVersionPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLayerVersionPolicy", varargs...)
	ret0, _ := ret[0].(*lambda.GetLayerVersionPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetLayerVersionPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLayerVersionPolicy", reflect.TypeOf((*MockLambdaClient)(nil).GetLayerVersionPolicy), varargs...)
}

func (m *MockLambdaClient) GetPolicy(arg0 context.Context, arg1 *lambda.GetPolicyInput, arg2 ...func(*lambda.Options)) (*lambda.GetPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPolicy", varargs...)
	ret0, _ := ret[0].(*lambda.GetPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) GetPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicy", reflect.TypeOf((*MockLambdaClient)(nil).GetPolicy), varargs...)
}

func (m *MockLambdaClient) ListAliases(arg0 context.Context, arg1 *lambda.ListAliasesInput, arg2 ...func(*lambda.Options)) (*lambda.ListAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAliases", varargs...)
	ret0, _ := ret[0].(*lambda.ListAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAliases", reflect.TypeOf((*MockLambdaClient)(nil).ListAliases), varargs...)
}

func (m *MockLambdaClient) ListEventSourceMappings(arg0 context.Context, arg1 *lambda.ListEventSourceMappingsInput, arg2 ...func(*lambda.Options)) (*lambda.ListEventSourceMappingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEventSourceMappings", varargs...)
	ret0, _ := ret[0].(*lambda.ListEventSourceMappingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListEventSourceMappings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEventSourceMappings", reflect.TypeOf((*MockLambdaClient)(nil).ListEventSourceMappings), varargs...)
}

func (m *MockLambdaClient) ListFunctionEventInvokeConfigs(arg0 context.Context, arg1 *lambda.ListFunctionEventInvokeConfigsInput, arg2 ...func(*lambda.Options)) (*lambda.ListFunctionEventInvokeConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFunctionEventInvokeConfigs", varargs...)
	ret0, _ := ret[0].(*lambda.ListFunctionEventInvokeConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListFunctionEventInvokeConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFunctionEventInvokeConfigs", reflect.TypeOf((*MockLambdaClient)(nil).ListFunctionEventInvokeConfigs), varargs...)
}

func (m *MockLambdaClient) ListFunctions(arg0 context.Context, arg1 *lambda.ListFunctionsInput, arg2 ...func(*lambda.Options)) (*lambda.ListFunctionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFunctions", varargs...)
	ret0, _ := ret[0].(*lambda.ListFunctionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListFunctions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFunctions", reflect.TypeOf((*MockLambdaClient)(nil).ListFunctions), varargs...)
}

func (m *MockLambdaClient) ListLayerVersions(arg0 context.Context, arg1 *lambda.ListLayerVersionsInput, arg2 ...func(*lambda.Options)) (*lambda.ListLayerVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLayerVersions", varargs...)
	ret0, _ := ret[0].(*lambda.ListLayerVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListLayerVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLayerVersions", reflect.TypeOf((*MockLambdaClient)(nil).ListLayerVersions), varargs...)
}

func (m *MockLambdaClient) ListLayers(arg0 context.Context, arg1 *lambda.ListLayersInput, arg2 ...func(*lambda.Options)) (*lambda.ListLayersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLayers", varargs...)
	ret0, _ := ret[0].(*lambda.ListLayersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListLayers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLayers", reflect.TypeOf((*MockLambdaClient)(nil).ListLayers), varargs...)
}

func (m *MockLambdaClient) ListProvisionedConcurrencyConfigs(arg0 context.Context, arg1 *lambda.ListProvisionedConcurrencyConfigsInput, arg2 ...func(*lambda.Options)) (*lambda.ListProvisionedConcurrencyConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProvisionedConcurrencyConfigs", varargs...)
	ret0, _ := ret[0].(*lambda.ListProvisionedConcurrencyConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListProvisionedConcurrencyConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProvisionedConcurrencyConfigs", reflect.TypeOf((*MockLambdaClient)(nil).ListProvisionedConcurrencyConfigs), varargs...)
}

func (m *MockLambdaClient) ListVersionsByFunction(arg0 context.Context, arg1 *lambda.ListVersionsByFunctionInput, arg2 ...func(*lambda.Options)) (*lambda.ListVersionsByFunctionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVersionsByFunction", varargs...)
	ret0, _ := ret[0].(*lambda.ListVersionsByFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLambdaClientMockRecorder) ListVersionsByFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVersionsByFunction", reflect.TypeOf((*MockLambdaClient)(nil).ListVersionsByFunction), varargs...)
}
