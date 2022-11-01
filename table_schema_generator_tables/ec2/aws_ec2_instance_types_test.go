package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEc2InstanceTypes(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)
	info := types.InstanceTypeInfo{}
	err := faker.FakeObject(&info)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeInstanceTypes(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeInstanceTypesOutput{
			InstanceTypes:	[]types.InstanceTypeInfo{info},
			NextToken:	nil,
			ResultMetadata:	middleware.Metadata{},
		}, nil)
	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestEc2InstanceTypes(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2InstanceTypesGenerator{}), buildEc2InstanceTypes, aws_client.TestOptions{})
}
