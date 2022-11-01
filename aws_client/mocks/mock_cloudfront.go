package mocks

import (
	context "context"
	reflect "reflect"

	cloudfront "github.com/aws/aws-sdk-go-v2/service/cloudfront"
	gomock "github.com/golang/mock/gomock"
)

type MockCloudfrontClient struct {
	ctrl		*gomock.Controller
	recorder	*MockCloudfrontClientMockRecorder
}

type MockCloudfrontClientMockRecorder struct {
	mock *MockCloudfrontClient
}

func NewMockCloudfrontClient(ctrl *gomock.Controller) *MockCloudfrontClient {
	mock := &MockCloudfrontClient{ctrl: ctrl}
	mock.recorder = &MockCloudfrontClientMockRecorder{mock}
	return mock
}

func (m *MockCloudfrontClient) EXPECT() *MockCloudfrontClientMockRecorder {
	return m.recorder
}

func (m *MockCloudfrontClient) GetDistribution(arg0 context.Context, arg1 *cloudfront.GetDistributionInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.GetDistributionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDistribution", varargs...)
	ret0, _ := ret[0].(*cloudfront.GetDistributionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) GetDistribution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDistribution", reflect.TypeOf((*MockCloudfrontClient)(nil).GetDistribution), varargs...)
}

func (m *MockCloudfrontClient) ListCachePolicies(arg0 context.Context, arg1 *cloudfront.ListCachePoliciesInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListCachePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCachePolicies", varargs...)
	ret0, _ := ret[0].(*cloudfront.ListCachePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListCachePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCachePolicies", reflect.TypeOf((*MockCloudfrontClient)(nil).ListCachePolicies), varargs...)
}

func (m *MockCloudfrontClient) ListDistributions(arg0 context.Context, arg1 *cloudfront.ListDistributionsInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDistributions", varargs...)
	ret0, _ := ret[0].(*cloudfront.ListDistributionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListDistributions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDistributions", reflect.TypeOf((*MockCloudfrontClient)(nil).ListDistributions), varargs...)
}

func (m *MockCloudfrontClient) ListDistributionsByWebACLId(arg0 context.Context, arg1 *cloudfront.ListDistributionsByWebACLIdInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListDistributionsByWebACLIdOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDistributionsByWebACLId", varargs...)
	ret0, _ := ret[0].(*cloudfront.ListDistributionsByWebACLIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListDistributionsByWebACLId(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDistributionsByWebACLId", reflect.TypeOf((*MockCloudfrontClient)(nil).ListDistributionsByWebACLId), varargs...)
}

func (m *MockCloudfrontClient) ListTagsForResource(arg0 context.Context, arg1 *cloudfront.ListTagsForResourceInput, arg2 ...func(*cloudfront.Options)) (*cloudfront.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*cloudfront.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCloudfrontClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockCloudfrontClient)(nil).ListTagsForResource), varargs...)
}
