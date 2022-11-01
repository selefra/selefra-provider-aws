package mocks

import (
	context "context"
	reflect "reflect"

	cloudhsmv2 "github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudHSMV2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockCloudHSMV2ClientMockRecorder
}

type MockCloudHSMV2ClientMockRecorder struct {
	mock *MockCloudHSMV2Client
}

func NewMockCloudHSMV2Client(ctrl *gomock.Controller) *MockCloudHSMV2Client {
	mock := &MockCloudHSMV2Client{ctrl: ctrl}
	mock.recorder = &MockCloudHSMV2ClientMockRecorder{mock}
	return mock
}

func (m *MockCloudHSMV2Client) EXPECT() *MockCloudHSMV2ClientMockRecorder {
	return m.recorder
}

func (m *MockCloudHSMV2Client) DescribeBackups(arg0 context.Context, arg1 *cloudhsmv2.DescribeBackupsInput, arg2 ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeBackupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBackups", varargs...)
	ret0, _ := ret[0].(*cloudhsmv2.DescribeBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudHSMV2ClientMockRecorder) DescribeBackups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackups", reflect.TypeOf((*MockCloudHSMV2Client)(nil).DescribeBackups), varargs...)
}

func (m *MockCloudHSMV2Client) DescribeClusters(arg0 context.Context, arg1 *cloudhsmv2.DescribeClustersInput, arg2 ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusters", varargs...)
	ret0, _ := ret[0].(*cloudhsmv2.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudHSMV2ClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusters", reflect.TypeOf((*MockCloudHSMV2Client)(nil).DescribeClusters), varargs...)
}
