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

func buildApigatewayDomainNames(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockApigatewayClient(ctrl)

	dm := types.DomainName{}
	err := faker.FakeObject(&dm)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDomainNames(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apigateway.GetDomainNamesOutput{
			Items: []types.DomainName{dm},
		}, nil)

	bpm := types.BasePathMapping{}
	err = faker.FakeObject(&bpm)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetBasePathMappings(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apigateway.GetBasePathMappingsOutput{
			Items: []types.BasePathMapping{bpm},
		}, nil)

	return aws_client.AwsServices{
		Apigateway: m,
	}
}

func TestDomainNames(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsApigatewayDomainNamesGenerator{}), buildApigatewayDomainNames, aws_client.TestOptions{})
}
