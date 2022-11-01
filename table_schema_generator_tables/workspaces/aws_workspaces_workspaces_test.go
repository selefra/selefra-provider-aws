package workspaces

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildWorkspaces(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockWorkspacesClient(ctrl)

	var workspace types.Workspace
	if err := faker.FakeObject(&workspace); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeWorkspaces(
		gomock.Any(),
		&workspaces.DescribeWorkspacesInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&workspaces.DescribeWorkspacesOutput{Workspaces: []types.Workspace{workspace}},
		nil,
	)

	return aws_client.AwsServices{Workspaces: mock}
}

func TestWorkspacesWorkspaces(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWorkspacesWorkspacesGenerator{}), buildWorkspaces, aws_client.TestOptions{})
}
