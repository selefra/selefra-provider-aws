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

func buildEc2Hosts(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)

	g := types.Host{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeHosts(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeHostsOutput{
			Hosts: []types.Host{g},
		}, nil)

	services := aws_client.AwsServices{
		EC2: m,
	}
	return services
}

func TestEc2Hosts(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2HostsGenerator{}), buildEc2Hosts, aws_client.TestOptions{})
}
