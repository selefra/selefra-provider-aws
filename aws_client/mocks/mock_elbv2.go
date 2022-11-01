package mocks

import (
	context "context"
	reflect "reflect"

	elasticloadbalancingv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	gomock "github.com/golang/mock/gomock"
)

type MockElbV2Client struct {
	ctrl		*gomock.Controller
	recorder	*MockElbV2ClientMockRecorder
}

type MockElbV2ClientMockRecorder struct {
	mock *MockElbV2Client
}

func NewMockElbV2Client(ctrl *gomock.Controller) *MockElbV2Client {
	mock := &MockElbV2Client{ctrl: ctrl}
	mock.recorder = &MockElbV2ClientMockRecorder{mock}
	return mock
}

func (m *MockElbV2Client) EXPECT() *MockElbV2ClientMockRecorder {
	return m.recorder
}

func (m *MockElbV2Client) DescribeListenerCertificates(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeListenerCertificatesInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeListenerCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeListenerCertificates", varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeListenerCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElbV2ClientMockRecorder) DescribeListenerCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeListenerCertificates", reflect.TypeOf((*MockElbV2Client)(nil).DescribeListenerCertificates), varargs...)
}

func (m *MockElbV2Client) DescribeListeners(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeListenersInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeListenersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeListeners", varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeListenersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElbV2ClientMockRecorder) DescribeListeners(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeListeners", reflect.TypeOf((*MockElbV2Client)(nil).DescribeListeners), varargs...)
}

func (m *MockElbV2Client) DescribeLoadBalancerAttributes(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeLoadBalancerAttributesInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoadBalancerAttributes", varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeLoadBalancerAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElbV2ClientMockRecorder) DescribeLoadBalancerAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancerAttributes", reflect.TypeOf((*MockElbV2Client)(nil).DescribeLoadBalancerAttributes), varargs...)
}

func (m *MockElbV2Client) DescribeLoadBalancers(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeLoadBalancersInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoadBalancers", varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElbV2ClientMockRecorder) DescribeLoadBalancers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancers", reflect.TypeOf((*MockElbV2Client)(nil).DescribeLoadBalancers), varargs...)
}

func (m *MockElbV2Client) DescribeTags(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTagsInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTags", varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElbV2ClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTags", reflect.TypeOf((*MockElbV2Client)(nil).DescribeTags), varargs...)
}

func (m *MockElbV2Client) DescribeTargetGroups(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTargetGroupsInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeTargetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTargetGroups", varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTargetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElbV2ClientMockRecorder) DescribeTargetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTargetGroups", reflect.TypeOf((*MockElbV2Client)(nil).DescribeTargetGroups), varargs...)
}

func (m *MockElbV2Client) DescribeTargetHealth(arg0 context.Context, arg1 *elasticloadbalancingv2.DescribeTargetHealthInput, arg2 ...func(*elasticloadbalancingv2.Options)) (*elasticloadbalancingv2.DescribeTargetHealthOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTargetHealth", varargs...)
	ret0, _ := ret[0].(*elasticloadbalancingv2.DescribeTargetHealthOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElbV2ClientMockRecorder) DescribeTargetHealth(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTargetHealth", reflect.TypeOf((*MockElbV2Client)(nil).DescribeTargetHealth), varargs...)
}
