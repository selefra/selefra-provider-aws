package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEgressOnlyInternetGateways(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)
	egressOutput := ec2.DescribeEgressOnlyInternetGatewaysOutput{}
	err := faker.FakeObject(&egressOutput)
	if err != nil {
		t.Fatal(err)
	}
	egressOutput.NextToken = nil
	m.EXPECT().DescribeEgressOnlyInternetGateways(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&egressOutput, nil)
	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestEgressOnlyInternetGateways(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2EgressOnlyInternetGatewaysGenerator{}), buildEgressOnlyInternetGateways, aws_client.TestOptions{})
}
