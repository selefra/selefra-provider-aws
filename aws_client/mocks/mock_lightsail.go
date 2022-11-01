package mocks

import (
	context "context"
	reflect "reflect"

	lightsail "github.com/aws/aws-sdk-go-v2/service/lightsail"
	gomock "github.com/golang/mock/gomock"
)

type MockLightsailClient struct {
	ctrl		*gomock.Controller
	recorder	*MockLightsailClientMockRecorder
}

type MockLightsailClientMockRecorder struct {
	mock *MockLightsailClient
}

func NewMockLightsailClient(ctrl *gomock.Controller) *MockLightsailClient {
	mock := &MockLightsailClient{ctrl: ctrl}
	mock.recorder = &MockLightsailClientMockRecorder{mock}
	return mock
}

func (m *MockLightsailClient) EXPECT() *MockLightsailClientMockRecorder {
	return m.recorder
}

func (m *MockLightsailClient) GetAlarms(arg0 context.Context, arg1 *lightsail.GetAlarmsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetAlarmsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAlarms", varargs...)
	ret0, _ := ret[0].(*lightsail.GetAlarmsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetAlarms(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlarms", reflect.TypeOf((*MockLightsailClient)(nil).GetAlarms), varargs...)
}

func (m *MockLightsailClient) GetBucketAccessKeys(arg0 context.Context, arg1 *lightsail.GetBucketAccessKeysInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetBucketAccessKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketAccessKeys", varargs...)
	ret0, _ := ret[0].(*lightsail.GetBucketAccessKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetBucketAccessKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketAccessKeys", reflect.TypeOf((*MockLightsailClient)(nil).GetBucketAccessKeys), varargs...)
}

func (m *MockLightsailClient) GetBuckets(arg0 context.Context, arg1 *lightsail.GetBucketsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetBucketsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBuckets", varargs...)
	ret0, _ := ret[0].(*lightsail.GetBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetBuckets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuckets", reflect.TypeOf((*MockLightsailClient)(nil).GetBuckets), varargs...)
}

func (m *MockLightsailClient) GetCertificates(arg0 context.Context, arg1 *lightsail.GetCertificatesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCertificates", varargs...)
	ret0, _ := ret[0].(*lightsail.GetCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificates", reflect.TypeOf((*MockLightsailClient)(nil).GetCertificates), varargs...)
}

func (m *MockLightsailClient) GetContainerImages(arg0 context.Context, arg1 *lightsail.GetContainerImagesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetContainerImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContainerImages", varargs...)
	ret0, _ := ret[0].(*lightsail.GetContainerImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetContainerImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerImages", reflect.TypeOf((*MockLightsailClient)(nil).GetContainerImages), varargs...)
}

func (m *MockLightsailClient) GetContainerServiceDeployments(arg0 context.Context, arg1 *lightsail.GetContainerServiceDeploymentsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetContainerServiceDeploymentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContainerServiceDeployments", varargs...)
	ret0, _ := ret[0].(*lightsail.GetContainerServiceDeploymentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetContainerServiceDeployments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerServiceDeployments", reflect.TypeOf((*MockLightsailClient)(nil).GetContainerServiceDeployments), varargs...)
}

func (m *MockLightsailClient) GetContainerServices(arg0 context.Context, arg1 *lightsail.GetContainerServicesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetContainerServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContainerServices", varargs...)
	ret0, _ := ret[0].(*lightsail.GetContainerServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetContainerServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerServices", reflect.TypeOf((*MockLightsailClient)(nil).GetContainerServices), varargs...)
}

func (m *MockLightsailClient) GetDiskSnapshots(arg0 context.Context, arg1 *lightsail.GetDiskSnapshotsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDiskSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDiskSnapshots", varargs...)
	ret0, _ := ret[0].(*lightsail.GetDiskSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDiskSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskSnapshots", reflect.TypeOf((*MockLightsailClient)(nil).GetDiskSnapshots), varargs...)
}

func (m *MockLightsailClient) GetDisks(arg0 context.Context, arg1 *lightsail.GetDisksInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDisksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDisks", varargs...)
	ret0, _ := ret[0].(*lightsail.GetDisksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDisks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDisks", reflect.TypeOf((*MockLightsailClient)(nil).GetDisks), varargs...)
}

func (m *MockLightsailClient) GetDistributionLatestCacheReset(arg0 context.Context, arg1 *lightsail.GetDistributionLatestCacheResetInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDistributionLatestCacheResetOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDistributionLatestCacheReset", varargs...)
	ret0, _ := ret[0].(*lightsail.GetDistributionLatestCacheResetOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDistributionLatestCacheReset(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDistributionLatestCacheReset", reflect.TypeOf((*MockLightsailClient)(nil).GetDistributionLatestCacheReset), varargs...)
}

func (m *MockLightsailClient) GetDistributions(arg0 context.Context, arg1 *lightsail.GetDistributionsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetDistributionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDistributions", varargs...)
	ret0, _ := ret[0].(*lightsail.GetDistributionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetDistributions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDistributions", reflect.TypeOf((*MockLightsailClient)(nil).GetDistributions), varargs...)
}

func (m *MockLightsailClient) GetInstanceAccessDetails(arg0 context.Context, arg1 *lightsail.GetInstanceAccessDetailsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstanceAccessDetailsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInstanceAccessDetails", varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstanceAccessDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstanceAccessDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceAccessDetails", reflect.TypeOf((*MockLightsailClient)(nil).GetInstanceAccessDetails), varargs...)
}

func (m *MockLightsailClient) GetInstancePortStates(arg0 context.Context, arg1 *lightsail.GetInstancePortStatesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstancePortStatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInstancePortStates", varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstancePortStatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstancePortStates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstancePortStates", reflect.TypeOf((*MockLightsailClient)(nil).GetInstancePortStates), varargs...)
}

func (m *MockLightsailClient) GetInstanceSnapshots(arg0 context.Context, arg1 *lightsail.GetInstanceSnapshotsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstanceSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInstanceSnapshots", varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstanceSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstanceSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceSnapshots", reflect.TypeOf((*MockLightsailClient)(nil).GetInstanceSnapshots), varargs...)
}

func (m *MockLightsailClient) GetInstances(arg0 context.Context, arg1 *lightsail.GetInstancesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInstances", varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstances", reflect.TypeOf((*MockLightsailClient)(nil).GetInstances), varargs...)
}

func (m *MockLightsailClient) GetLoadBalancerTlsCertificates(arg0 context.Context, arg1 *lightsail.GetLoadBalancerTlsCertificatesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetLoadBalancerTlsCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLoadBalancerTlsCertificates", varargs...)
	ret0, _ := ret[0].(*lightsail.GetLoadBalancerTlsCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetLoadBalancerTlsCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoadBalancerTlsCertificates", reflect.TypeOf((*MockLightsailClient)(nil).GetLoadBalancerTlsCertificates), varargs...)
}

func (m *MockLightsailClient) GetLoadBalancers(arg0 context.Context, arg1 *lightsail.GetLoadBalancersInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetLoadBalancersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLoadBalancers", varargs...)
	ret0, _ := ret[0].(*lightsail.GetLoadBalancersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetLoadBalancers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoadBalancers", reflect.TypeOf((*MockLightsailClient)(nil).GetLoadBalancers), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseEvents(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseEventsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabaseEvents", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabaseEvents", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseEvents), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseLogEvents(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseLogEventsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabaseLogEvents", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseLogEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabaseLogEvents", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseLogEvents), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseLogStreams(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseLogStreamsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabaseLogStreams", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseLogStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseLogStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabaseLogStreams", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseLogStreams), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseParameters(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseParametersInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabaseParameters", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabaseParameters", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseParameters), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabaseSnapshots(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseSnapshotsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabaseSnapshots", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabaseSnapshots", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseSnapshots), varargs...)
}

func (m *MockLightsailClient) GetRelationalDatabases(arg0 context.Context, arg1 *lightsail.GetRelationalDatabasesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabases", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetRelationalDatabases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabases", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabases), varargs...)
}

func (m *MockLightsailClient) GetStaticIps(arg0 context.Context, arg1 *lightsail.GetStaticIpsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetStaticIpsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStaticIps", varargs...)
	ret0, _ := ret[0].(*lightsail.GetStaticIpsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLightsailClientMockRecorder) GetStaticIps(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStaticIps", reflect.TypeOf((*MockLightsailClient)(nil).GetStaticIps), varargs...)
}
