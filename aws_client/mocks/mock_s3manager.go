package mocks

import (
	context "context"
	reflect "reflect"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/golang/mock/gomock"
)

type MockS3ManagerClient struct {
	ctrl		*gomock.Controller
	recorder	*MockS3ManagerClientMockRecorder
}

type MockS3ManagerClientMockRecorder struct {
	mock *MockS3ManagerClient
}

func NewMockS3ManagerClient(ctrl *gomock.Controller) *MockS3ManagerClient {
	mock := &MockS3ManagerClient{ctrl: ctrl}
	mock.recorder = &MockS3ManagerClientMockRecorder{mock}
	return mock
}

func (m *MockS3ManagerClient) EXPECT() *MockS3ManagerClientMockRecorder {
	return m.recorder
}

func (m *MockS3ManagerClient) GetBucketRegion(arg0 context.Context, arg1 string, arg2 ...func(*s3.Options)) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketRegion", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ManagerClientMockRecorder) GetBucketRegion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketRegion", reflect.TypeOf((*MockS3ManagerClient)(nil).GetBucketRegion), varargs...)
}
