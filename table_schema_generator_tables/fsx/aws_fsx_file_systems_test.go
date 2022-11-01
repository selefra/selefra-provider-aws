package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildFilesystemsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockFsxClient(ctrl)

	var f types.FileSystem
	err := faker.FakeObject(&f, faker.WithMaxDepth(5))
	if err != nil {
		t.Fatalf("FakeObject returned error: %v", err)
	}
	f.FileSystemType = types.FileSystemTypeLustre
	f.Lifecycle = types.FileSystemLifecycleAvailable
	f.StorageType = types.StorageTypeHdd
	m.EXPECT().DescribeFileSystems(
		gomock.Any(),
		&fsx.DescribeFileSystemsInput{MaxResults: aws.Int32(1000)},
	).AnyTimes().Return(
		&fsx.DescribeFileSystemsOutput{FileSystems: []types.FileSystem{f}},
		nil,
	)

	return aws_client.AwsServices{
		FSX: m,
	}
}

func TestFilesystems(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFsxFileSystemsGenerator{}), buildFilesystemsMock, aws_client.TestOptions{})
}
