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

func buildEc2NetworkInterfaces(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)

	niOutput := ec2.DescribeNetworkInterfacesOutput{}
	err := faker.FakeObject(&niOutput)
	if err != nil {
		t.Fatal(err)
	}
	niOutput.NextToken = nil
	m.EXPECT().DescribeNetworkInterfaces(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&niOutput, nil)
	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestEc2NetworkInterfaces(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2NetworkInterfacesGenerator{}), buildEc2NetworkInterfaces, aws_client.TestOptions{})
}
