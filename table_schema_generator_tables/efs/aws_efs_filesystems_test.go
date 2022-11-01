package efs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEfsFilesystemsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEfsClient(ctrl)
	l := types.FileSystemDescription{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeFileSystems(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&efs.DescribeFileSystemsOutput{
			FileSystems: []types.FileSystemDescription{l},
		}, nil)

	b := efs.DescribeBackupPolicyOutput{}
	err = faker.FakeObject(&b)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeBackupPolicy(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&b, nil)

	return aws_client.AwsServices{
		EFS: m,
	}
}

func TestEfsFilesystems(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEfsFilesystemsGenerator{}), buildEfsFilesystemsMock, aws_client.TestOptions{})
}
