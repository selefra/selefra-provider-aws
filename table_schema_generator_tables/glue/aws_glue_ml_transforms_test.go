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

func buildMlTransformsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlueClient(ctrl)

	var transforms glue.GetMLTransformsOutput
	require.NoError(t, faker.FakeObject(&transforms))
	transforms.NextToken = nil
	m.EXPECT().GetMLTransforms(
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(&transforms, nil)

	var runs glue.GetMLTaskRunsOutput
	require.NoError(t, faker.FakeObject(&runs))
	runs.NextToken = nil
	m.EXPECT().GetMLTaskRuns(
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(&runs, nil)

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

func TestMlTransforms(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlueMlTransformsGenerator{}), buildMlTransformsMock, aws_client.TestOptions{})
}
