package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEc2RegionalConfig(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)
	m.EXPECT().GetEbsDefaultKmsKeyId(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&ec2.GetEbsDefaultKmsKeyIdOutput{KmsKeyId: aws.String("some/key/id")}, nil)
	m.EXPECT().GetEbsEncryptionByDefault(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&ec2.GetEbsEncryptionByDefaultOutput{EbsEncryptionByDefault: aws.Bool(true)}, nil)

	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestEc2RegionalConfig(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2RegionalConfigGenerator{}), buildEc2RegionalConfig, aws_client.TestOptions{})
}
