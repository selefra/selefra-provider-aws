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

func buildSageMakerNotebookInstances(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockSageMakerClient(ctrl)

	summ := types.NotebookInstanceSummary{}
	if err := faker.FakeObject(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListNotebookInstances(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sagemaker.ListNotebookInstancesOutput{NotebookInstances: []types.NotebookInstanceSummary{summ}},
		nil,
	)

	note := sagemaker.DescribeNotebookInstanceOutput{}
	if err := faker.FakeObject(&note); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeNotebookInstance(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
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

func TestSageMakerNotebookInstances(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSagemakerNotebookInstancesGenerator{}), buildSageMakerNotebookInstances, aws_client.TestOptions{})
}
