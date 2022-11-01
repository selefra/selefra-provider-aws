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

func buildSageMakerEndpointConfigs(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockSageMakerClient(ctrl)

	summ := types.EndpointConfigSummary{}
	if err := faker.FakeObject(&summ); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListEndpointConfigs(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sagemaker.ListEndpointConfigsOutput{EndpointConfigs: []types.EndpointConfigSummary{summ}},
		nil,
	)

	endpointConfig := sagemaker.DescribeEndpointConfigOutput{}
	if err := faker.FakeObject(&endpointConfig); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeEndpointConfig(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&endpointConfig,
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

func TestSageMakerEndpointConfigurations(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSagemakerEndpointConfigurationsGenerator{}), buildSageMakerEndpointConfigs, aws_client.TestOptions{})
}
