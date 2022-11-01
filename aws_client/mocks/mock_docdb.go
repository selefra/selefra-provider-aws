package mocks

import (
	context "context"
	reflect "reflect"

	docdb "github.com/aws/aws-sdk-go-v2/service/docdb"
	gomock "github.com/golang/mock/gomock"
)

type MockDocDBClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDocDBClientMockRecorder
}

type MockDocDBClientMockRecorder struct {
	mock *MockDocDBClient
}

func NewMockDocDBClient(ctrl *gomock.Controller) *MockDocDBClient {
	mock := &MockDocDBClient{ctrl: ctrl}
	mock.recorder = &MockDocDBClientMockRecorder{mock}
	return mock
}

func (m *MockDocDBClient) EXPECT() *MockDocDBClientMockRecorder {
	return m.recorder
}

func (m *MockDocDBClient) DescribeDBClusterSnapshotAttributes(arg0 context.Context, arg1 *docdb.DescribeDBClusterSnapshotAttributesInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterSnapshotAttributes", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClusterSnapshotAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocDBClientMockRecorder) DescribeDBClusterSnapshotAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterSnapshotAttributes", reflect.TypeOf((*MockDocDBClient)(nil).DescribeDBClusterSnapshotAttributes), varargs...)
}

func (m *MockDocDBClient) DescribeDBClusterSnapshots(arg0 context.Context, arg1 *docdb.DescribeDBClusterSnapshotsInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClusterSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterSnapshots", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClusterSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocDBClientMockRecorder) DescribeDBClusterSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterSnapshots", reflect.TypeOf((*MockDocDBClient)(nil).DescribeDBClusterSnapshots), varargs...)
}

func (m *MockDocDBClient) DescribeDBClusters(arg0 context.Context, arg1 *docdb.DescribeDBClustersInput, arg2 ...func(*docdb.Options)) (*docdb.DescribeDBClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusters", varargs...)
	ret0, _ := ret[0].(*docdb.DescribeDBClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocDBClientMockRecorder) DescribeDBClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusters", reflect.TypeOf((*MockDocDBClient)(nil).DescribeDBClusters), varargs...)
}

func (m *MockDocDBClient) ListTagsForResource(arg0 context.Context, arg1 *docdb.ListTagsForResourceInput, arg2 ...func(*docdb.Options)) (*docdb.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*docdb.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDocDBClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockDocDBClient)(nil).ListTagsForResource), varargs...)
}
