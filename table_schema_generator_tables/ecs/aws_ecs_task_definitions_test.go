package ecs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEcsTaskDefinitions(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEcsClient(ctrl)

	listTaskDefinitionsOutput := ecs.ListTaskDefinitionsOutput{}
	err := faker.FakeObject(&listTaskDefinitionsOutput)
	if err != nil {
		t.Fatal(err)
	}
	listTaskDefinitionsOutput.NextToken = nil
	m.EXPECT().ListTaskDefinitions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&listTaskDefinitionsOutput, nil)

	taskDefinition := &ecs.DescribeTaskDefinitionOutput{}
	err = faker.FakeObject(&taskDefinition)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTaskDefinition(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(taskDefinition, nil)

	return aws_client.AwsServices{
		ECS: m,
	}
}

func TestEcsTaskDefinitions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEcsTaskDefinitionsGenerator{}), buildEcsTaskDefinitions, aws_client.TestOptions{})
}
