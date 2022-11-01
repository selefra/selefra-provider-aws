package mocks

import (
	context "context"
	reflect "reflect"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/golang/mock/gomock"
)

type MockS3Client struct {
	ctrl		*gomock.Controller
	recorder	*MockS3ClientMockRecorder
}

type MockS3ClientMockRecorder struct {
	mock *MockS3Client
}

func NewMockS3Client(ctrl *gomock.Controller) *MockS3Client {
	mock := &MockS3Client{ctrl: ctrl}
	mock.recorder = &MockS3ClientMockRecorder{mock}
	return mock
}

func (m *MockS3Client) EXPECT() *MockS3ClientMockRecorder {
	return m.recorder
}

func (m *MockS3Client) GetBucketAcl(arg0 context.Context, arg1 *s3.GetBucketAclInput, arg2 ...func(*s3.Options)) (*s3.GetBucketAclOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketAcl", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketAclOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketAcl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketAcl", reflect.TypeOf((*MockS3Client)(nil).GetBucketAcl), varargs...)
}

func (m *MockS3Client) GetBucketCors(arg0 context.Context, arg1 *s3.GetBucketCorsInput, arg2 ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketCors", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketCorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketCors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketCors", reflect.TypeOf((*MockS3Client)(nil).GetBucketCors), varargs...)
}

func (m *MockS3Client) GetBucketEncryption(arg0 context.Context, arg1 *s3.GetBucketEncryptionInput, arg2 ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketEncryption", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketEncryptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketEncryption(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketEncryption", reflect.TypeOf((*MockS3Client)(nil).GetBucketEncryption), varargs...)
}

func (m *MockS3Client) GetBucketLifecycleConfiguration(arg0 context.Context, arg1 *s3.GetBucketLifecycleConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketLifecycleConfiguration", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketLifecycleConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketLifecycleConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketLifecycleConfiguration", reflect.TypeOf((*MockS3Client)(nil).GetBucketLifecycleConfiguration), varargs...)
}

func (m *MockS3Client) GetBucketLocation(arg0 context.Context, arg1 *s3.GetBucketLocationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketLocation", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketLocationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketLocation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketLocation", reflect.TypeOf((*MockS3Client)(nil).GetBucketLocation), varargs...)
}

func (m *MockS3Client) GetBucketLogging(arg0 context.Context, arg1 *s3.GetBucketLoggingInput, arg2 ...func(*s3.Options)) (*s3.GetBucketLoggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketLogging", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketLoggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketLogging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketLogging", reflect.TypeOf((*MockS3Client)(nil).GetBucketLogging), varargs...)
}

func (m *MockS3Client) GetBucketOwnershipControls(arg0 context.Context, arg1 *s3.GetBucketOwnershipControlsInput, arg2 ...func(*s3.Options)) (*s3.GetBucketOwnershipControlsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketOwnershipControls", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketOwnershipControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketOwnershipControls(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketOwnershipControls", reflect.TypeOf((*MockS3Client)(nil).GetBucketOwnershipControls), varargs...)
}

func (m *MockS3Client) GetBucketPolicy(arg0 context.Context, arg1 *s3.GetBucketPolicyInput, arg2 ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketPolicy", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketPolicy", reflect.TypeOf((*MockS3Client)(nil).GetBucketPolicy), varargs...)
}

func (m *MockS3Client) GetBucketReplication(arg0 context.Context, arg1 *s3.GetBucketReplicationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketReplicationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketReplication", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketReplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketReplication(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketReplication", reflect.TypeOf((*MockS3Client)(nil).GetBucketReplication), varargs...)
}

func (m *MockS3Client) GetBucketTagging(arg0 context.Context, arg1 *s3.GetBucketTaggingInput, arg2 ...func(*s3.Options)) (*s3.GetBucketTaggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketTagging", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketTaggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketTagging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketTagging", reflect.TypeOf((*MockS3Client)(nil).GetBucketTagging), varargs...)
}

func (m *MockS3Client) GetBucketVersioning(arg0 context.Context, arg1 *s3.GetBucketVersioningInput, arg2 ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketVersioning", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketVersioningOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetBucketVersioning(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketVersioning", reflect.TypeOf((*MockS3Client)(nil).GetBucketVersioning), varargs...)
}

func (m *MockS3Client) GetPublicAccessBlock(arg0 context.Context, arg1 *s3.GetPublicAccessBlockInput, arg2 ...func(*s3.Options)) (*s3.GetPublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPublicAccessBlock", varargs...)
	ret0, _ := ret[0].(*s3.GetPublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) GetPublicAccessBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicAccessBlock", reflect.TypeOf((*MockS3Client)(nil).GetPublicAccessBlock), varargs...)
}

func (m *MockS3Client) ListBuckets(arg0 context.Context, arg1 *s3.ListBucketsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBuckets", varargs...)
	ret0, _ := ret[0].(*s3.ListBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockS3ClientMockRecorder) ListBuckets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockS3Client)(nil).ListBuckets), varargs...)
}
