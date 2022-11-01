package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEcrRepositoriesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEcrClient(ctrl)
	l := types.Repository{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	i := types.ImageDetail{}
	err = faker.FakeObject(&i)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeRepositories(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ecr.DescribeRepositoriesOutput{
			Repositories: []types.Repository{l},
		}, nil)

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ecr.DescribeImagesOutput{
			ImageDetails: []types.ImageDetail{i},
		}, nil)

	tagResponse := ecr.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagResponse)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tagResponse, nil)

	return aws_client.AwsServices{
		ECR: m,
	}
}

func TestEcrRepositories(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEcrRepositoriesGenerator{}), buildEcrRepositoriesMock, aws_client.TestOptions{})
}
