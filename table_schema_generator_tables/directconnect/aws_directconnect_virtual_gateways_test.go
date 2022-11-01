package directconnect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildDirectconnectVirtualGatewaysMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.VirtualGateway{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVirtualGateways(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&directconnect.DescribeVirtualGatewaysOutput{
			VirtualGateways: []types.VirtualGateway{l},
		}, nil)
	return aws_client.AwsServices{
		Directconnect: m,
	}
}

func TestDirectconnecVirtualGateways(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDirectconnectVirtualGatewaysGenerator{}), buildDirectconnectVirtualGatewaysMock, aws_client.TestOptions{})
}
