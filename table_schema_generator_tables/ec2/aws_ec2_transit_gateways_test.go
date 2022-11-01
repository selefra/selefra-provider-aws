package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEc2TransitGateways(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)
	tgw := types.TransitGateway{}
	err := faker.FakeObject(&tgw)
	if err != nil {
		t.Fatal(err)
	}

	tgwvpca := types.TransitGatewayVpcAttachment{}
	err = faker.FakeObject(&tgwvpca)
	if err != nil {
		t.Fatal(err)
	}

	tgwpeera := types.TransitGatewayPeeringAttachment{}
	err = faker.FakeObject(&tgwpeera)
	if err != nil {
		t.Fatal(err)
	}

	tgwrt := types.TransitGatewayRouteTable{}
	err = faker.FakeObject(&tgwrt)
	if err != nil {
		t.Fatal(err)
	}

	tgwmcd := types.TransitGatewayMulticastDomain{}
	err = faker.FakeObject(&tgwmcd)
	if err != nil {
		t.Fatal(err)
	}

	tgwa := types.TransitGatewayAttachment{}
	err = faker.FakeObject(&tgwa)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTransitGatewayVpcAttachments(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeTransitGatewayVpcAttachmentsOutput{
			TransitGatewayVpcAttachments: []types.TransitGatewayVpcAttachment{tgwvpca},
		}, nil)

	m.EXPECT().DescribeTransitGatewayPeeringAttachments(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeTransitGatewayPeeringAttachmentsOutput{
			TransitGatewayPeeringAttachments: []types.TransitGatewayPeeringAttachment{tgwpeera},
		}, nil)

	m.EXPECT().DescribeTransitGatewayRouteTables(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeTransitGatewayRouteTablesOutput{
			TransitGatewayRouteTables: []types.TransitGatewayRouteTable{tgwrt},
		}, nil)

	m.EXPECT().DescribeTransitGatewayMulticastDomains(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeTransitGatewayMulticastDomainsOutput{
			TransitGatewayMulticastDomains: []types.TransitGatewayMulticastDomain{tgwmcd},
		}, nil)
	m.EXPECT().DescribeTransitGatewayAttachments(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeTransitGatewayAttachmentsOutput{
			TransitGatewayAttachments: []types.TransitGatewayAttachment{tgwa},
		}, nil)
	m.EXPECT().DescribeTransitGateways(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeTransitGatewaysOutput{
			TransitGateways: []types.TransitGateway{tgw},
		}, nil)
	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestEc2TransitGateways(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2TransitGatewaysGenerator{}), buildEc2TransitGateways, aws_client.TestOptions{})
}
