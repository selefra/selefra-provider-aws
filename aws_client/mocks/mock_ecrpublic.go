package mocks

import (
	context "context"
	reflect "reflect"

	ecrpublic "github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	gomock "github.com/golang/mock/gomock"
)

type MockEcrPublicClient struct {
	ctrl		*gomock.Controller
	recorder	*MockEcrPublicClientMockRecorder
}

type MockEcrPublicClientMockRecorder struct {
	mock *MockEcrPublicClient
}

func NewMockEcrPublicClient(ctrl *gomock.Controller) *MockEcrPublicClient {
	mock := &MockEcrPublicClient{ctrl: ctrl}
	mock.recorder = &MockEcrPublicClientMockRecorder{mock}
	return mock
}

func (m *MockEcrPublicClient) EXPECT() *MockEcrPublicClientMockRecorder {
	return m.recorder
}

func (m *MockEcrPublicClient) DescribeImageTags(arg0 context.Context, arg1 *ecrpublic.DescribeImageTagsInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImageTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImageTags", varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeImageTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrPublicClientMockRecorder) DescribeImageTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImageTags", reflect.TypeOf((*MockEcrPublicClient)(nil).DescribeImageTags), varargs...)
}

func (m *MockEcrPublicClient) DescribeImages(arg0 context.Context, arg1 *ecrpublic.DescribeImagesInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImages", varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrPublicClientMockRecorder) DescribeImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImages", reflect.TypeOf((*MockEcrPublicClient)(nil).DescribeImages), varargs...)
}

func (m *MockEcrPublicClient) DescribeRegistries(arg0 context.Context, arg1 *ecrpublic.DescribeRegistriesInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRegistriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRegistries", varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeRegistriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrPublicClientMockRecorder) DescribeRegistries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRegistries", reflect.TypeOf((*MockEcrPublicClient)(nil).DescribeRegistries), varargs...)
}

func (m *MockEcrPublicClient) DescribeRepositories(arg0 context.Context, arg1 *ecrpublic.DescribeRepositoriesInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRepositoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRepositories", varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeRepositoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrPublicClientMockRecorder) DescribeRepositories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRepositories", reflect.TypeOf((*MockEcrPublicClient)(nil).DescribeRepositories), varargs...)
}

func (m *MockEcrPublicClient) GetRepositoryPolicy(arg0 context.Context, arg1 *ecrpublic.GetRepositoryPolicyInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRepositoryPolicy", varargs...)
	ret0, _ := ret[0].(*ecrpublic.GetRepositoryPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrPublicClientMockRecorder) GetRepositoryPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryPolicy", reflect.TypeOf((*MockEcrPublicClient)(nil).GetRepositoryPolicy), varargs...)
}

func (m *MockEcrPublicClient) ListTagsForResource(arg0 context.Context, arg1 *ecrpublic.ListTagsForResourceInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*ecrpublic.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEcrPublicClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockEcrPublicClient)(nil).ListTagsForResource), varargs...)
}
