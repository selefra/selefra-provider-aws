package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func buildWorkflowsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlueClient(ctrl)

	var name string
	require.NoError(t, faker.FakeObject(&name))
	m.EXPECT().ListWorkflows(
		gomock.Any(),
		&glue.ListWorkflowsInput{MaxResults: aws.Int32(25)},
	).AnyTimes().Return(
		&glue.ListWorkflowsOutput{Workflows: []string{name}},
		nil,
	)

	var w types.Workflow
	require.NoError(t, faker.FakeObject(&w))
	w.Name = &name
	m.EXPECT().GetWorkflow(
		gomock.Any(),
		&glue.GetWorkflowInput{Name: aws.String(name)},
	).AnyTimes().Return(
		&glue.GetWorkflowOutput{Workflow: &w},
		nil,
	)

	m.EXPECT().GetTags(
		gomock.Any(),
		&glue.GetTagsInput{ResourceArn: aws.String("arn:aws:glue:us-east-1:testAccount:workflow/" + name)},
	).AnyTimes().Return(
		&glue.GetTagsOutput{Tags: map[string]string{"key": "value"}},
		nil,
	)

	return aws_client.AwsServices{
		Glue: m,
	}
}

func TestWorkflows(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlueWorkflowsGenerator{}), buildWorkflowsMock, aws_client.TestOptions{})
}
