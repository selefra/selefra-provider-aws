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

func buildApigatewayUsagePlans(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockApigatewayClient(ctrl)

	u := types.UsagePlan{}
	err := faker.FakeObject(&u)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetUsagePlans(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apigateway.GetUsagePlansOutput{
			Items: []types.UsagePlan{u},
		}, nil)

	uk := types.UsagePlanKey{}
	err = faker.FakeObject(&uk)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetUsagePlanKeys(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apigateway.GetUsagePlanKeysOutput{
			Items: []types.UsagePlanKey{uk},
		}, nil)

	return aws_client.AwsServices{
		Apigateway: m,
	}
}

func TestUsagePlans(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsApigatewayUsagePlansGenerator{}), buildApigatewayUsagePlans, aws_client.TestOptions{})
}
