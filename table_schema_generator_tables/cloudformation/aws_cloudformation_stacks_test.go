package cloudformation

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildStacks(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockCloudFormationClient(ctrl)

	var stack types.Stack
	if err := faker.FakeObject(&stack); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeStacks(
		gomock.Any(),
		&cloudformation.DescribeStacksInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&cloudformation.DescribeStacksOutput{Stacks: []types.Stack{stack}},
		nil,
	)

	var resource types.StackResourceSummary
	if err := faker.FakeObject(&resource); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListStackResources(
		gomock.Any(),
		&cloudformation.ListStackResourcesInput{StackName: stack.StackName},
		gomock.Any(),
	).AnyTimes().Return(
		&cloudformation.ListStackResourcesOutput{StackResourceSummaries: []types.StackResourceSummary{resource}},
		nil,
	)

	return aws_client.AwsServices{Cloudformation: mock}
}

func TestCloudformationStacks(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCloudformationStacksGenerator{}), buildStacks, aws_client.TestOptions{})
}
