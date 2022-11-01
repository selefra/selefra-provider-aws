package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEc2NetworkAcls(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)

	l := types.NetworkAcl{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	l.IsDefault = aws.Bool(false)
	m.EXPECT().DescribeNetworkAcls(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeNetworkAclsOutput{
			NetworkAcls: []types.NetworkAcl{l},
		}, nil)
	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestEc2NetworkAclsMockTest(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2NetworkAclsGenerator{}), buildEc2NetworkAcls, aws_client.TestOptions{})
}
