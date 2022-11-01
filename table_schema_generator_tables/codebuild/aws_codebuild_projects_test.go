package codebuild

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildCodebuildProjects(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockCodebuildClient(ctrl)

	projectsList := codebuild.ListProjectsOutput{}
	if err := faker.FakeObject(&projectsList); err != nil {
		t.Fatal(err)
	}
	projectsList.NextToken = nil
	m.EXPECT().ListProjects(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(
		&projectsList,
		nil,
	)

	projects := codebuild.BatchGetProjectsOutput{}
	if err := faker.FakeObject(&projects); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().BatchGetProjects(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(
		&projects,
		nil,
	)

	return aws_client.AwsServices{Codebuild: m}
}

func TestCodebuildProjects(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCodebuildProjectsGenerator{}), buildCodebuildProjects, aws_client.TestOptions{})
}
