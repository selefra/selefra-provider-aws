package mocks

import (
	context "context"
	reflect "reflect"

	apigatewayv2 "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	gomock "github.com/golang/mock/gomock"
)

type MockApigatewayv2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockApigatewayv2ClientMockRecorder
}

type MockApigatewayv2ClientMockRecorder struct {
	mock *MockApigatewayv2Client
}

func NewMockApigatewayv2Client(ctrl *gomock.Controller) *MockApigatewayv2Client {
	mock := &MockApigatewayv2Client{ctrl: ctrl}
	mock.recorder = &MockApigatewayv2ClientMockRecorder{mock}
	return mock
}

func (m *MockApigatewayv2Client) EXPECT() *MockApigatewayv2ClientMockRecorder {
	return m.recorder
}

func (m *MockApigatewayv2Client) GetApiMappings(arg0 context.Context, arg1 *apigatewayv2.GetApiMappingsInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApiMappingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApiMappings", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetApiMappingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetApiMappings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiMappings", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetApiMappings), varargs...)
}

func (m *MockApigatewayv2Client) GetApis(arg0 context.Context, arg1 *apigatewayv2.GetApisInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetApisOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApis", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetApisOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetApis(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApis", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetApis), varargs...)
}

func (m *MockApigatewayv2Client) GetAuthorizers(arg0 context.Context, arg1 *apigatewayv2.GetAuthorizersInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetAuthorizersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAuthorizers", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetAuthorizersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetAuthorizers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizers", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetAuthorizers), varargs...)
}

func (m *MockApigatewayv2Client) GetDeployments(arg0 context.Context, arg1 *apigatewayv2.GetDeploymentsInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDeploymentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeployments", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetDeploymentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetDeployments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployments", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetDeployments), varargs...)
}

func (m *MockApigatewayv2Client) GetDomainNames(arg0 context.Context, arg1 *apigatewayv2.GetDomainNamesInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetDomainNamesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDomainNames", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetDomainNamesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetDomainNames(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainNames", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetDomainNames), varargs...)
}

func (m *MockApigatewayv2Client) GetIntegrationResponses(arg0 context.Context, arg1 *apigatewayv2.GetIntegrationResponsesInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationResponsesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIntegrationResponses", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetIntegrationResponsesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetIntegrationResponses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntegrationResponses", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetIntegrationResponses), varargs...)
}

func (m *MockApigatewayv2Client) GetIntegrations(arg0 context.Context, arg1 *apigatewayv2.GetIntegrationsInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetIntegrationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIntegrations", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetIntegrationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetIntegrations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntegrations", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetIntegrations), varargs...)
}

func (m *MockApigatewayv2Client) GetModelTemplate(arg0 context.Context, arg1 *apigatewayv2.GetModelTemplateInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelTemplateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetModelTemplate", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetModelTemplateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetModelTemplate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelTemplate", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetModelTemplate), varargs...)
}

func (m *MockApigatewayv2Client) GetModels(arg0 context.Context, arg1 *apigatewayv2.GetModelsInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetModelsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetModels", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetModelsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModels", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetModels), varargs...)
}

func (m *MockApigatewayv2Client) GetRouteResponses(arg0 context.Context, arg1 *apigatewayv2.GetRouteResponsesInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRouteResponsesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRouteResponses", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetRouteResponsesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetRouteResponses(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRouteResponses", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetRouteResponses), varargs...)
}

func (m *MockApigatewayv2Client) GetRoutes(arg0 context.Context, arg1 *apigatewayv2.GetRoutesInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetRoutesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRoutes", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetRoutesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetRoutes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoutes", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetRoutes), varargs...)
}

func (m *MockApigatewayv2Client) GetStages(arg0 context.Context, arg1 *apigatewayv2.GetStagesInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetStagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStages", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetStagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetStages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStages", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetStages), varargs...)
}

func (m *MockApigatewayv2Client) GetTags(arg0 context.Context, arg1 *apigatewayv2.GetTagsInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTags", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetTags), varargs...)
}

func (m *MockApigatewayv2Client) GetVpcLinks(arg0 context.Context, arg1 *apigatewayv2.GetVpcLinksInput, arg2 ...func(*apigatewayv2.Options)) (*apigatewayv2.GetVpcLinksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVpcLinks", varargs...)
	ret0, _ := ret[0].(*apigatewayv2.GetVpcLinksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockApigatewayv2ClientMockRecorder) GetVpcLinks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVpcLinks", reflect.TypeOf((*MockApigatewayv2Client)(nil).GetVpcLinks), varargs...)
}
