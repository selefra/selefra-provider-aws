package mocks

import (
	context "context"
	reflect "reflect"

	sns "github.com/aws/aws-sdk-go-v2/service/sns"
	gomock "github.com/golang/mock/gomock"
)

type MockSnsClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSnsClientMockRecorder
}

type MockSnsClientMockRecorder struct {
	mock *MockSnsClient
}

func NewMockSnsClient(ctrl *gomock.Controller) *MockSnsClient {
	mock := &MockSnsClient{ctrl: ctrl}
	mock.recorder = &MockSnsClientMockRecorder{mock}
	return mock
}

func (m *MockSnsClient) EXPECT() *MockSnsClientMockRecorder {
	return m.recorder
}

func (m *MockSnsClient) GetSubscriptionAttributes(arg0 context.Context, arg1 *sns.GetSubscriptionAttributesInput, arg2 ...func(*sns.Options)) (*sns.GetSubscriptionAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSubscriptionAttributes", varargs...)
	ret0, _ := ret[0].(*sns.GetSubscriptionAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) GetSubscriptionAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionAttributes", reflect.TypeOf((*MockSnsClient)(nil).GetSubscriptionAttributes), varargs...)
}

func (m *MockSnsClient) GetTopicAttributes(arg0 context.Context, arg1 *sns.GetTopicAttributesInput, arg2 ...func(*sns.Options)) (*sns.GetTopicAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTopicAttributes", varargs...)
	ret0, _ := ret[0].(*sns.GetTopicAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) GetTopicAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopicAttributes", reflect.TypeOf((*MockSnsClient)(nil).GetTopicAttributes), varargs...)
}

func (m *MockSnsClient) ListSubscriptions(arg0 context.Context, arg1 *sns.ListSubscriptionsInput, arg2 ...func(*sns.Options)) (*sns.ListSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSubscriptions", varargs...)
	ret0, _ := ret[0].(*sns.ListSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSubscriptions", reflect.TypeOf((*MockSnsClient)(nil).ListSubscriptions), varargs...)
}

func (m *MockSnsClient) ListTagsForResource(arg0 context.Context, arg1 *sns.ListTagsForResourceInput, arg2 ...func(*sns.Options)) (*sns.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*sns.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockSnsClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockSnsClient) ListTopics(arg0 context.Context, arg1 *sns.ListTopicsInput, arg2 ...func(*sns.Options)) (*sns.ListTopicsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTopics", varargs...)
	ret0, _ := ret[0].(*sns.ListTopicsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSnsClientMockRecorder) ListTopics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTopics", reflect.TypeOf((*MockSnsClient)(nil).ListTopics), varargs...)
}
