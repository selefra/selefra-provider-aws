package mocks

import (
	context "context"
	reflect "reflect"

	autoscaling "github.com/aws/aws-sdk-go-v2/service/autoscaling"
	gomock "github.com/golang/mock/gomock"
)

type MockAutoscalingClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAutoscalingClientMockRecorder
}

type MockAutoscalingClientMockRecorder struct {
	mock *MockAutoscalingClient
}

func NewMockAutoscalingClient(ctrl *gomock.Controller) *MockAutoscalingClient {
	mock := &MockAutoscalingClient{ctrl: ctrl}
	mock.recorder = &MockAutoscalingClientMockRecorder{mock}
	return mock
}

func (m *MockAutoscalingClient) EXPECT() *MockAutoscalingClientMockRecorder {
	return m.recorder
}

func (m *MockAutoscalingClient) DescribeAutoScalingGroups(arg0 context.Context, arg1 *autoscaling.DescribeAutoScalingGroupsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAutoScalingGroups", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeAutoScalingGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeAutoScalingGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAutoScalingGroups", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeAutoScalingGroups), varargs...)
}

func (m *MockAutoscalingClient) DescribeLaunchConfigurations(arg0 context.Context, arg1 *autoscaling.DescribeLaunchConfigurationsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLaunchConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLaunchConfigurations", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLaunchConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeLaunchConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLaunchConfigurations", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLaunchConfigurations), varargs...)
}

func (m *MockAutoscalingClient) DescribeLifecycleHooks(arg0 context.Context, arg1 *autoscaling.DescribeLifecycleHooksInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLifecycleHooksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLifecycleHooks", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLifecycleHooksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeLifecycleHooks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLifecycleHooks", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLifecycleHooks), varargs...)
}

func (m *MockAutoscalingClient) DescribeLoadBalancerTargetGroups(arg0 context.Context, arg1 *autoscaling.DescribeLoadBalancerTargetGroupsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoadBalancerTargetGroups", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLoadBalancerTargetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeLoadBalancerTargetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancerTargetGroups", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLoadBalancerTargetGroups), varargs...)
}

func (m *MockAutoscalingClient) DescribeLoadBalancers(arg0 context.Context, arg1 *autoscaling.DescribeLoadBalancersInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLoadBalancers", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeLoadBalancers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLoadBalancers", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeLoadBalancers), varargs...)
}

func (m *MockAutoscalingClient) DescribeNotificationConfigurations(arg0 context.Context, arg1 *autoscaling.DescribeNotificationConfigurationsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeNotificationConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeNotificationConfigurations", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeNotificationConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeNotificationConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeNotificationConfigurations", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeNotificationConfigurations), varargs...)
}

func (m *MockAutoscalingClient) DescribePolicies(arg0 context.Context, arg1 *autoscaling.DescribePoliciesInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribePolicies", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribePolicies", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribePolicies), varargs...)
}

func (m *MockAutoscalingClient) DescribeScheduledActions(arg0 context.Context, arg1 *autoscaling.DescribeScheduledActionsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeScheduledActionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeScheduledActions", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeScheduledActionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeScheduledActions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeScheduledActions", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeScheduledActions), varargs...)
}

func (m *MockAutoscalingClient) DescribeTags(arg0 context.Context, arg1 *autoscaling.DescribeTagsInput, arg2 ...func(*autoscaling.Options)) (*autoscaling.DescribeTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTags", varargs...)
	ret0, _ := ret[0].(*autoscaling.DescribeTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAutoscalingClientMockRecorder) DescribeTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTags", reflect.TypeOf((*MockAutoscalingClient)(nil).DescribeTags), varargs...)
}
