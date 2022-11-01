package mocks

import (
	context "context"
	reflect "reflect"

	fsx "github.com/aws/aws-sdk-go-v2/service/fsx"
	gomock "github.com/golang/mock/gomock"
)

type MockFsxClient struct {
	ctrl		*gomock.Controller
	recorder	*MockFsxClientMockRecorder
}

type MockFsxClientMockRecorder struct {
	mock *MockFsxClient
}

func NewMockFsxClient(ctrl *gomock.Controller) *MockFsxClient {
	mock := &MockFsxClient{ctrl: ctrl}
	mock.recorder = &MockFsxClientMockRecorder{mock}
	return mock
}

func (m *MockFsxClient) EXPECT() *MockFsxClientMockRecorder {
	return m.recorder
}

func (m *MockFsxClient) DescribeBackups(arg0 context.Context, arg1 *fsx.DescribeBackupsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeBackupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBackups", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeBackups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackups", reflect.TypeOf((*MockFsxClient)(nil).DescribeBackups), varargs...)
}

func (m *MockFsxClient) DescribeDataRepositoryAssociations(arg0 context.Context, arg1 *fsx.DescribeDataRepositoryAssociationsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDataRepositoryAssociations", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeDataRepositoryAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeDataRepositoryAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDataRepositoryAssociations", reflect.TypeOf((*MockFsxClient)(nil).DescribeDataRepositoryAssociations), varargs...)
}

func (m *MockFsxClient) DescribeDataRepositoryTasks(arg0 context.Context, arg1 *fsx.DescribeDataRepositoryTasksInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeDataRepositoryTasksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDataRepositoryTasks", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeDataRepositoryTasksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeDataRepositoryTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDataRepositoryTasks", reflect.TypeOf((*MockFsxClient)(nil).DescribeDataRepositoryTasks), varargs...)
}

func (m *MockFsxClient) DescribeFileSystems(arg0 context.Context, arg1 *fsx.DescribeFileSystemsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeFileSystemsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFileSystems", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeFileSystemsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeFileSystems(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFileSystems", reflect.TypeOf((*MockFsxClient)(nil).DescribeFileSystems), varargs...)
}

func (m *MockFsxClient) DescribeSnapshots(arg0 context.Context, arg1 *fsx.DescribeSnapshotsInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSnapshots", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSnapshots", reflect.TypeOf((*MockFsxClient)(nil).DescribeSnapshots), varargs...)
}

func (m *MockFsxClient) DescribeStorageVirtualMachines(arg0 context.Context, arg1 *fsx.DescribeStorageVirtualMachinesInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeStorageVirtualMachinesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStorageVirtualMachines", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeStorageVirtualMachinesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeStorageVirtualMachines(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStorageVirtualMachines", reflect.TypeOf((*MockFsxClient)(nil).DescribeStorageVirtualMachines), varargs...)
}

func (m *MockFsxClient) DescribeVolumes(arg0 context.Context, arg1 *fsx.DescribeVolumesInput, arg2 ...func(*fsx.Options)) (*fsx.DescribeVolumesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVolumes", varargs...)
	ret0, _ := ret[0].(*fsx.DescribeVolumesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFsxClientMockRecorder) DescribeVolumes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVolumes", reflect.TypeOf((*MockFsxClient)(nil).DescribeVolumes), varargs...)
}
