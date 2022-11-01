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

func buildCertificatesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockLightsailClient(ctrl)

	b := lightsail.GetCertificatesOutput{}
	err := faker.FakeObject(&b)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetCertificates(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&b, nil)

	return aws_client.AwsServices{
		Lightsail: m,
	}
}

func TestCertificates(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLightsailCertificatesGenerator{}), buildCertificatesMock, aws_client.TestOptions{})
}
