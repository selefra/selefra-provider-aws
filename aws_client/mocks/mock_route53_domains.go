package mocks

import (
	context "context"
	reflect "reflect"

	route53domains "github.com/aws/aws-sdk-go-v2/service/route53domains"
	gomock "github.com/golang/mock/gomock"
)

type MockRoute53DomainsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockRoute53DomainsClientMockRecorder
}

type MockRoute53DomainsClientMockRecorder struct {
	mock *MockRoute53DomainsClient
}

func NewMockRoute53DomainsClient(ctrl *gomock.Controller) *MockRoute53DomainsClient {
	mock := &MockRoute53DomainsClient{ctrl: ctrl}
	mock.recorder = &MockRoute53DomainsClientMockRecorder{mock}
	return mock
}

func (m *MockRoute53DomainsClient) EXPECT() *MockRoute53DomainsClientMockRecorder {
	return m.recorder
}

func (m *MockRoute53DomainsClient) GetDomainDetail(arg0 context.Context, arg1 *route53domains.GetDomainDetailInput, arg2 ...func(*route53domains.Options)) (*route53domains.GetDomainDetailOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDomainDetail", varargs...)
	ret0, _ := ret[0].(*route53domains.GetDomainDetailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53DomainsClientMockRecorder) GetDomainDetail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainDetail", reflect.TypeOf((*MockRoute53DomainsClient)(nil).GetDomainDetail), varargs...)
}

func (m *MockRoute53DomainsClient) ListDomains(arg0 context.Context, arg1 *route53domains.ListDomainsInput, arg2 ...func(*route53domains.Options)) (*route53domains.ListDomainsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDomains", varargs...)
	ret0, _ := ret[0].(*route53domains.ListDomainsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53DomainsClientMockRecorder) ListDomains(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDomains", reflect.TypeOf((*MockRoute53DomainsClient)(nil).ListDomains), varargs...)
}

func (m *MockRoute53DomainsClient) ListTagsForDomain(arg0 context.Context, arg1 *route53domains.ListTagsForDomainInput, arg2 ...func(*route53domains.Options)) (*route53domains.ListTagsForDomainOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForDomain", varargs...)
	ret0, _ := ret[0].(*route53domains.ListTagsForDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRoute53DomainsClientMockRecorder) ListTagsForDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForDomain", reflect.TypeOf((*MockRoute53DomainsClient)(nil).ListTagsForDomain), varargs...)
}
