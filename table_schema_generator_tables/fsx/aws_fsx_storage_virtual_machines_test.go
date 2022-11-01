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
	"github.com/stretchr/testify/require"
)

func buildStorageVmsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockFsxClient(ctrl)

	var vm types.StorageVirtualMachine
	require.NoError(t, faker.FakeObject(&vm))
	m.EXPECT().DescribeStorageVirtualMachines(
		gomock.Any(),
		&fsx.DescribeStorageVirtualMachinesInput{MaxResults: aws.Int32(1000)},
	).AnyTimes().Return(
		&fsx.DescribeStorageVirtualMachinesOutput{StorageVirtualMachines: []types.StorageVirtualMachine{vm}},
		nil,
	)
	return aws_client.AwsServices{
		FSX: m,
	}
}

func TestStorageVms(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFsxStorageVirtualMachinesGenerator{}), buildStorageVmsMock, aws_client.TestOptions{})
}
