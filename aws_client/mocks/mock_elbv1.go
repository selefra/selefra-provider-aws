package mocks

import (
	context "context"
	reflect "reflect"

	elasticloadbalancing "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	gomock "github.com/golang/mock/gomock"
)

type MockElbV1Client struct {
	ctrl		*gomock.Controller
	recorder	*MockElbV1ClientMockRecorder
}

type MockElbV1ClientMockRecorder struct {
	mock *MockElbV1Client
}

func NewMockElbV1Client(ctrl *gomock.Controller) *MockElbV1Client {
	mock := &MockElbV1Client{ctrl: ctrl}
	mock.recorder = &MockElbV1ClientMockRecorder{mock}
	return mock
}

func (m *MockElbV1Client) EXPECT() *MockElbV1ClientMockRecorder {
	return m.recorder
}

func (m *MockElbV1Client) DescribeLoadBalancerAttributes(arg0 context.Context, arg1 *elasticloadbalancing.DescribeLoadBalancerAttributesInput, arg2 ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancerAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoadBalancerAttributes", varargs...)
	ret0, _ := ret[0].(*elasticloadbalancing.DescribeLoadBalancerAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElbV1ClientMockRecorder) DescribeLoadBalancerAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancerAttributes", reflect.TypeOf((*MockElbV1Client)(nil).DescribeLoadBalancerAttributes), varargs...)
}

func (m *MockElbV1Client) DescribeLoadBalancerPolicies(arg0 context.Context, arg1 *elasticloadbalancing.DescribeLoadBalancerPoliciesInput, arg2 ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancerPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoadBalancerPolicies", varargs...)
	ret0, _ := ret[0].(*elasticloadbalancing.DescribeLoadBalancerPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElbV1ClientMockRecorder) DescribeLoadBalancerPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancerPolicies", reflect.TypeOf((*MockElbV1Client)(nil).DescribeLoadBalancerPolicies), varargs...)
}

func (m *MockElbV1Client) DescribeLoadBalancers(arg0 context.Context, arg1 *elasticloadbalancing.DescribeLoadBalancersInput, arg2 ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoadBalancers", varargs...)
	ret0, _ := ret[0].(*elasticloadbalancing.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElbV1ClientMockRecorder) DescribeLoadBalancers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancers", reflect.TypeOf((*MockElbV1Client)(nil).DescribeLoadBalancers), varargs...)
}

func (m *MockElbV1Client) DescribeTags(arg0 context.Context, arg1 *elasticloadbalancing.DescribeTagsInput, arg2 ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTags", varargs...)
	ret0, _ := ret[0].(*elasticloadbalancing.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElbV1ClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTags", reflect.TypeOf((*MockElbV1Client)(nil).DescribeTags), varargs...)
}
