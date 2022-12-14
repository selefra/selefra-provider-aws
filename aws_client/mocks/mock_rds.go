package mocks

import (
	context "context"
	reflect "reflect"

	rds "github.com/aws/aws-sdk-go-v2/service/rds"
	gomock "github.com/golang/mock/gomock"
)

type MockRdsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockRdsClientMockRecorder
}

type MockRdsClientMockRecorder struct {
	mock *MockRdsClient
}

func NewMockRdsClient(ctrl *gomock.Controller) *MockRdsClient {
	mock := &MockRdsClient{ctrl: ctrl}
	mock.recorder = &MockRdsClientMockRecorder{mock}
	return mock
}

func (m *MockRdsClient) EXPECT() *MockRdsClientMockRecorder {
	return m.recorder
}

func (m *MockRdsClient) DescribeCertificates(arg0 context.Context, arg1 *rds.DescribeCertificatesInput, arg2 ...func(*rds.Options)) (*rds.DescribeCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCertificates", varargs...)
	ret0, _ := ret[0].(*rds.DescribeCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCertificates", reflect.TypeOf((*MockRdsClient)(nil).DescribeCertificates), varargs...)
}

func (m *MockRdsClient) DescribeDBClusterParameterGroups(arg0 context.Context, arg1 *rds.DescribeDBClusterParameterGroupsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClusterParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterParameterGroups", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClusterParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusterParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterParameterGroups", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusterParameterGroups), varargs...)
}

func (m *MockRdsClient) DescribeDBClusterParameters(arg0 context.Context, arg1 *rds.DescribeDBClusterParametersInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClusterParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterParameters", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClusterParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusterParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterParameters", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusterParameters), varargs...)
}

func (m *MockRdsClient) DescribeDBClusterSnapshotAttributes(arg0 context.Context, arg1 *rds.DescribeDBClusterSnapshotAttributesInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClusterSnapshotAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterSnapshotAttributes", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClusterSnapshotAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusterSnapshotAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterSnapshotAttributes", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusterSnapshotAttributes), varargs...)
}

func (m *MockRdsClient) DescribeDBClusterSnapshots(arg0 context.Context, arg1 *rds.DescribeDBClusterSnapshotsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClusterSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusterSnapshots", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClusterSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusterSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusterSnapshots", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusterSnapshots), varargs...)
}

func (m *MockRdsClient) DescribeDBClusters(arg0 context.Context, arg1 *rds.DescribeDBClustersInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBClusters", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBClusters", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBClusters), varargs...)
}

func (m *MockRdsClient) DescribeDBInstances(arg0 context.Context, arg1 *rds.DescribeDBInstancesInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBInstances", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBInstances", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBInstances), varargs...)
}

func (m *MockRdsClient) DescribeDBParameterGroups(arg0 context.Context, arg1 *rds.DescribeDBParameterGroupsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBParameterGroups", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBParameterGroups", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBParameterGroups), varargs...)
}

func (m *MockRdsClient) DescribeDBParameters(arg0 context.Context, arg1 *rds.DescribeDBParametersInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBParameters", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBParameters", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBParameters), varargs...)
}

func (m *MockRdsClient) DescribeDBSecurityGroups(arg0 context.Context, arg1 *rds.DescribeDBSecurityGroupsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBSecurityGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBSecurityGroups", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBSecurityGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBSecurityGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBSecurityGroups", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBSecurityGroups), varargs...)
}

func (m *MockRdsClient) DescribeDBSnapshotAttributes(arg0 context.Context, arg1 *rds.DescribeDBSnapshotAttributesInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBSnapshotAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBSnapshotAttributes", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBSnapshotAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBSnapshotAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBSnapshotAttributes", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBSnapshotAttributes), varargs...)
}

func (m *MockRdsClient) DescribeDBSnapshots(arg0 context.Context, arg1 *rds.DescribeDBSnapshotsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBSnapshots", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBSnapshots", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBSnapshots), varargs...)
}

func (m *MockRdsClient) DescribeDBSubnetGroups(arg0 context.Context, arg1 *rds.DescribeDBSubnetGroupsInput, arg2 ...func(*rds.Options)) (*rds.DescribeDBSubnetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDBSubnetGroups", varargs...)
	ret0, _ := ret[0].(*rds.DescribeDBSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeDBSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDBSubnetGroups", reflect.TypeOf((*MockRdsClient)(nil).DescribeDBSubnetGroups), varargs...)
}

func (m *MockRdsClient) DescribeEventSubscriptions(arg0 context.Context, arg1 *rds.DescribeEventSubscriptionsInput, arg2 ...func(*rds.Options)) (*rds.DescribeEventSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEventSubscriptions", varargs...)
	ret0, _ := ret[0].(*rds.DescribeEventSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) DescribeEventSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEventSubscriptions", reflect.TypeOf((*MockRdsClient)(nil).DescribeEventSubscriptions), varargs...)
}

func (m *MockRdsClient) ListTagsForResource(arg0 context.Context, arg1 *rds.ListTagsForResourceInput, arg2 ...func(*rds.Options)) (*rds.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*rds.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRdsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockRdsClient)(nil).ListTagsForResource), varargs...)
}
