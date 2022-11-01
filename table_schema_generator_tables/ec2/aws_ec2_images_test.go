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

func buildEc2ImagesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEc2Client(ctrl)
	services := aws_client.AwsServices{
		EC2: m,
	}
	g := types.Image{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}

	creationDate := "1994-11-05T08:15:30-05:00"
	g.OwnerId = aws.String("testAccount")
	g.CreationDate = &creationDate
	deprecationTime := "2050-11-05T08:15:30-05:00"
	g.DeprecationTime = &deprecationTime

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ec2.DescribeImagesOutput{
			Images: []types.Image{g},
		}, nil).Times(2)

	return services
}

func TestEc2Images(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEc2ImagesGenerator{}), buildEc2ImagesMock, aws_client.TestOptions{})
}
