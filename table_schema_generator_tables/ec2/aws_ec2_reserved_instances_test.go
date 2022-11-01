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

func buildReservedEc2Instances(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.ReservedInstances{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeReservedInstances(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).AnyTimes().Return(
		&ec2.DescribeReservedInstancesOutput{
			ReservedInstances: []types.ReservedInstances{l},
		}, nil)
	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestReservedEc2Instances(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2ReservedInstancesGenerator{}), buildReservedEc2Instances, aws_client.TestOptions{})
}
