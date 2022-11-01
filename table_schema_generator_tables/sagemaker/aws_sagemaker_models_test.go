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

func buildSageMakerModels(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockSageMakerClient(ctrl)

	summ := types.ModelSummary{}
	if err := faker.FakeObject(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListModels(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sagemaker.ListModelsOutput{Models: []types.ModelSummary{summ}},
		nil,
	)

	model := sagemaker.DescribeModelOutput{}
	if err := faker.FakeObject(&model); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeModel(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&model,
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

func TestSageMakerModels(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSagemakerModelsGenerator{}), buildSageMakerModels, aws_client.TestOptions{})
}
