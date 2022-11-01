package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildRDSClientForClusterSnapshots(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockRdsClient(ctrl)

	var s types.DBClusterSnapshot
	if err := faker.FakeObject(&s); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBClusterSnapshots(
		gomock.Any(),
		&rds.DescribeDBClusterSnapshotsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&rds.DescribeDBClusterSnapshotsOutput{DBClusterSnapshots: []types.DBClusterSnapshot{s}},
		nil,
	)

	var attrs []types.DBClusterSnapshotAttribute
	if err := faker.FakeObject(&attrs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBClusterSnapshotAttributes(
		gomock.Any(),
		&rds.DescribeDBClusterSnapshotAttributesInput{DBClusterSnapshotIdentifier: s.DBClusterSnapshotIdentifier},
		gomock.Any(),
	).AnyTimes().Return(
		&rds.DescribeDBClusterSnapshotAttributesOutput{
			DBClusterSnapshotAttributesResult: &types.DBClusterSnapshotAttributesResult{DBClusterSnapshotAttributes: attrs},
		},
		nil,
	)
	return aws_client.AwsServices{RDS: mock}
}

func TestRDSDBClusterSnapshots(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRdsClusterSnapshotsGenerator{}), buildRDSClientForClusterSnapshots, aws_client.TestOptions{})
}
