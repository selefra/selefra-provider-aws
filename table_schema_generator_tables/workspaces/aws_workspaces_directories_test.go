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

func buildDirectories(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockWorkspacesClient(ctrl)

	var directory types.WorkspaceDirectory
	if err := faker.FakeObject(&directory); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeWorkspaceDirectories(
		gomock.Any(),
		&workspaces.DescribeWorkspaceDirectoriesInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&workspaces.DescribeWorkspaceDirectoriesOutput{Directories: []types.WorkspaceDirectory{directory}},
		nil,
	)

	return aws_client.AwsServices{Workspaces: mock}
}

func TestWorkspacesDirectories(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWorkspacesDirectoriesGenerator{}), buildDirectories, aws_client.TestOptions{})
}
