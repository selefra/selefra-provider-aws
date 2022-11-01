package mocks

import (
	context "context"
	reflect "reflect"

	dax "github.com/aws/aws-sdk-go-v2/service/dax"
	gomock "github.com/golang/mock/gomock"
)

type MockDAXClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDAXClientMockRecorder
}

type MockDAXClientMockRecorder struct {
	mock *MockDAXClient
}

func NewMockDAXClient(ctrl *gomock.Controller) *MockDAXClient {
	mock := &MockDAXClient{ctrl: ctrl}
	mock.recorder = &MockDAXClientMockRecorder{mock}
	return mock
}

func (m *MockDAXClient) EXPECT() *MockDAXClientMockRecorder {
	return m.recorder
}

func (m *MockDAXClient) DescribeClusters(arg0 context.Context, arg1 *dax.DescribeClustersInput, arg2 ...func(*dax.Options)) (*dax.DescribeClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusters", varargs...)
	ret0, _ := ret[0].(*dax.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDAXClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusters", reflect.TypeOf((*MockDAXClient)(nil).DescribeClusters), varargs...)
}

func (m *MockDAXClient) ListTags(arg0 context.Context, arg1 *dax.ListTagsInput, arg2 ...func(*dax.Options)) (*dax.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTags", varargs...)
	ret0, _ := ret[0].(*dax.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDAXClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockDAXClient)(nil).ListTags), varargs...)
}
