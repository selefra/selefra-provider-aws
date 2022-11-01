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

func buildLoadBalancers(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockLightsailClient(ctrl)

	var lb lightsail.GetLoadBalancersOutput
	if err := faker.FakeObject(&lb); err != nil {
		t.Fatal(err)
	}
	lb.NextPageToken = nil

	mock.EXPECT().GetLoadBalancers(
		gomock.Any(),
		&lightsail.GetLoadBalancersInput{},
		gomock.Any(),
	).AnyTimes().Return(&lb, nil)

	var lbc lightsail.GetLoadBalancerTlsCertificatesOutput
	if err := faker.FakeObject(&lbc); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().GetLoadBalancerTlsCertificates(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(&lbc, nil)

	return aws_client.AwsServices{Lightsail: mock}
}

func TestLoadBalancers(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLightsailLoadBalancersGenerator{}), buildLoadBalancers, aws_client.TestOptions{})
}
