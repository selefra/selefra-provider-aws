package mocks

import (
	context "context"
	reflect "reflect"

	glue "github.com/aws/aws-sdk-go-v2/service/glue"
	gomock "github.com/golang/mock/gomock"
)

type MockGlueClient struct {
	ctrl		*gomock.Controller
	recorder	*MockGlueClientMockRecorder
}

type MockGlueClientMockRecorder struct {
	mock *MockGlueClient
}

func NewMockGlueClient(ctrl *gomock.Controller) *MockGlueClient {
	mock := &MockGlueClient{ctrl: ctrl}
	mock.recorder = &MockGlueClientMockRecorder{mock}
	return mock
}

func (m *MockGlueClient) EXPECT() *MockGlueClientMockRecorder {
	return m.recorder
}

func (m *MockGlueClient) GetClassifiers(arg0 context.Context, arg1 *glue.GetClassifiersInput, arg2 ...func(*glue.Options)) (*glue.GetClassifiersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetClassifiers", varargs...)
	ret0, _ := ret[0].(*glue.GetClassifiersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetClassifiers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassifiers", reflect.TypeOf((*MockGlueClient)(nil).GetClassifiers), varargs...)
}

func (m *MockGlueClient) GetConnections(arg0 context.Context, arg1 *glue.GetConnectionsInput, arg2 ...func(*glue.Options)) (*glue.GetConnectionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConnections", varargs...)
	ret0, _ := ret[0].(*glue.GetConnectionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetConnections(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnections", reflect.TypeOf((*MockGlueClient)(nil).GetConnections), varargs...)
}

func (m *MockGlueClient) GetCrawlers(arg0 context.Context, arg1 *glue.GetCrawlersInput, arg2 ...func(*glue.Options)) (*glue.GetCrawlersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCrawlers", varargs...)
	ret0, _ := ret[0].(*glue.GetCrawlersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetCrawlers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCrawlers", reflect.TypeOf((*MockGlueClient)(nil).GetCrawlers), varargs...)
}

func (m *MockGlueClient) GetDataCatalogEncryptionSettings(arg0 context.Context, arg1 *glue.GetDataCatalogEncryptionSettingsInput, arg2 ...func(*glue.Options)) (*glue.GetDataCatalogEncryptionSettingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDataCatalogEncryptionSettings", varargs...)
	ret0, _ := ret[0].(*glue.GetDataCatalogEncryptionSettingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDataCatalogEncryptionSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataCatalogEncryptionSettings", reflect.TypeOf((*MockGlueClient)(nil).GetDataCatalogEncryptionSettings), varargs...)
}

func (m *MockGlueClient) GetDatabases(arg0 context.Context, arg1 *glue.GetDatabasesInput, arg2 ...func(*glue.Options)) (*glue.GetDatabasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDatabases", varargs...)
	ret0, _ := ret[0].(*glue.GetDatabasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDatabases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabases", reflect.TypeOf((*MockGlueClient)(nil).GetDatabases), varargs...)
}

func (m *MockGlueClient) GetDevEndpoints(arg0 context.Context, arg1 *glue.GetDevEndpointsInput, arg2 ...func(*glue.Options)) (*glue.GetDevEndpointsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDevEndpoints", varargs...)
	ret0, _ := ret[0].(*glue.GetDevEndpointsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetDevEndpoints(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevEndpoints", reflect.TypeOf((*MockGlueClient)(nil).GetDevEndpoints), varargs...)
}

func (m *MockGlueClient) GetJobRuns(arg0 context.Context, arg1 *glue.GetJobRunsInput, arg2 ...func(*glue.Options)) (*glue.GetJobRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetJobRuns", varargs...)
	ret0, _ := ret[0].(*glue.GetJobRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetJobRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobRuns", reflect.TypeOf((*MockGlueClient)(nil).GetJobRuns), varargs...)
}

func (m *MockGlueClient) GetJobs(arg0 context.Context, arg1 *glue.GetJobsInput, arg2 ...func(*glue.Options)) (*glue.GetJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetJobs", varargs...)
	ret0, _ := ret[0].(*glue.GetJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobs", reflect.TypeOf((*MockGlueClient)(nil).GetJobs), varargs...)
}

func (m *MockGlueClient) GetMLTaskRuns(arg0 context.Context, arg1 *glue.GetMLTaskRunsInput, arg2 ...func(*glue.Options)) (*glue.GetMLTaskRunsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMLTaskRuns", varargs...)
	ret0, _ := ret[0].(*glue.GetMLTaskRunsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetMLTaskRuns(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMLTaskRuns", reflect.TypeOf((*MockGlueClient)(nil).GetMLTaskRuns), varargs...)
}

func (m *MockGlueClient) GetMLTransforms(arg0 context.Context, arg1 *glue.GetMLTransformsInput, arg2 ...func(*glue.Options)) (*glue.GetMLTransformsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMLTransforms", varargs...)
	ret0, _ := ret[0].(*glue.GetMLTransformsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetMLTransforms(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMLTransforms", reflect.TypeOf((*MockGlueClient)(nil).GetMLTransforms), varargs...)
}

func (m *MockGlueClient) GetPartitionIndexes(arg0 context.Context, arg1 *glue.GetPartitionIndexesInput, arg2 ...func(*glue.Options)) (*glue.GetPartitionIndexesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPartitionIndexes", varargs...)
	ret0, _ := ret[0].(*glue.GetPartitionIndexesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetPartitionIndexes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitionIndexes", reflect.TypeOf((*MockGlueClient)(nil).GetPartitionIndexes), varargs...)
}

func (m *MockGlueClient) GetSchema(arg0 context.Context, arg1 *glue.GetSchemaInput, arg2 ...func(*glue.Options)) (*glue.GetSchemaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSchema", varargs...)
	ret0, _ := ret[0].(*glue.GetSchemaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetSchema(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchema", reflect.TypeOf((*MockGlueClient)(nil).GetSchema), varargs...)
}

func (m *MockGlueClient) GetSchemaVersion(arg0 context.Context, arg1 *glue.GetSchemaVersionInput, arg2 ...func(*glue.Options)) (*glue.GetSchemaVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSchemaVersion", varargs...)
	ret0, _ := ret[0].(*glue.GetSchemaVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetSchemaVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchemaVersion", reflect.TypeOf((*MockGlueClient)(nil).GetSchemaVersion), varargs...)
}

func (m *MockGlueClient) GetSecurityConfigurations(arg0 context.Context, arg1 *glue.GetSecurityConfigurationsInput, arg2 ...func(*glue.Options)) (*glue.GetSecurityConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSecurityConfigurations", varargs...)
	ret0, _ := ret[0].(*glue.GetSecurityConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetSecurityConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecurityConfigurations", reflect.TypeOf((*MockGlueClient)(nil).GetSecurityConfigurations), varargs...)
}

func (m *MockGlueClient) GetTables(arg0 context.Context, arg1 *glue.GetTablesInput, arg2 ...func(*glue.Options)) (*glue.GetTablesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTables", varargs...)
	ret0, _ := ret[0].(*glue.GetTablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetTables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTables", reflect.TypeOf((*MockGlueClient)(nil).GetTables), varargs...)
}

func (m *MockGlueClient) GetTags(arg0 context.Context, arg1 *glue.GetTagsInput, arg2 ...func(*glue.Options)) (*glue.GetTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTags", varargs...)
	ret0, _ := ret[0].(*glue.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockGlueClient)(nil).GetTags), varargs...)
}

func (m *MockGlueClient) GetTrigger(arg0 context.Context, arg1 *glue.GetTriggerInput, arg2 ...func(*glue.Options)) (*glue.GetTriggerOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTrigger", varargs...)
	ret0, _ := ret[0].(*glue.GetTriggerOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetTrigger(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrigger", reflect.TypeOf((*MockGlueClient)(nil).GetTrigger), varargs...)
}

func (m *MockGlueClient) GetWorkflow(arg0 context.Context, arg1 *glue.GetWorkflowInput, arg2 ...func(*glue.Options)) (*glue.GetWorkflowOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkflow", varargs...)
	ret0, _ := ret[0].(*glue.GetWorkflowOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) GetWorkflow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkflow", reflect.TypeOf((*MockGlueClient)(nil).GetWorkflow), varargs...)
}

func (m *MockGlueClient) ListRegistries(arg0 context.Context, arg1 *glue.ListRegistriesInput, arg2 ...func(*glue.Options)) (*glue.ListRegistriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRegistries", varargs...)
	ret0, _ := ret[0].(*glue.ListRegistriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListRegistries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRegistries", reflect.TypeOf((*MockGlueClient)(nil).ListRegistries), varargs...)
}

func (m *MockGlueClient) ListSchemaVersions(arg0 context.Context, arg1 *glue.ListSchemaVersionsInput, arg2 ...func(*glue.Options)) (*glue.ListSchemaVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSchemaVersions", varargs...)
	ret0, _ := ret[0].(*glue.ListSchemaVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListSchemaVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchemaVersions", reflect.TypeOf((*MockGlueClient)(nil).ListSchemaVersions), varargs...)
}

func (m *MockGlueClient) ListSchemas(arg0 context.Context, arg1 *glue.ListSchemasInput, arg2 ...func(*glue.Options)) (*glue.ListSchemasOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSchemas", varargs...)
	ret0, _ := ret[0].(*glue.ListSchemasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListSchemas(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSchemas", reflect.TypeOf((*MockGlueClient)(nil).ListSchemas), varargs...)
}

func (m *MockGlueClient) ListTriggers(arg0 context.Context, arg1 *glue.ListTriggersInput, arg2 ...func(*glue.Options)) (*glue.ListTriggersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTriggers", varargs...)
	ret0, _ := ret[0].(*glue.ListTriggersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListTriggers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTriggers", reflect.TypeOf((*MockGlueClient)(nil).ListTriggers), varargs...)
}

func (m *MockGlueClient) ListWorkflows(arg0 context.Context, arg1 *glue.ListWorkflowsInput, arg2 ...func(*glue.Options)) (*glue.ListWorkflowsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWorkflows", varargs...)
	ret0, _ := ret[0].(*glue.ListWorkflowsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) ListWorkflows(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkflows", reflect.TypeOf((*MockGlueClient)(nil).ListWorkflows), varargs...)
}

func (m *MockGlueClient) QuerySchemaVersionMetadata(arg0 context.Context, arg1 *glue.QuerySchemaVersionMetadataInput, arg2 ...func(*glue.Options)) (*glue.QuerySchemaVersionMetadataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QuerySchemaVersionMetadata", varargs...)
	ret0, _ := ret[0].(*glue.QuerySchemaVersionMetadataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGlueClientMockRecorder) QuerySchemaVersionMetadata(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerySchemaVersionMetadata", reflect.TypeOf((*MockGlueClient)(nil).QuerySchemaVersionMetadata), varargs...)
}
