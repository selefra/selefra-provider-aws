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

func buildEc2EbsVolumes(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)
	volumesOutput := ec2.DescribeVolumesOutput{}
	err := faker.FakeObject(&volumesOutput)
	if err != nil {
		t.Fatal(err)
	}
	volumesOutput.NextToken = nil
	m.EXPECT().DescribeVolumes(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&volumesOutput, nil)
	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestEc2EbsVolumes(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2EbsVolumesGenerator{}), buildEc2EbsVolumes, aws_client.TestOptions{})
}
