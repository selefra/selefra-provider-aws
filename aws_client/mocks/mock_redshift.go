package mocks

import (
	context "context"
	reflect "reflect"

	redshift "github.com/aws/aws-sdk-go-v2/service/redshift"
	gomock "github.com/golang/mock/gomock"
)

type MockRedshiftClient struct {
	ctrl		*gomock.Controller
	recorder	*MockRedshiftClientMockRecorder
}

type MockRedshiftClientMockRecorder struct {
	mock *MockRedshiftClient
}

func NewMockRedshiftClient(ctrl *gomock.Controller) *MockRedshiftClient {
	mock := &MockRedshiftClient{ctrl: ctrl}
	mock.recorder = &MockRedshiftClientMockRecorder{mock}
	return mock
}

func (m *MockRedshiftClient) EXPECT() *MockRedshiftClientMockRecorder {
	return m.recorder
}

func (m *MockRedshiftClient) DescribeClusterParameters(arg0 context.Context, arg1 *redshift.DescribeClusterParametersInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClusterParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusterParameters", varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClusterParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusterParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusterParameters", reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusterParameters), varargs...)
}

func (m *MockRedshiftClient) DescribeClusterSnapshots(arg0 context.Context, arg1 *redshift.DescribeClusterSnapshotsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClusterSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusterSnapshots", varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClusterSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusterSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusterSnapshots", reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusterSnapshots), varargs...)
}

func (m *MockRedshiftClient) DescribeClusterSubnetGroups(arg0 context.Context, arg1 *redshift.DescribeClusterSubnetGroupsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClusterSubnetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusterSubnetGroups", varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClusterSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusterSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusterSubnetGroups", reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusterSubnetGroups), varargs...)
}

func (m *MockRedshiftClient) DescribeClusters(arg0 context.Context, arg1 *redshift.DescribeClustersInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusters", varargs...)
	ret0, _ := ret[0].(*redshift.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusters", reflect.TypeOf((*MockRedshiftClient)(nil).DescribeClusters), varargs...)
}

func (m *MockRedshiftClient) DescribeEventSubscriptions(arg0 context.Context, arg1 *redshift.DescribeEventSubscriptionsInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeEventSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEventSubscriptions", varargs...)
	ret0, _ := ret[0].(*redshift.DescribeEventSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeEventSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEventSubscriptions", reflect.TypeOf((*MockRedshiftClient)(nil).DescribeEventSubscriptions), varargs...)
}

func (m *MockRedshiftClient) DescribeLoggingStatus(arg0 context.Context, arg1 *redshift.DescribeLoggingStatusInput, arg2 ...func(*redshift.Options)) (*redshift.DescribeLoggingStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoggingStatus", varargs...)
	ret0, _ := ret[0].(*redshift.DescribeLoggingStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedshiftClientMockRecorder) DescribeLoggingStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoggingStatus", reflect.TypeOf((*MockRedshiftClient)(nil).DescribeLoggingStatus), varargs...)
}
