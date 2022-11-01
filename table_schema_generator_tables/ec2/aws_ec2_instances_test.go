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

func buildEc2Instances(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.Reservation{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	l.Instances[0].StateTransitionReason = aws.String("User initiated (2021-11-26 11:33:00 GMT)")
	creationDate := "1994-11-05T08:15:30-05:00"
	l.Instances[0].ElasticGpuAssociations[0].ElasticGpuAssociationTime = &creationDate
	nextToken := "test"

	m.EXPECT().DescribeInstances(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).AnyTimes().Return(
		&ec2.DescribeInstancesOutput{
			Reservations:	[]types.Reservation{},
			NextToken:	&nextToken,
		}, nil)
	m.EXPECT().DescribeInstances(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).AnyTimes().Return(
		&ec2.DescribeInstancesOutput{
			Reservations:	[]types.Reservation{l},
			NextToken:	nil,
		}, nil)
	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestEc2Instances(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2InstancesGenerator{}), buildEc2Instances, aws_client.TestOptions{})
}
