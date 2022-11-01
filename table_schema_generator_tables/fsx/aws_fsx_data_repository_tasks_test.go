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

func buildDataRepoTasksMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockFsxClient(ctrl)

	var task types.DataRepositoryTask
	require.NoError(t, faker.FakeObject(&task))
	m.EXPECT().DescribeDataRepositoryTasks(
		gomock.Any(),
		&fsx.DescribeDataRepositoryTasksInput{MaxResults: aws.Int32(1000)},
	).AnyTimes().Return(
		&fsx.DescribeDataRepositoryTasksOutput{DataRepositoryTasks: []types.DataRepositoryTask{task}},
		nil,
	)

	return aws_client.AwsServices{
		FSX: m,
	}
}

func TestDataRepoTasks(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFsxDataRepositoryTasksGenerator{}), buildDataRepoTasksMock, aws_client.TestOptions{})
}
