package apigateway

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildApigatewayClientCertificates(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockApigatewayClient(ctrl)

	c := types.ClientCertificate{}
	err := faker.FakeObject(&c)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetClientCertificates(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apigateway.GetClientCertificatesOutput{
			Items: []types.ClientCertificate{c},
		}, nil)

	return aws_client.AwsServices{
		Apigateway: m,
	}
}

func TestClientCertificates(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsApigatewayClientCertificatesGenerator{}), buildApigatewayClientCertificates, aws_client.TestOptions{})
}
