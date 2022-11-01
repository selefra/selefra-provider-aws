package mocks

import (
	context "context"
	reflect "reflect"

	cognitoidentityprovider "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	gomock "github.com/golang/mock/gomock"
)

type MockCognitoUserPoolsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCognitoUserPoolsClientMockRecorder
}

type MockCognitoUserPoolsClientMockRecorder struct {
	mock *MockCognitoUserPoolsClient
}

func NewMockCognitoUserPoolsClient(ctrl *gomock.Controller) *MockCognitoUserPoolsClient {
	mock := &MockCognitoUserPoolsClient{ctrl: ctrl}
	mock.recorder = &MockCognitoUserPoolsClientMockRecorder{mock}
	return mock
}

func (m *MockCognitoUserPoolsClient) EXPECT() *MockCognitoUserPoolsClientMockRecorder {
	return m.recorder
}

func (m *MockCognitoUserPoolsClient) DescribeIdentityProvider(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeIdentityProviderInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeIdentityProvider", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeIdentityProviderOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoUserPoolsClientMockRecorder) DescribeIdentityProvider(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeIdentityProvider", reflect.TypeOf((*MockCognitoUserPoolsClient)(nil).DescribeIdentityProvider), varargs...)
}

func (m *MockCognitoUserPoolsClient) DescribeUserPool(arg0 context.Context, arg1 *cognitoidentityprovider.DescribeUserPoolInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUserPool", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.DescribeUserPoolOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoUserPoolsClientMockRecorder) DescribeUserPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUserPool", reflect.TypeOf((*MockCognitoUserPoolsClient)(nil).DescribeUserPool), varargs...)
}

func (m *MockCognitoUserPoolsClient) ListIdentityProviders(arg0 context.Context, arg1 *cognitoidentityprovider.ListIdentityProvidersInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListIdentityProvidersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIdentityProviders", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListIdentityProvidersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoUserPoolsClientMockRecorder) ListIdentityProviders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIdentityProviders", reflect.TypeOf((*MockCognitoUserPoolsClient)(nil).ListIdentityProviders), varargs...)
}

func (m *MockCognitoUserPoolsClient) ListUserPools(arg0 context.Context, arg1 *cognitoidentityprovider.ListUserPoolsInput, arg2 ...func(*cognitoidentityprovider.Options)) (*cognitoidentityprovider.ListUserPoolsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUserPools", varargs...)
	ret0, _ := ret[0].(*cognitoidentityprovider.ListUserPoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoUserPoolsClientMockRecorder) ListUserPools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserPools", reflect.TypeOf((*MockCognitoUserPoolsClient)(nil).ListUserPools), varargs...)
}
