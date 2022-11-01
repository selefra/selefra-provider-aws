package mocks

import (
	context "context"
	reflect "reflect"

	acm "github.com/aws/aws-sdk-go-v2/service/acm"
	gomock "github.com/golang/mock/gomock"
)

type MockACMClient struct {
	ctrl		*gomock.Controller
	recorder	*MockACMClientMockRecorder
}

type MockACMClientMockRecorder struct {
	mock *MockACMClient
}

func NewMockACMClient(ctrl *gomock.Controller) *MockACMClient {
	mock := &MockACMClient{ctrl: ctrl}
	mock.recorder = &MockACMClientMockRecorder{mock}
	return mock
}

func (m *MockACMClient) EXPECT() *MockACMClientMockRecorder {
	return m.recorder
}

func (m *MockACMClient) DescribeCertificate(arg0 context.Context, arg1 *acm.DescribeCertificateInput, arg2 ...func(*acm.Options)) (*acm.DescribeCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCertificate", varargs...)
	ret0, _ := ret[0].(*acm.DescribeCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockACMClientMockRecorder) DescribeCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCertificate", reflect.TypeOf((*MockACMClient)(nil).DescribeCertificate), varargs...)
}

func (m *MockACMClient) ListCertificates(arg0 context.Context, arg1 *acm.ListCertificatesInput, arg2 ...func(*acm.Options)) (*acm.ListCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCertificates", varargs...)
	ret0, _ := ret[0].(*acm.ListCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockACMClientMockRecorder) ListCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificates", reflect.TypeOf((*MockACMClient)(nil).ListCertificates), varargs...)
}

func (m *MockACMClient) ListTagsForCertificate(arg0 context.Context, arg1 *acm.ListTagsForCertificateInput, arg2 ...func(*acm.Options)) (*acm.ListTagsForCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForCertificate", varargs...)
	ret0, _ := ret[0].(*acm.ListTagsForCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockACMClientMockRecorder) ListTagsForCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForCertificate", reflect.TypeOf((*MockACMClient)(nil).ListTagsForCertificate), varargs...)
}
