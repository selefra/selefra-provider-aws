package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func buildJobsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlueClient(ctrl)

	node := types.CodeGenConfigurationNode{}

	require.NoError(t, faker.FakeObject(&node))
	job := types.Job{
		CodeGenConfigurationNodes:	map[string]types.CodeGenConfigurationNode{"test": node},
		ExecutionClass:			types.ExecutionClassFlex,
	}
	require.NoError(t, faker.FakeObject(&job))
	m.EXPECT().GetJobs(gomock.Any(), gomock.Any()).AnyTimes().Return(&glue.GetJobsOutput{Jobs: []types.Job{job}}, nil)

	var jobRuns glue.GetJobRunsOutput
	require.NoError(t, faker.FakeObject(&jobRuns))
	jobRuns.NextToken = nil
	m.EXPECT().GetJobRuns(gomock.Any(), gomock.Any()).AnyTimes().Return(&jobRuns, nil)

	m.EXPECT().GetTags(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&glue.GetTagsOutput{Tags: map[string]string{"key": "value"}},
		nil,
	)

	return aws_client.AwsServices{
		Glue: m,
	}
}

func TestJobs(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlueJobsGenerator{}), buildJobsMock, aws_client.TestOptions{})
}
