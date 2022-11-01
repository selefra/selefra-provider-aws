package mocks

import (
	context "context"
	reflect "reflect"

	athena "github.com/aws/aws-sdk-go-v2/service/athena"
	gomock "github.com/golang/mock/gomock"
)

type MockAthenaClient struct {
	ctrl		*gomock.Controller
	recorder	*MockAthenaClientMockRecorder
}

type MockAthenaClientMockRecorder struct {
	mock *MockAthenaClient
}

func NewMockAthenaClient(ctrl *gomock.Controller) *MockAthenaClient {
	mock := &MockAthenaClient{ctrl: ctrl}
	mock.recorder = &MockAthenaClientMockRecorder{mock}
	return mock
}

func (m *MockAthenaClient) EXPECT() *MockAthenaClientMockRecorder {
	return m.recorder
}

func (m *MockAthenaClient) GetDataCatalog(arg0 context.Context, arg1 *athena.GetDataCatalogInput, arg2 ...func(*athena.Options)) (*athena.GetDataCatalogOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDataCatalog", varargs...)
	ret0, _ := ret[0].(*athena.GetDataCatalogOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetDataCatalog(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataCatalog", reflect.TypeOf((*MockAthenaClient)(nil).GetDataCatalog), varargs...)
}

func (m *MockAthenaClient) GetNamedQuery(arg0 context.Context, arg1 *athena.GetNamedQueryInput, arg2 ...func(*athena.Options)) (*athena.GetNamedQueryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNamedQuery", varargs...)
	ret0, _ := ret[0].(*athena.GetNamedQueryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetNamedQuery(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamedQuery", reflect.TypeOf((*MockAthenaClient)(nil).GetNamedQuery), varargs...)
}

func (m *MockAthenaClient) GetPreparedStatement(arg0 context.Context, arg1 *athena.GetPreparedStatementInput, arg2 ...func(*athena.Options)) (*athena.GetPreparedStatementOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPreparedStatement", varargs...)
	ret0, _ := ret[0].(*athena.GetPreparedStatementOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetPreparedStatement(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreparedStatement", reflect.TypeOf((*MockAthenaClient)(nil).GetPreparedStatement), varargs...)
}

func (m *MockAthenaClient) GetQueryExecution(arg0 context.Context, arg1 *athena.GetQueryExecutionInput, arg2 ...func(*athena.Options)) (*athena.GetQueryExecutionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetQueryExecution", varargs...)
	ret0, _ := ret[0].(*athena.GetQueryExecutionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetQueryExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryExecution", reflect.TypeOf((*MockAthenaClient)(nil).GetQueryExecution), varargs...)
}

func (m *MockAthenaClient) GetWorkGroup(arg0 context.Context, arg1 *athena.GetWorkGroupInput, arg2 ...func(*athena.Options)) (*athena.GetWorkGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkGroup", varargs...)
	ret0, _ := ret[0].(*athena.GetWorkGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) GetWorkGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkGroup", reflect.TypeOf((*MockAthenaClient)(nil).GetWorkGroup), varargs...)
}

func (m *MockAthenaClient) ListDataCatalogs(arg0 context.Context, arg1 *athena.ListDataCatalogsInput, arg2 ...func(*athena.Options)) (*athena.ListDataCatalogsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDataCatalogs", varargs...)
	ret0, _ := ret[0].(*athena.ListDataCatalogsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListDataCatalogs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDataCatalogs", reflect.TypeOf((*MockAthenaClient)(nil).ListDataCatalogs), varargs...)
}

func (m *MockAthenaClient) ListDatabases(arg0 context.Context, arg1 *athena.ListDatabasesInput, arg2 ...func(*athena.Options)) (*athena.ListDatabasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDatabases", varargs...)
	ret0, _ := ret[0].(*athena.ListDatabasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListDatabases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDatabases", reflect.TypeOf((*MockAthenaClient)(nil).ListDatabases), varargs...)
}

func (m *MockAthenaClient) ListNamedQueries(arg0 context.Context, arg1 *athena.ListNamedQueriesInput, arg2 ...func(*athena.Options)) (*athena.ListNamedQueriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNamedQueries", varargs...)
	ret0, _ := ret[0].(*athena.ListNamedQueriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListNamedQueries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamedQueries", reflect.TypeOf((*MockAthenaClient)(nil).ListNamedQueries), varargs...)
}

func (m *MockAthenaClient) ListPreparedStatements(arg0 context.Context, arg1 *athena.ListPreparedStatementsInput, arg2 ...func(*athena.Options)) (*athena.ListPreparedStatementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPreparedStatements", varargs...)
	ret0, _ := ret[0].(*athena.ListPreparedStatementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListPreparedStatements(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPreparedStatements", reflect.TypeOf((*MockAthenaClient)(nil).ListPreparedStatements), varargs...)
}

func (m *MockAthenaClient) ListQueryExecutions(arg0 context.Context, arg1 *athena.ListQueryExecutionsInput, arg2 ...func(*athena.Options)) (*athena.ListQueryExecutionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListQueryExecutions", varargs...)
	ret0, _ := ret[0].(*athena.ListQueryExecutionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListQueryExecutions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListQueryExecutions", reflect.TypeOf((*MockAthenaClient)(nil).ListQueryExecutions), varargs...)
}

func (m *MockAthenaClient) ListTableMetadata(arg0 context.Context, arg1 *athena.ListTableMetadataInput, arg2 ...func(*athena.Options)) (*athena.ListTableMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTableMetadata", varargs...)
	ret0, _ := ret[0].(*athena.ListTableMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListTableMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTableMetadata", reflect.TypeOf((*MockAthenaClient)(nil).ListTableMetadata), varargs...)
}

func (m *MockAthenaClient) ListTagsForResource(arg0 context.Context, arg1 *athena.ListTagsForResourceInput, arg2 ...func(*athena.Options)) (*athena.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*athena.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAthenaClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockAthenaClient) ListWorkGroups(arg0 context.Context, arg1 *athena.ListWorkGroupsInput, arg2 ...func(*athena.Options)) (*athena.ListWorkGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWorkGroups", varargs...)
	ret0, _ := ret[0].(*athena.ListWorkGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAthenaClientMockRecorder) ListWorkGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkGroups", reflect.TypeOf((*MockAthenaClient)(nil).ListWorkGroups), varargs...)
}
