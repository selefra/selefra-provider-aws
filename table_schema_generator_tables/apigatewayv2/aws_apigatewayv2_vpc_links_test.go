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

func buildApigatewayv2VpcLinks(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockApigatewayv2Client(ctrl)

	v := types.VpcLink{}
	err := faker.FakeObject(&v)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetVpcLinks(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apigatewayv2.GetVpcLinksOutput{
			Items: []types.VpcLink{v},
		}, nil)

	return aws_client.AwsServices{
		Apigatewayv2: m,
	}
}

func TestApigatewayv2VpcLinks(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsApigatewayv2VpcLinksGenerator{}), buildApigatewayv2VpcLinks, aws_client.TestOptions{})
}
