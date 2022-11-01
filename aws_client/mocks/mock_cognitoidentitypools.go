package mocks

import (
	context "context"
	reflect "reflect"

	cognitoidentity "github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	gomock "github.com/golang/mock/gomock"
)

type MockCognitoIdentityPoolsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCognitoIdentityPoolsClientMockRecorder
}

type MockCognitoIdentityPoolsClientMockRecorder struct {
	mock *MockCognitoIdentityPoolsClient
}

func NewMockCognitoIdentityPoolsClient(ctrl *gomock.Controller) *MockCognitoIdentityPoolsClient {
	mock := &MockCognitoIdentityPoolsClient{ctrl: ctrl}
	mock.recorder = &MockCognitoIdentityPoolsClientMockRecorder{mock}
	return mock
}

func (m *MockCognitoIdentityPoolsClient) EXPECT() *MockCognitoIdentityPoolsClientMockRecorder {
	return m.recorder
}

func (m *MockCognitoIdentityPoolsClient) DescribeIdentityPool(arg0 context.Context, arg1 *cognitoidentity.DescribeIdentityPoolInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.DescribeIdentityPoolOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeIdentityPool", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.DescribeIdentityPoolOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoIdentityPoolsClientMockRecorder) DescribeIdentityPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeIdentityPool", reflect.TypeOf((*MockCognitoIdentityPoolsClient)(nil).DescribeIdentityPool), varargs...)
}

func (m *MockCognitoIdentityPoolsClient) ListIdentityPools(arg0 context.Context, arg1 *cognitoidentity.ListIdentityPoolsInput, arg2 ...func(*cognitoidentity.Options)) (*cognitoidentity.ListIdentityPoolsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListIdentityPools", varargs...)
	ret0, _ := ret[0].(*cognitoidentity.ListIdentityPoolsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCognitoIdentityPoolsClientMockRecorder) ListIdentityPools(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIdentityPools", reflect.TypeOf((*MockCognitoIdentityPoolsClient)(nil).ListIdentityPools), varargs...)
}
