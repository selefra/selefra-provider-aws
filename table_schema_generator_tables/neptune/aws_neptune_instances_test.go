package neptune

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildNeptuneDBInstances(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockNeptuneClient(ctrl)
	l := types.DBInstance{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBInstances(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&neptune.DescribeDBInstancesOutput{
			DBInstances: []types.DBInstance{l},
		}, nil)
	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: l.DBInstanceArn},
		gomock.Any(),
	).AnyTimes().Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)

	return aws_client.AwsServices{
		Neptune: m,
	}
}

func TestNeptuneInstances(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsNeptuneInstancesGenerator{}), buildNeptuneDBInstances, aws_client.TestOptions{})
}
