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

func buildEc2VpcEndpointServiceConfigurations(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)
	sc := types.ServiceConfiguration{}
	if err := faker.FakeObject(&sc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVpcEndpointServiceConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeVpcEndpointServiceConfigurationsOutput{
			ServiceConfigurations: []types.ServiceConfiguration{sc},
		}, nil)
	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestEc2VpcEndpointServiceConfigurations(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2VpcEndpointServiceConfigurationsGenerator{}), buildEc2VpcEndpointServiceConfigurations, aws_client.TestOptions{})
}
