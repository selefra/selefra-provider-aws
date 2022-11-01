package mocks

import (
	context "context"
	reflect "reflect"

	wafregional "github.com/aws/aws-sdk-go-v2/service/wafregional"
	gomock "github.com/golang/mock/gomock"
)

type MockWafRegionalClient struct {
	ctrl		*gomock.Controller
	recorder	*MockWafRegionalClientMockRecorder
}

type MockWafRegionalClientMockRecorder struct {
	mock *MockWafRegionalClient
}

func NewMockWafRegionalClient(ctrl *gomock.Controller) *MockWafRegionalClient {
	mock := &MockWafRegionalClient{ctrl: ctrl}
	mock.recorder = &MockWafRegionalClientMockRecorder{mock}
	return mock
}

func (m *MockWafRegionalClient) EXPECT() *MockWafRegionalClientMockRecorder {
	return m.recorder
}

func (m *MockWafRegionalClient) GetRateBasedRule(arg0 context.Context, arg1 *wafregional.GetRateBasedRuleInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRateBasedRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRateBasedRule", varargs...)
	ret0, _ := ret[0].(*wafregional.GetRateBasedRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafRegionalClientMockRecorder) GetRateBasedRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRateBasedRule", reflect.TypeOf((*MockWafRegionalClient)(nil).GetRateBasedRule), varargs...)
}

func (m *MockWafRegionalClient) GetRule(arg0 context.Context, arg1 *wafregional.GetRuleInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRuleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRule", varargs...)
	ret0, _ := ret[0].(*wafregional.GetRuleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafRegionalClientMockRecorder) GetRule(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRule", reflect.TypeOf((*MockWafRegionalClient)(nil).GetRule), varargs...)
}

func (m *MockWafRegionalClient) GetRuleGroup(arg0 context.Context, arg1 *wafregional.GetRuleGroupInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetRuleGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRuleGroup", varargs...)
	ret0, _ := ret[0].(*wafregional.GetRuleGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafRegionalClientMockRecorder) GetRuleGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRuleGroup", reflect.TypeOf((*MockWafRegionalClient)(nil).GetRuleGroup), varargs...)
}

func (m *MockWafRegionalClient) GetWebACL(arg0 context.Context, arg1 *wafregional.GetWebACLInput, arg2 ...func(*wafregional.Options)) (*wafregional.GetWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWebACL", varargs...)
	ret0, _ := ret[0].(*wafregional.GetWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafRegionalClientMockRecorder) GetWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWebACL", reflect.TypeOf((*MockWafRegionalClient)(nil).GetWebACL), varargs...)
}

func (m *MockWafRegionalClient) ListRateBasedRules(arg0 context.Context, arg1 *wafregional.ListRateBasedRulesInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRateBasedRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRateBasedRules", varargs...)
	ret0, _ := ret[0].(*wafregional.ListRateBasedRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafRegionalClientMockRecorder) ListRateBasedRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRateBasedRules", reflect.TypeOf((*MockWafRegionalClient)(nil).ListRateBasedRules), varargs...)
}

func (m *MockWafRegionalClient) ListResourcesForWebACL(arg0 context.Context, arg1 *wafregional.ListResourcesForWebACLInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListResourcesForWebACLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourcesForWebACL", varargs...)
	ret0, _ := ret[0].(*wafregional.ListResourcesForWebACLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafRegionalClientMockRecorder) ListResourcesForWebACL(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourcesForWebACL", reflect.TypeOf((*MockWafRegionalClient)(nil).ListResourcesForWebACL), varargs...)
}

func (m *MockWafRegionalClient) ListRuleGroups(arg0 context.Context, arg1 *wafregional.ListRuleGroupsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRuleGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRuleGroups", varargs...)
	ret0, _ := ret[0].(*wafregional.ListRuleGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafRegionalClientMockRecorder) ListRuleGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRuleGroups", reflect.TypeOf((*MockWafRegionalClient)(nil).ListRuleGroups), varargs...)
}

func (m *MockWafRegionalClient) ListRules(arg0 context.Context, arg1 *wafregional.ListRulesInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRules", varargs...)
	ret0, _ := ret[0].(*wafregional.ListRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafRegionalClientMockRecorder) ListRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRules", reflect.TypeOf((*MockWafRegionalClient)(nil).ListRules), varargs...)
}

func (m *MockWafRegionalClient) ListTagsForResource(arg0 context.Context, arg1 *wafregional.ListTagsForResourceInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*wafregional.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafRegionalClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockWafRegionalClient)(nil).ListTagsForResource), varargs...)
}

func (m *MockWafRegionalClient) ListWebACLs(arg0 context.Context, arg1 *wafregional.ListWebACLsInput, arg2 ...func(*wafregional.Options)) (*wafregional.ListWebACLsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListWebACLs", varargs...)
	ret0, _ := ret[0].(*wafregional.ListWebACLsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWafRegionalClientMockRecorder) ListWebACLs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWebACLs", reflect.TypeOf((*MockWafRegionalClient)(nil).ListWebACLs), varargs...)
}
