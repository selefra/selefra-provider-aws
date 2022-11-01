package mocks

import (
	context "context"
	reflect "reflect"

	neptune "github.com/aws/aws-sdk-go-v2/service/neptune"
	gomock "github.com/golang/mock/gomock"
)

type MockNeptuneClient struct {
	ctrl		*gomock.Controller
	recorder	*MockNeptuneClientMockRecorder
}

type MockNeptuneClientMockRecorder struct {
	mock *MockNeptuneClient
}

func NewMockNeptuneClient(ctrl *gomock.Controller) *MockNeptuneClient {
	mock := &MockNeptuneClient{ctrl: ctrl}
	mock.recorder = &MockNeptuneClientMockRecorder{mock}
	return mock
}

func (m *MockNeptuneClient) EXPECT() *MockNeptuneClientMockRecorder {
	return m.recorder
}

func (m *MockNeptuneClient) DescribeDBClusterEndpoints(arg0 context.Context, arg1 *neptune.DescribeDBClusterEndpointsInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeDBClusterEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterEndpoints", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeDBClusterEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeDBClusterEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterEndpoints", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeDBClusterEndpoints), varargs...)
}

func (m *MockNeptuneClient) DescribeDBClusterParameterGroups(arg0 context.Context, arg1 *neptune.DescribeDBClusterParameterGroupsInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeDBClusterParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterParameterGroups", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeDBClusterParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeDBClusterParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterParameterGroups", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeDBClusterParameterGroups), varargs...)
}

func (m *MockNeptuneClient) DescribeDBClusterParameters(arg0 context.Context, arg1 *neptune.DescribeDBClusterParametersInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeDBClusterParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterParameters", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeDBClusterParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeDBClusterParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterParameters", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeDBClusterParameters), varargs...)
}

func (m *MockNeptuneClient) DescribeDBClusterSnapshotAttributes(arg0 context.Context, arg1 *neptune.DescribeDBClusterSnapshotAttributesInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterSnapshotAttributes", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeDBClusterSnapshotAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeDBClusterSnapshotAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterSnapshotAttributes", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeDBClusterSnapshotAttributes), varargs...)
}

func (m *MockNeptuneClient) DescribeDBClusterSnapshots(arg0 context.Context, arg1 *neptune.DescribeDBClusterSnapshotsInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeDBClusterSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterSnapshots", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeDBClusterSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeDBClusterSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterSnapshots", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeDBClusterSnapshots), varargs...)
}

func (m *MockNeptuneClient) DescribeDBClusters(arg0 context.Context, arg1 *neptune.DescribeDBClustersInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeDBClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusters", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeDBClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeDBClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusters", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeDBClusters), varargs...)
}

func (m *MockNeptuneClient) DescribeDBInstances(arg0 context.Context, arg1 *neptune.DescribeDBInstancesInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeDBInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBInstances", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeDBInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeDBInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBInstances", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeDBInstances), varargs...)
}

func (m *MockNeptuneClient) DescribeDBParameterGroups(arg0 context.Context, arg1 *neptune.DescribeDBParameterGroupsInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeDBParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBParameterGroups", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeDBParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeDBParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBParameterGroups", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeDBParameterGroups), varargs...)
}

func (m *MockNeptuneClient) DescribeDBParameters(arg0 context.Context, arg1 *neptune.DescribeDBParametersInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeDBParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBParameters", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeDBParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeDBParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBParameters", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeDBParameters), varargs...)
}

func (m *MockNeptuneClient) DescribeDBSubnetGroups(arg0 context.Context, arg1 *neptune.DescribeDBSubnetGroupsInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeDBSubnetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBSubnetGroups", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeDBSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeDBSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBSubnetGroups", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeDBSubnetGroups), varargs...)
}

func (m *MockNeptuneClient) DescribeEventSubscriptions(arg0 context.Context, arg1 *neptune.DescribeEventSubscriptionsInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeEventSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEventSubscriptions", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeEventSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeEventSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEventSubscriptions", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeEventSubscriptions), varargs...)
}

func (m *MockNeptuneClient) DescribeGlobalClusters(arg0 context.Context, arg1 *neptune.DescribeGlobalClustersInput, arg2 ...func(*neptune.Options)) (*neptune.DescribeGlobalClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeGlobalClusters", varargs...)
	ret0, _ := ret[0].(*neptune.DescribeGlobalClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) DescribeGlobalClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeGlobalClusters", reflect.TypeOf((*MockNeptuneClient)(nil).DescribeGlobalClusters), varargs...)
}

func (m *MockNeptuneClient) ListTagsForResource(arg0 context.Context, arg1 *neptune.ListTagsForResourceInput, arg2 ...func(*neptune.Options)) (*neptune.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*neptune.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNeptuneClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockNeptuneClient)(nil).ListTagsForResource), varargs...)
}
