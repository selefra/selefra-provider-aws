package codepipeline

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildPipelines(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockCodePipelineClient(ctrl)

	var pipeSummary types.PipelineSummary
	if err := faker.FakeObject(&pipeSummary); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListPipelines(
		gomock.Any(),
		&codepipeline.ListPipelinesInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&codepipeline.ListPipelinesOutput{Pipelines: []types.PipelineSummary{pipeSummary}},
		nil,
	)

	var resource codepipeline.GetPipelineOutput
	if err := faker.FakeObject(&resource); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetPipeline(
		gomock.Any(),
		&codepipeline.GetPipelineInput{Name: pipeSummary.Name},
		gomock.Any(),
	).AnyTimes().Return(
		&resource,
		nil,
	)

	tags := &codepipeline.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		tags,
		nil,
	)

	return aws_client.AwsServices{CodePipeline: mock}
}

func TestCodePipelinePipelines(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCodepipelinePipelinesGenerator{}), buildPipelines, aws_client.TestOptions{})
}
