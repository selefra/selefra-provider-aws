package cloudfront

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildCloudfrontDistributionsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := aws_client.AwsServices{
		Cloudfront: m,
	}
	ds := cloudfrontTypes.DistributionSummary{}
	if err := faker.FakeObject(&ds); err != nil {
		t.Fatal(err)
	}
	cloudfrontOutput := &cloudfront.ListDistributionsOutput{
		DistributionList: &cloudfrontTypes.DistributionList{
			Items: []cloudfrontTypes.DistributionSummary{ds},
		},
	}
	m.EXPECT().ListDistributions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		cloudfrontOutput,
		nil,
	)

	distribution := &cloudfront.GetDistributionOutput{}
	if err := faker.FakeObject(&distribution); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDistribution(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		distribution,
		nil,
	)

	tags := &cloudfront.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		tags,
		nil,
	)
	return services
}

func TestCloudfrontDistributions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCloudfrontDistributionsGenerator{}), buildCloudfrontDistributionsMock, aws_client.TestOptions{})
}
