package mocks

import (
	context "context"
	reflect "reflect"

	iot "github.com/aws/aws-sdk-go-v2/service/iot"
	gomock "github.com/golang/mock/gomock"
)

type MockIOTClient struct {
	ctrl		*gomock.Controller
	recorder	*MockIOTClientMockRecorder
}

type MockIOTClientMockRecorder struct {
	mock *MockIOTClient
}

func NewMockIOTClient(ctrl *gomock.Controller) *MockIOTClient {
	mock := &MockIOTClient{ctrl: ctrl}
	mock.recorder = &MockIOTClientMockRecorder{mock}
	return mock
}

func (m *MockIOTClient) EXPECT() *MockIOTClientMockRecorder {
	return m.recorder
}

func (m *MockIOTClient) DescribeBillingGroup(arg0 context.Context, arg1 *iot.DescribeBillingGroupInput, arg2 ...func(*iot.Options)) (*iot.DescribeBillingGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBillingGroup", varargs...)
	ret0, _ := ret[0].(*iot.DescribeBillingGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) DescribeBillingGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBillingGroup", reflect.TypeOf((*MockIOTClient)(nil).DescribeBillingGroup), varargs...)
}

func (m *MockIOTClient) DescribeCACertificate(arg0 context.Context, arg1 *iot.DescribeCACertificateInput, arg2 ...func(*iot.Options)) (*iot.DescribeCACertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCACertificate", varargs...)
	ret0, _ := ret[0].(*iot.DescribeCACertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) DescribeCACertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCACertificate", reflect.TypeOf((*MockIOTClient)(nil).DescribeCACertificate), varargs...)
}

func (m *MockIOTClient) DescribeCertificate(arg0 context.Context, arg1 *iot.DescribeCertificateInput, arg2 ...func(*iot.Options)) (*iot.DescribeCertificateOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCertificate", varargs...)
	ret0, _ := ret[0].(*iot.DescribeCertificateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) DescribeCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCertificate", reflect.TypeOf((*MockIOTClient)(nil).DescribeCertificate), varargs...)
}

func (m *MockIOTClient) DescribeJob(arg0 context.Context, arg1 *iot.DescribeJobInput, arg2 ...func(*iot.Options)) (*iot.DescribeJobOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeJob", varargs...)
	ret0, _ := ret[0].(*iot.DescribeJobOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) DescribeJob(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeJob", reflect.TypeOf((*MockIOTClient)(nil).DescribeJob), varargs...)
}

func (m *MockIOTClient) DescribeSecurityProfile(arg0 context.Context, arg1 *iot.DescribeSecurityProfileInput, arg2 ...func(*iot.Options)) (*iot.DescribeSecurityProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSecurityProfile", varargs...)
	ret0, _ := ret[0].(*iot.DescribeSecurityProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) DescribeSecurityProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSecurityProfile", reflect.TypeOf((*MockIOTClient)(nil).DescribeSecurityProfile), varargs...)
}

func (m *MockIOTClient) DescribeStream(arg0 context.Context, arg1 *iot.DescribeStreamInput, arg2 ...func(*iot.Options)) (*iot.DescribeStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStream", varargs...)
	ret0, _ := ret[0].(*iot.DescribeStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) DescribeStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStream", reflect.TypeOf((*MockIOTClient)(nil).DescribeStream), varargs...)
}

func (m *MockIOTClient) DescribeThingGroup(arg0 context.Context, arg1 *iot.DescribeThingGroupInput, arg2 ...func(*iot.Options)) (*iot.DescribeThingGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeThingGroup", varargs...)
	ret0, _ := ret[0].(*iot.DescribeThingGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) DescribeThingGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeThingGroup", reflect.TypeOf((*MockIOTClient)(nil).DescribeThingGroup), varargs...)
}

func (m *MockIOTClient) GetPolicy(arg0 context.Context, arg1 *iot.GetPolicyInput, arg2 ...func(*iot.Options)) (*iot.GetPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPolicy", varargs...)
	ret0, _ := ret[0].(*iot.GetPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) GetPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicy", reflect.TypeOf((*MockIOTClient)(nil).GetPolicy), varargs...)
}

func (m *MockIOTClient) GetTopicRule(arg0 context.Context, arg1 *iot.GetTopicRuleInput, arg2 ...func(*iot.Options)) (*iot.GetTopicRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTopicRule", varargs...)
	ret0, _ := ret[0].(*iot.GetTopicRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) GetTopicRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopicRule", reflect.TypeOf((*MockIOTClient)(nil).GetTopicRule), varargs...)
}

func (m *MockIOTClient) ListAttachedPolicies(arg0 context.Context, arg1 *iot.ListAttachedPoliciesInput, arg2 ...func(*iot.Options)) (*iot.ListAttachedPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAttachedPolicies", varargs...)
	ret0, _ := ret[0].(*iot.ListAttachedPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListAttachedPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedPolicies", reflect.TypeOf((*MockIOTClient)(nil).ListAttachedPolicies), varargs...)
}

func (m *MockIOTClient) ListBillingGroups(arg0 context.Context, arg1 *iot.ListBillingGroupsInput, arg2 ...func(*iot.Options)) (*iot.ListBillingGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBillingGroups", varargs...)
	ret0, _ := ret[0].(*iot.ListBillingGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListBillingGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBillingGroups", reflect.TypeOf((*MockIOTClient)(nil).ListBillingGroups), varargs...)
}

func (m *MockIOTClient) ListCACertificates(arg0 context.Context, arg1 *iot.ListCACertificatesInput, arg2 ...func(*iot.Options)) (*iot.ListCACertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCACertificates", varargs...)
	ret0, _ := ret[0].(*iot.ListCACertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListCACertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCACertificates", reflect.TypeOf((*MockIOTClient)(nil).ListCACertificates), varargs...)
}

func (m *MockIOTClient) ListCertificates(arg0 context.Context, arg1 *iot.ListCertificatesInput, arg2 ...func(*iot.Options)) (*iot.ListCertificatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCertificates", varargs...)
	ret0, _ := ret[0].(*iot.ListCertificatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListCertificates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificates", reflect.TypeOf((*MockIOTClient)(nil).ListCertificates), varargs...)
}

func (m *MockIOTClient) ListCertificatesByCA(arg0 context.Context, arg1 *iot.ListCertificatesByCAInput, arg2 ...func(*iot.Options)) (*iot.ListCertificatesByCAOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCertificatesByCA", varargs...)
	ret0, _ := ret[0].(*iot.ListCertificatesByCAOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListCertificatesByCA(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCertificatesByCA", reflect.TypeOf((*MockIOTClient)(nil).ListCertificatesByCA), varargs...)
}

func (m *MockIOTClient) ListJobs(arg0 context.Context, arg1 *iot.ListJobsInput, arg2 ...func(*iot.Options)) (*iot.ListJobsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListJobs", varargs...)
	ret0, _ := ret[0].(*iot.ListJobsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListJobs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJobs", reflect.TypeOf((*MockIOTClient)(nil).ListJobs), varargs...)
}

func (m *MockIOTClient) ListPolicies(arg0 context.Context, arg1 *iot.ListPoliciesInput, arg2 ...func(*iot.Options)) (*iot.ListPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPolicies", varargs...)
	ret0, _ := ret[0].(*iot.ListPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPolicies", reflect.TypeOf((*MockIOTClient)(nil).ListPolicies), varargs...)
}

func (m *MockIOTClient) ListSecurityProfiles(arg0 context.Context, arg1 *iot.ListSecurityProfilesInput, arg2 ...func(*iot.Options)) (*iot.ListSecurityProfilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSecurityProfiles", varargs...)
	ret0, _ := ret[0].(*iot.ListSecurityProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListSecurityProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecurityProfiles", reflect.TypeOf((*MockIOTClient)(nil).ListSecurityProfiles), varargs...)
}

func (m *MockIOTClient) ListStreams(arg0 context.Context, arg1 *iot.ListStreamsInput, arg2 ...func(*iot.Options)) (*iot.ListStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStreams", varargs...)
	ret0, _ := ret[0].(*iot.ListStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStreams", reflect.TypeOf((*MockIOTClient)(nil).ListStreams), varargs...)
}

func (m *MockIOTClient) ListTagsForResource(arg0 context.Context, arg1 *iot.ListTagsForResourceInput, arg2 ...func(*iot.Options)) (*iot.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*iot.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockIOTClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockIOTClient) ListTargetsForSecurityProfile(arg0 context.Context, arg1 *iot.ListTargetsForSecurityProfileInput, arg2 ...func(*iot.Options)) (*iot.ListTargetsForSecurityProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTargetsForSecurityProfile", varargs...)
	ret0, _ := ret[0].(*iot.ListTargetsForSecurityProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListTargetsForSecurityProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTargetsForSecurityProfile", reflect.TypeOf((*MockIOTClient)(nil).ListTargetsForSecurityProfile), varargs...)
}

func (m *MockIOTClient) ListThingGroups(arg0 context.Context, arg1 *iot.ListThingGroupsInput, arg2 ...func(*iot.Options)) (*iot.ListThingGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListThingGroups", varargs...)
	ret0, _ := ret[0].(*iot.ListThingGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListThingGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThingGroups", reflect.TypeOf((*MockIOTClient)(nil).ListThingGroups), varargs...)
}

func (m *MockIOTClient) ListThingPrincipals(arg0 context.Context, arg1 *iot.ListThingPrincipalsInput, arg2 ...func(*iot.Options)) (*iot.ListThingPrincipalsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListThingPrincipals", varargs...)
	ret0, _ := ret[0].(*iot.ListThingPrincipalsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListThingPrincipals(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThingPrincipals", reflect.TypeOf((*MockIOTClient)(nil).ListThingPrincipals), varargs...)
}

func (m *MockIOTClient) ListThingTypes(arg0 context.Context, arg1 *iot.ListThingTypesInput, arg2 ...func(*iot.Options)) (*iot.ListThingTypesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListThingTypes", varargs...)
	ret0, _ := ret[0].(*iot.ListThingTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListThingTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThingTypes", reflect.TypeOf((*MockIOTClient)(nil).ListThingTypes), varargs...)
}

func (m *MockIOTClient) ListThings(arg0 context.Context, arg1 *iot.ListThingsInput, arg2 ...func(*iot.Options)) (*iot.ListThingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListThings", varargs...)
	ret0, _ := ret[0].(*iot.ListThingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListThings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThings", reflect.TypeOf((*MockIOTClient)(nil).ListThings), varargs...)
}

func (m *MockIOTClient) ListThingsInBillingGroup(arg0 context.Context, arg1 *iot.ListThingsInBillingGroupInput, arg2 ...func(*iot.Options)) (*iot.ListThingsInBillingGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListThingsInBillingGroup", varargs...)
	ret0, _ := ret[0].(*iot.ListThingsInBillingGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListThingsInBillingGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThingsInBillingGroup", reflect.TypeOf((*MockIOTClient)(nil).ListThingsInBillingGroup), varargs...)
}

func (m *MockIOTClient) ListThingsInThingGroup(arg0 context.Context, arg1 *iot.ListThingsInThingGroupInput, arg2 ...func(*iot.Options)) (*iot.ListThingsInThingGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListThingsInThingGroup", varargs...)
	ret0, _ := ret[0].(*iot.ListThingsInThingGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListThingsInThingGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThingsInThingGroup", reflect.TypeOf((*MockIOTClient)(nil).ListThingsInThingGroup), varargs...)
}

func (m *MockIOTClient) ListTopicRules(arg0 context.Context, arg1 *iot.ListTopicRulesInput, arg2 ...func(*iot.Options)) (*iot.ListTopicRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTopicRules", varargs...)
	ret0, _ := ret[0].(*iot.ListTopicRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIOTClientMockRecorder) ListTopicRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTopicRules", reflect.TypeOf((*MockIOTClient)(nil).ListTopicRules), varargs...)
}
