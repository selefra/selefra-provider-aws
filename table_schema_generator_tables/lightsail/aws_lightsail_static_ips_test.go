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

func buildStaticIps(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockLightsailClient(ctrl)

	var ips lightsail.GetStaticIpsOutput
	if err := faker.FakeObject(&ips); err != nil {
		t.Fatal(err)
	}
	ips.NextPageToken = nil

	mock.EXPECT().GetStaticIps(gomock.Any(), &lightsail.GetStaticIpsInput{}, gomock.Any()).AnyTimes().Return(&ips, nil)

	return aws_client.AwsServices{Lightsail: mock}
}

func TestStaticIps(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLightsailStaticIpsGenerator{}), buildStaticIps, aws_client.TestOptions{})
}
