package mocks

import (
	context "context"
	reflect "reflect"

	elasticsearchservice "github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	gomock "github.com/golang/mock/gomock"
)

type MockElasticSearch struct {
	ctrl		*gomock.Controller
	recorder	*MockElasticSearchMockRecorder
}

type MockElasticSearchMockRecorder struct {
	mock *MockElasticSearch
}

func NewMockElasticSearch(ctrl *gomock.Controller) *MockElasticSearch {
	mock := &MockElasticSearch{ctrl: ctrl}
	mock.recorder = &MockElasticSearchMockRecorder{mock}
	return mock
}

func (m *MockElasticSearch) EXPECT() *MockElasticSearchMockRecorder {
	return m.recorder
}

func (m *MockElasticSearch) DescribeElasticsearchDomain(arg0 context.Context, arg1 *elasticsearchservice.DescribeElasticsearchDomainInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.DescribeElasticsearchDomainOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeElasticsearchDomain", varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.DescribeElasticsearchDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticSearchMockRecorder) DescribeElasticsearchDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeElasticsearchDomain", reflect.TypeOf((*MockElasticSearch)(nil).DescribeElasticsearchDomain), varargs...)
}

func (m *MockElasticSearch) ListDomainNames(arg0 context.Context, arg1 *elasticsearchservice.ListDomainNamesInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListDomainNamesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDomainNames", varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListDomainNamesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticSearchMockRecorder) ListDomainNames(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDomainNames", reflect.TypeOf((*MockElasticSearch)(nil).ListDomainNames), varargs...)
}

func (m *MockElasticSearch) ListTags(arg0 context.Context, arg1 *elasticsearchservice.ListTagsInput, arg2 ...func(*elasticsearchservice.Options)) (*elasticsearchservice.ListTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTags", varargs...)
	ret0, _ := ret[0].(*elasticsearchservice.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElasticSearchMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockElasticSearch)(nil).ListTags), varargs...)
}
