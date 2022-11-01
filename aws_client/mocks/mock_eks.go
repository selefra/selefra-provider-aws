package mocks

import (
	context "context"
	reflect "reflect"

	eks "github.com/aws/aws-sdk-go-v2/service/eks"
	gomock "github.com/golang/mock/gomock"
)

type MockEksClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEksClientMockRecorder
}

type MockEksClientMockRecorder struct {
	mock *MockEksClient
}

func NewMockEksClient(ctrl *gomock.Controller) *MockEksClient {
	mock := &MockEksClient{ctrl: ctrl}
	mock.recorder = &MockEksClientMockRecorder{mock}
	return mock
}

func (m *MockEksClient) EXPECT() *MockEksClientMockRecorder {
	return m.recorder
}

func (m *MockEksClient) DescribeCluster(arg0 context.Context, arg1 *eks.DescribeClusterInput, arg2 ...func(*eks.Options)) (*eks.DescribeClusterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCluster", varargs...)
	ret0, _ := ret[0].(*eks.DescribeClusterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) DescribeCluster(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCluster", reflect.TypeOf((*MockEksClient)(nil).DescribeCluster), varargs...)
}

func (m *MockEksClient) ListClusters(arg0 context.Context, arg1 *eks.ListClustersInput, arg2 ...func(*eks.Options)) (*eks.ListClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].(*eks.ListClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEksClientMockRecorder) ListClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockEksClient)(nil).ListClusters), varargs...)
}
