package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func buildDevEndpointsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlueClient(ctrl)

	var devEndpoint glue.GetDevEndpointsOutput
	require.NoError(t, faker.FakeObject(&devEndpoint))
	devEndpoint.NextToken = nil
	m.EXPECT().GetDevEndpoints(
		gomock.Any(),
		&glue.GetDevEndpointsInput{},
	).AnyTimes().Return(&devEndpoint, nil)

	m.EXPECT().GetTags(
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(
		&glue.GetTagsOutput{Tags: map[string]string{"key": "value"}},
		nil,
	)

	return aws_client.AwsServices{
		Glue: m,
	}
}

func TestDevEndpoints(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlueDevEndpointsGenerator{}), buildDevEndpointsMock, aws_client.TestOptions{})
}
