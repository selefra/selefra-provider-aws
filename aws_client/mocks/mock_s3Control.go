package mocks

import (
	context "context"
	reflect "reflect"

	s3control "github.com/aws/aws-sdk-go-v2/service/s3control"
	gomock "github.com/golang/mock/gomock"
)

type MockS3ControlClient struct {
	ctrl		*gomock.Controller
	recorder	*MockS3ControlClientMockRecorder
}

type MockS3ControlClientMockRecorder struct {
	mock *MockS3ControlClient
}

func NewMockS3ControlClient(ctrl *gomock.Controller) *MockS3ControlClient {
	mock := &MockS3ControlClient{ctrl: ctrl}
	mock.recorder = &MockS3ControlClientMockRecorder{mock}
	return mock
}

func (m *MockS3ControlClient) EXPECT() *MockS3ControlClientMockRecorder {
	return m.recorder
}

func (m *MockS3ControlClient) GetPublicAccessBlock(arg0 context.Context, arg1 *s3control.GetPublicAccessBlockInput, arg2 ...func(*s3control.Options)) (*s3control.GetPublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPublicAccessBlock", varargs...)
	ret0, _ := ret[0].(*s3control.GetPublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ControlClientMockRecorder) GetPublicAccessBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicAccessBlock", reflect.TypeOf((*MockS3ControlClient)(nil).GetPublicAccessBlock), varargs...)
}
