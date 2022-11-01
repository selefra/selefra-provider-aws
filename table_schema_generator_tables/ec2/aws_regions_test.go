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

func buildRegionsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)
	r := types.Region{}
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	r.OptInStatus = aws.String("opted-in")
	r.RegionName = aws.String("us-east-1")
	m.EXPECT().DescribeRegions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeRegionsOutput{
			Regions: []types.Region{r},
		}, nil)

	return aws_client.AwsServices{
		EC2: m,
	}
}

func TestRegions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRegionsGenerator{}), buildRegionsMock, aws_client.TestOptions{})
}
