package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildDistributions(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockLightsailClient(ctrl)

	var d lightsail.GetDistributionsOutput
	if err := faker.FakeObject(&d); err != nil {
		t.Fatal(err)
	}
	d.NextPageToken = nil
	mock.EXPECT().GetDistributions(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(&d, nil)

	var r lightsail.GetDistributionLatestCacheResetOutput
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetDistributionLatestCacheReset(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(&r, nil)

	return aws_client.AwsServices{Lightsail: mock}
}

func TestLightsailDistributions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLightsailDistributionsGenerator{}), buildDistributions, aws_client.TestOptions{})
}
