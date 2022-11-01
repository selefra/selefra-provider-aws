package sagemaker

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	types "github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildSageMakerTrainingJobs(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockSageMakerClient(ctrl)

	summ := types.TrainingJobSummary{}
	if err := faker.FakeObject(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTrainingJobs(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sagemaker.ListTrainingJobsOutput{TrainingJobSummaries: []types.TrainingJobSummary{summ}},
		nil,
	)

	note := sagemaker.DescribeTrainingJobOutput{}
	if err := faker.FakeObject(&note); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTrainingJob(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&note,
		nil,
	)

	var tagsOut sagemaker.ListTagsOutput
	if err := faker.FakeObject(&tagsOut); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&tagsOut, nil,
	)

	return aws_client.AwsServices{
		SageMaker: m,
	}
}

func TestSageMakerTrainingJobs(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSagemakerTrainingJobsGenerator{}), buildSageMakerTrainingJobs, aws_client.TestOptions{})
}
