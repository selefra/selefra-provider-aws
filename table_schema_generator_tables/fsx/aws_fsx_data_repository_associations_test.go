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

func buildDataRepoAssociationsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockFsxClient(ctrl)

	var a types.DataRepositoryAssociation
	require.NoError(t, faker.FakeObject(&a))
	m.EXPECT().DescribeDataRepositoryAssociations(
		gomock.Any(),
		&fsx.DescribeDataRepositoryAssociationsInput{MaxResults: aws.Int32(25)},
	).AnyTimes().Return(
		&fsx.DescribeDataRepositoryAssociationsOutput{Associations: []types.DataRepositoryAssociation{a}},
		nil,
	)

	return aws_client.AwsServices{
		FSX: m,
	}
}

func TestDataRepoAssociations(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFsxDataRepositoryAssociationsGenerator{}), buildDataRepoAssociationsMock, aws_client.TestOptions{})
}
