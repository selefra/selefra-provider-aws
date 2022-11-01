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

func buildNeptuneDBClusters(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockNeptuneClient(ctrl)
	l := types.DBCluster{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBClusters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&neptune.DescribeDBClustersOutput{
			DBClusters: []types.DBCluster{l},
		}, nil)
	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: l.DBClusterArn},
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

func TestNeptuneClusters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsNeptuneClustersGenerator{}), buildNeptuneDBClusters, aws_client.TestOptions{})
}
