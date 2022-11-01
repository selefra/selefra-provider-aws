package mocks

import (
	context "context"
	reflect "reflect"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	gomock "github.com/golang/mock/gomock"
)

type MockDynamoDBClient struct {
	ctrl		*gomock.Controller
	recorder	*MockDynamoDBClientMockRecorder
}

type MockDynamoDBClientMockRecorder struct {
	mock *MockDynamoDBClient
}

func NewMockDynamoDBClient(ctrl *gomock.Controller) *MockDynamoDBClient {
	mock := &MockDynamoDBClient{ctrl: ctrl}
	mock.recorder = &MockDynamoDBClientMockRecorder{mock}
	return mock
}

func (m *MockDynamoDBClient) EXPECT() *MockDynamoDBClientMockRecorder {
	return m.recorder
}

func (m *MockDynamoDBClient) DescribeContinuousBackups(arg0 context.Context, arg1 *dynamodb.DescribeContinuousBackupsInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeContinuousBackups", varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeContinuousBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamoDBClientMockRecorder) DescribeContinuousBackups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeContinuousBackups", reflect.TypeOf((*MockDynamoDBClient)(nil).DescribeContinuousBackups), varargs...)
}

func (m *MockDynamoDBClient) DescribeTable(arg0 context.Context, arg1 *dynamodb.DescribeTableInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTable", varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamoDBClientMockRecorder) DescribeTable(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTable", reflect.TypeOf((*MockDynamoDBClient)(nil).DescribeTable), varargs...)
}

func (m *MockDynamoDBClient) DescribeTableReplicaAutoScaling(arg0 context.Context, arg1 *dynamodb.DescribeTableReplicaAutoScalingInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeTableReplicaAutoScaling", varargs...)
	ret0, _ := ret[0].(*dynamodb.DescribeTableReplicaAutoScalingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamoDBClientMockRecorder) DescribeTableReplicaAutoScaling(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTableReplicaAutoScaling", reflect.TypeOf((*MockDynamoDBClient)(nil).DescribeTableReplicaAutoScaling), varargs...)
}

func (m *MockDynamoDBClient) ListTables(arg0 context.Context, arg1 *dynamodb.ListTablesInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.ListTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTables", varargs...)
	ret0, _ := ret[0].(*dynamodb.ListTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamoDBClientMockRecorder) ListTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTables", reflect.TypeOf((*MockDynamoDBClient)(nil).ListTables), varargs...)
}

func (m *MockDynamoDBClient) ListTagsOfResource(arg0 context.Context, arg1 *dynamodb.ListTagsOfResourceInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.ListTagsOfResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsOfResource", varargs...)
	ret0, _ := ret[0].(*dynamodb.ListTagsOfResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDynamoDBClientMockRecorder) ListTagsOfResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsOfResource", reflect.TypeOf((*MockDynamoDBClient)(nil).ListTagsOfResource), varargs...)
}
