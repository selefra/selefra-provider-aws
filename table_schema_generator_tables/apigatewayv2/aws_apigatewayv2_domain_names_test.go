package apigatewayv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildApigatewayv2DomainNames(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockApigatewayv2Client(ctrl)

	dn := types.DomainName{}
	err := faker.FakeObject(&dn)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDomainNames(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apigatewayv2.GetDomainNamesOutput{
			Items: []types.DomainName{dn},
		}, nil)

	am := types.ApiMapping{}
	err = faker.FakeObject(&am)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetApiMappings(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apigatewayv2.GetApiMappingsOutput{
			Items: []types.ApiMapping{am},
		}, nil)

	return aws_client.AwsServices{
		Apigatewayv2: m,
	}
}

func TestApigatewayv2DomainNames(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsApigatewayv2DomainNamesGenerator{}), buildApigatewayv2DomainNames, aws_client.TestOptions{})
}
