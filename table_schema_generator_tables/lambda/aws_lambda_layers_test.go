package lambda

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildLambdaLayersMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockLambdaClient(ctrl)

	creationDate := "1994-11-05T08:15:30.000+0500"

	l := types.LayersListItem{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	l.LatestMatchingVersion.CreatedDate = &creationDate
	m.EXPECT().ListLayers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lambda.ListLayersOutput{
			Layers: []types.LayersListItem{l},
		}, nil)

	lv := types.LayerVersionsListItem{}
	err = faker.FakeObject(&lv)
	if err != nil {
		t.Fatal(err)
	}
	arn := "arn:aws:s3:::my_corporate_bucket/test:exampleobject.png:1"
	lv.LayerVersionArn = &arn
	lv.CreatedDate = &creationDate
	m.EXPECT().ListLayerVersions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lambda.ListLayerVersionsOutput{
			LayerVersions: []types.LayerVersionsListItem{lv},
		}, nil)

	lvp := lambda.GetLayerVersionPolicyOutput{}
	err = faker.FakeObject(&lvp)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetLayerVersionPolicy(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lvp, nil)

	return aws_client.AwsServices{
		Lambda: m,
	}
}

func TestLambdaLayers(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLambdaLayersGenerator{}), buildLambdaLayersMock, aws_client.TestOptions{})
}
