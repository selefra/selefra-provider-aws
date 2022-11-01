package mocks

import (
	context "context"
	reflect "reflect"

	qldb "github.com/aws/aws-sdk-go-v2/service/qldb"
	gomock "github.com/golang/mock/gomock"
)

type MockQLDBClient struct {
	ctrl		*gomock.Controller
	recorder	*MockQLDBClientMockRecorder
}

type MockQLDBClientMockRecorder struct {
	mock *MockQLDBClient
}

func NewMockQLDBClient(ctrl *gomock.Controller) *MockQLDBClient {
	mock := &MockQLDBClient{ctrl: ctrl}
	mock.recorder = &MockQLDBClientMockRecorder{mock}
	return mock
}

func (m *MockQLDBClient) EXPECT() *MockQLDBClientMockRecorder {
	return m.recorder
}

func (m *MockQLDBClient) DescribeLedger(arg0 context.Context, arg1 *qldb.DescribeLedgerInput, arg2 ...func(*qldb.Options)) (*qldb.DescribeLedgerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeLedger", varargs...)
	ret0, _ := ret[0].(*qldb.DescribeLedgerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQLDBClientMockRecorder) DescribeLedger(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeLedger", reflect.TypeOf((*MockQLDBClient)(nil).DescribeLedger), varargs...)
}

func (m *MockQLDBClient) ListJournalKinesisStreamsForLedger(arg0 context.Context, arg1 *qldb.ListJournalKinesisStreamsForLedgerInput, arg2 ...func(*qldb.Options)) (*qldb.ListJournalKinesisStreamsForLedgerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListJournalKinesisStreamsForLedger", varargs...)
	ret0, _ := ret[0].(*qldb.ListJournalKinesisStreamsForLedgerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQLDBClientMockRecorder) ListJournalKinesisStreamsForLedger(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJournalKinesisStreamsForLedger", reflect.TypeOf((*MockQLDBClient)(nil).ListJournalKinesisStreamsForLedger), varargs...)
}

func (m *MockQLDBClient) ListJournalS3ExportsForLedger(arg0 context.Context, arg1 *qldb.ListJournalS3ExportsForLedgerInput, arg2 ...func(*qldb.Options)) (*qldb.ListJournalS3ExportsForLedgerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListJournalS3ExportsForLedger", varargs...)
	ret0, _ := ret[0].(*qldb.ListJournalS3ExportsForLedgerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQLDBClientMockRecorder) ListJournalS3ExportsForLedger(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJournalS3ExportsForLedger", reflect.TypeOf((*MockQLDBClient)(nil).ListJournalS3ExportsForLedger), varargs...)
}

func (m *MockQLDBClient) ListLedgers(arg0 context.Context, arg1 *qldb.ListLedgersInput, arg2 ...func(*qldb.Options)) (*qldb.ListLedgersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListLedgers", varargs...)
	ret0, _ := ret[0].(*qldb.ListLedgersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQLDBClientMockRecorder) ListLedgers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLedgers", reflect.TypeOf((*MockQLDBClient)(nil).ListLedgers), varargs...)
}

func (m *MockQLDBClient) ListTagsForResource(arg0 context.Context, arg1 *qldb.ListTagsForResourceInput, arg2 ...func(*qldb.Options)) (*qldb.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*qldb.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockQLDBClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockQLDBClient)(nil).ListTagsForResource), varargs...)
}
