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

func buildNeptuneGlobalClusters(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockNeptuneClient(ctrl)
	var gc types.GlobalCluster
	if err := faker.FakeObject(&gc); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeGlobalClusters(gomock.Any(), &neptune.DescribeGlobalClustersInput{}, gomock.Any()).AnyTimes().Return(
		&neptune.DescribeGlobalClustersOutput{GlobalClusters: []types.GlobalCluster{gc}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: gc.GlobalClusterArn},
		gomock.Any(),
	).AnyTimes().Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)
	return aws_client.AwsServices{Neptune: mock}
}

func TestNeptuneGlobalCluster(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsNeptuneGlobalClustersGenerator{}), buildNeptuneGlobalClusters, aws_client.TestOptions{})
}
