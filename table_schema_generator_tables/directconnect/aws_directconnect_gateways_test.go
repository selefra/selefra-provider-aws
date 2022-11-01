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

func buildDirectconnectGatewaysMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.DirectConnectGateway{}
	association := types.DirectConnectGatewayAssociation{}
	attachment := types.DirectConnectGatewayAttachment{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeObject(&association)
	if err != nil {
		t.Fatal(err)
	}
	err = faker.FakeObject(&attachment)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeDirectConnectGateways(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&directconnect.DescribeDirectConnectGatewaysOutput{
			DirectConnectGateways: []types.DirectConnectGateway{l},
		}, nil)
	m.EXPECT().DescribeDirectConnectGatewayAssociations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&directconnect.DescribeDirectConnectGatewayAssociationsOutput{
			DirectConnectGatewayAssociations: []types.DirectConnectGatewayAssociation{association},
		}, nil)
	m.EXPECT().DescribeDirectConnectGatewayAttachments(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&directconnect.DescribeDirectConnectGatewayAttachmentsOutput{
			DirectConnectGatewayAttachments: []types.DirectConnectGatewayAttachment{attachment},
		}, nil)
	return aws_client.AwsServices{
		Directconnect: m,
	}
}

func TestDirectconnectGateways(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDirectconnectGatewaysGenerator{}), buildDirectconnectGatewaysMock, aws_client.TestOptions{})
}
