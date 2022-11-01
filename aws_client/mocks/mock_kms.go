package mocks

import (
	context "context"
	reflect "reflect"

	kms "github.com/aws/aws-sdk-go-v2/service/kms"
	gomock "github.com/golang/mock/gomock"
)

type MockKmsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockKmsClientMockRecorder
}

type MockKmsClientMockRecorder struct {
	mock *MockKmsClient
}

func NewMockKmsClient(ctrl *gomock.Controller) *MockKmsClient {
	mock := &MockKmsClient{ctrl: ctrl}
	mock.recorder = &MockKmsClientMockRecorder{mock}
	return mock
}

func (m *MockKmsClient) EXPECT() *MockKmsClientMockRecorder {
	return m.recorder
}

func (m *MockKmsClient) DescribeKey(arg0 context.Context, arg1 *kms.DescribeKeyInput, arg2 ...func(*kms.Options)) (*kms.DescribeKeyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeKey", varargs...)
	ret0, _ := ret[0].(*kms.DescribeKeyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) DescribeKey(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeKey", reflect.TypeOf((*MockKmsClient)(nil).DescribeKey), varargs...)
}

func (m *MockKmsClient) GetKeyRotationStatus(arg0 context.Context, arg1 *kms.GetKeyRotationStatusInput, arg2 ...func(*kms.Options)) (*kms.GetKeyRotationStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetKeyRotationStatus", varargs...)
	ret0, _ := ret[0].(*kms.GetKeyRotationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) GetKeyRotationStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeyRotationStatus", reflect.TypeOf((*MockKmsClient)(nil).GetKeyRotationStatus), varargs...)
}

func (m *MockKmsClient) ListAliases(arg0 context.Context, arg1 *kms.ListAliasesInput, arg2 ...func(*kms.Options)) (*kms.ListAliasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAliases", varargs...)
	ret0, _ := ret[0].(*kms.ListAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) ListAliases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAliases", reflect.TypeOf((*MockKmsClient)(nil).ListAliases), varargs...)
}

func (m *MockKmsClient) ListKeys(arg0 context.Context, arg1 *kms.ListKeysInput, arg2 ...func(*kms.Options)) (*kms.ListKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListKeys", varargs...)
	ret0, _ := ret[0].(*kms.ListKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) ListKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockKmsClient)(nil).ListKeys), varargs...)
}

func (m *MockKmsClient) ListResourceTags(arg0 context.Context, arg1 *kms.ListResourceTagsInput, arg2 ...func(*kms.Options)) (*kms.ListResourceTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourceTags", varargs...)
	ret0, _ := ret[0].(*kms.ListResourceTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKmsClientMockRecorder) ListResourceTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceTags", reflect.TypeOf((*MockKmsClient)(nil).ListResourceTags), varargs...)
}
