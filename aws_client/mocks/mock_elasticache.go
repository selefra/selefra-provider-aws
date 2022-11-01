package mocks

import (
	context "context"
	reflect "reflect"

	elasticache "github.com/aws/aws-sdk-go-v2/service/elasticache"
	gomock "github.com/golang/mock/gomock"
)

type MockElastiCache struct {
	ctrl		*gomock.Controller
	recorder	*MockElastiCacheMockRecorder
}

type MockElastiCacheMockRecorder struct {
	mock *MockElastiCache
}

func NewMockElastiCache(ctrl *gomock.Controller) *MockElastiCache {
	mock := &MockElastiCache{ctrl: ctrl}
	mock.recorder = &MockElastiCacheMockRecorder{mock}
	return mock
}

func (m *MockElastiCache) EXPECT() *MockElastiCacheMockRecorder {
	return m.recorder
}

func (m *MockElastiCache) DescribeCacheClusters(arg0 context.Context, arg1 *elasticache.DescribeCacheClustersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheClustersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheClusters", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeCacheClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheClusters", reflect.TypeOf((*MockElastiCache)(nil).DescribeCacheClusters), varargs...)
}

func (m *MockElastiCache) DescribeCacheEngineVersions(arg0 context.Context, arg1 *elasticache.DescribeCacheEngineVersionsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheEngineVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheEngineVersions", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheEngineVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeCacheEngineVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheEngineVersions", reflect.TypeOf((*MockElastiCache)(nil).DescribeCacheEngineVersions), varargs...)
}

func (m *MockElastiCache) DescribeCacheParameterGroups(arg0 context.Context, arg1 *elasticache.DescribeCacheParameterGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheParameterGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheParameterGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheParameterGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeCacheParameterGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheParameterGroups", reflect.TypeOf((*MockElastiCache)(nil).DescribeCacheParameterGroups), varargs...)
}

func (m *MockElastiCache) DescribeCacheParameters(arg0 context.Context, arg1 *elasticache.DescribeCacheParametersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheParameters", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeCacheParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheParameters", reflect.TypeOf((*MockElastiCache)(nil).DescribeCacheParameters), varargs...)
}

func (m *MockElastiCache) DescribeCacheSubnetGroups(arg0 context.Context, arg1 *elasticache.DescribeCacheSubnetGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeCacheSubnetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCacheSubnetGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeCacheSubnetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeCacheSubnetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCacheSubnetGroups", reflect.TypeOf((*MockElastiCache)(nil).DescribeCacheSubnetGroups), varargs...)
}

func (m *MockElastiCache) DescribeGlobalReplicationGroups(arg0 context.Context, arg1 *elasticache.DescribeGlobalReplicationGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeGlobalReplicationGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeGlobalReplicationGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeGlobalReplicationGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeGlobalReplicationGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeGlobalReplicationGroups", reflect.TypeOf((*MockElastiCache)(nil).DescribeGlobalReplicationGroups), varargs...)
}

func (m *MockElastiCache) DescribeReplicationGroups(arg0 context.Context, arg1 *elasticache.DescribeReplicationGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReplicationGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReplicationGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReplicationGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeReplicationGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReplicationGroups", reflect.TypeOf((*MockElastiCache)(nil).DescribeReplicationGroups), varargs...)
}

func (m *MockElastiCache) DescribeReservedCacheNodes(arg0 context.Context, arg1 *elasticache.DescribeReservedCacheNodesInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReservedCacheNodes", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReservedCacheNodesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeReservedCacheNodes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReservedCacheNodes", reflect.TypeOf((*MockElastiCache)(nil).DescribeReservedCacheNodes), varargs...)
}

func (m *MockElastiCache) DescribeReservedCacheNodesOfferings(arg0 context.Context, arg1 *elasticache.DescribeReservedCacheNodesOfferingsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeReservedCacheNodesOfferings", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeReservedCacheNodesOfferingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeReservedCacheNodesOfferings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeReservedCacheNodesOfferings", reflect.TypeOf((*MockElastiCache)(nil).DescribeReservedCacheNodesOfferings), varargs...)
}

func (m *MockElastiCache) DescribeServiceUpdates(arg0 context.Context, arg1 *elasticache.DescribeServiceUpdatesInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeServiceUpdatesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeServiceUpdates", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeServiceUpdatesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeServiceUpdates(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServiceUpdates", reflect.TypeOf((*MockElastiCache)(nil).DescribeServiceUpdates), varargs...)
}

func (m *MockElastiCache) DescribeSnapshots(arg0 context.Context, arg1 *elasticache.DescribeSnapshotsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSnapshots", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSnapshots", reflect.TypeOf((*MockElastiCache)(nil).DescribeSnapshots), varargs...)
}

func (m *MockElastiCache) DescribeUserGroups(arg0 context.Context, arg1 *elasticache.DescribeUserGroupsInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeUserGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUserGroups", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeUserGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeUserGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUserGroups", reflect.TypeOf((*MockElastiCache)(nil).DescribeUserGroups), varargs...)
}

func (m *MockElastiCache) DescribeUsers(arg0 context.Context, arg1 *elasticache.DescribeUsersInput, arg2 ...func(*elasticache.Options)) (*elasticache.DescribeUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUsers", varargs...)
	ret0, _ := ret[0].(*elasticache.DescribeUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockElastiCacheMockRecorder) DescribeUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUsers", reflect.TypeOf((*MockElastiCache)(nil).DescribeUsers), varargs...)
}
