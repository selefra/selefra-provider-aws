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

func buildEc2CustomerGateways(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.CustomerGateway{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeCustomerGateways(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeCustomerGatewaysOutput{
			CustomerGateways: []types.CustomerGateway{l},
		}, nil)
	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestEc2CustomerGateways(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2CustomerGatewaysGenerator{}), buildEc2CustomerGateways, aws_client.TestOptions{})
}
