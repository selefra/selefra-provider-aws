package mocks

import (
	context "context"
	reflect "reflect"

	appsync "github.com/aws/aws-sdk-go-v2/service/appsync"
	gomock "github.com/golang/mock/gomock"
)

type MockAppSyncClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAppSyncClientMockRecorder
}

type MockAppSyncClientMockRecorder struct {
	mock *MockAppSyncClient
}

func NewMockAppSyncClient(ctrl *gomock.Controller) *MockAppSyncClient {
	mock := &MockAppSyncClient{ctrl: ctrl}
	mock.recorder = &MockAppSyncClientMockRecorder{mock}
	return mock
}

func (m *MockAppSyncClient) EXPECT() *MockAppSyncClientMockRecorder {
	return m.recorder
}

func (m *MockAppSyncClient) ListGraphqlApis(arg0 context.Context, arg1 *appsync.ListGraphqlApisInput, arg2 ...func(*appsync.Options)) (*appsync.ListGraphqlApisOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGraphqlApis", varargs...)
	ret0, _ := ret[0].(*appsync.ListGraphqlApisOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAppSyncClientMockRecorder) ListGraphqlApis(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGraphqlApis", reflect.TypeOf((*MockAppSyncClient)(nil).ListGraphqlApis), varargs...)
}
