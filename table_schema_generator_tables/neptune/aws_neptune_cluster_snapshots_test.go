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

func buildNeptuneClientForClusterSnapshots(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockNeptuneClient(ctrl)

	var s types.DBClusterSnapshot
	if err := faker.FakeObject(&s); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBClusterSnapshots(
		gomock.Any(),
		&neptune.DescribeDBClusterSnapshotsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&neptune.DescribeDBClusterSnapshotsOutput{DBClusterSnapshots: []types.DBClusterSnapshot{s}},
		nil,
	)

	var attrs []types.DBClusterSnapshotAttribute
	if err := faker.FakeObject(&attrs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBClusterSnapshotAttributes(
		gomock.Any(),
		&neptune.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
		gomock.Any(),
	).AnyTimes().Return(
		&neptune.DescribeDBClusterSnapshotAttributesOutput{
			DBClusterSnapshotAttributesResult: &types.DBClusterSnapshotAttributesResult{DBClusterSnapshotAttributes: attrs},
		},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: s.DBClusterSnapshotArn},
		gomock.Any(),
	).AnyTimes().Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)

	return aws_client.AwsServices{Neptune: mock}
}

func TestNeptuneDBClusterSnapshots(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsNeptuneClusterSnapshotsGenerator{}), buildNeptuneClientForClusterSnapshots, aws_client.TestOptions{})
}
