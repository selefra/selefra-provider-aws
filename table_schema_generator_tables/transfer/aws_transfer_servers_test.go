package transfer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func buildServersMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockTransferClient(ctrl)

	var ls types.ListedServer
	require.NoError(t, faker.FakeObject(&ls))
	m.EXPECT().ListServers(
		gomock.Any(),
		&transfer.ListServersInput{MaxResults: aws.Int32(1000)},
	).AnyTimes().Return(
		&transfer.ListServersOutput{Servers: []types.ListedServer{ls}},
		nil,
	)

	var ds types.DescribedServer
	require.NoError(t, faker.FakeObject(&ds))
	ds.ServerId = ls.ServerId
	ds.Arn = ls.Arn
	m.EXPECT().DescribeServer(
		gomock.Any(),
		&transfer.DescribeServerInput{ServerId: ls.ServerId},
	).AnyTimes().Return(
		&transfer.DescribeServerOutput{Server: &ds},
		nil,
	)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&transfer.ListTagsForResourceInput{Arn: ds.Arn},
	).AnyTimes().Return(
		&transfer.ListTagsForResourceOutput{Tags: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}}},
		nil,
	)

	return aws_client.AwsServices{
		Transfer: m,
	}
}

func TestServers(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsTransferServersGenerator{}), buildServersMock, aws_client.TestOptions{})
}
