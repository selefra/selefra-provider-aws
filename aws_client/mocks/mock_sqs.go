package mocks

import (
	context "context"
	reflect "reflect"

	sqs "github.com/aws/aws-sdk-go-v2/service/sqs"
	gomock "github.com/golang/mock/gomock"
)

type MockSQSClient struct {
	ctrl		*gomock.Controller
	recorder	*MockSQSClientMockRecorder
}

type MockSQSClientMockRecorder struct {
	mock *MockSQSClient
}

func NewMockSQSClient(ctrl *gomock.Controller) *MockSQSClient {
	mock := &MockSQSClient{ctrl: ctrl}
	mock.recorder = &MockSQSClientMockRecorder{mock}
	return mock
}

func (m *MockSQSClient) EXPECT() *MockSQSClientMockRecorder {
	return m.recorder
}

func (m *MockSQSClient) GetQueueAttributes(arg0 context.Context, arg1 *sqs.GetQueueAttributesInput, arg2 ...func(*sqs.Options)) (*sqs.GetQueueAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueueAttributes", varargs...)
	ret0, _ := ret[0].(*sqs.GetQueueAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQSClientMockRecorder) GetQueueAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueueAttributes", reflect.TypeOf((*MockSQSClient)(nil).GetQueueAttributes), varargs...)
}

func (m *MockSQSClient) ListQueueTags(arg0 context.Context, arg1 *sqs.ListQueueTagsInput, arg2 ...func(*sqs.Options)) (*sqs.ListQueueTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueueTags", varargs...)
	ret0, _ := ret[0].(*sqs.ListQueueTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQSClientMockRecorder) ListQueueTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueueTags", reflect.TypeOf((*MockSQSClient)(nil).ListQueueTags), varargs...)
}

func (m *MockSQSClient) ListQueues(arg0 context.Context, arg1 *sqs.ListQueuesInput, arg2 ...func(*sqs.Options)) (*sqs.ListQueuesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueues", varargs...)
	ret0, _ := ret[0].(*sqs.ListQueuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQSClientMockRecorder) ListQueues(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueues", reflect.TypeOf((*MockSQSClient)(nil).ListQueues), varargs...)
}
