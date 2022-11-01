package ecrpublic

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	"github.com/aws/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEcrPublicRepositoriesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEcrPublicClient(ctrl)
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
		&ecrpublic.DescribeRepositoriesOutput{
			Repositories: []types.Repository{l},
		}, nil)

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ecrpublic.DescribeImagesOutput{
			ImageDetails: []types.ImageDetail{i},
		}, nil)

	tagResponse := ecrpublic.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagResponse)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tagResponse, nil)

	return aws_client.AwsServices{
		ECRPublic: m,
	}
}

func TestEcrPublicRepositories(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEcrpublicRepositoriesGenerator{}), buildEcrPublicRepositoriesMock, aws_client.TestOptions{})
}
