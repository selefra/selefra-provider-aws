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

func buildRDSClient(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockRdsClient(ctrl)

	var s types.DBSnapshot
	if err := faker.FakeObject(&s); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBSnapshots(
		gomock.Any(),
		&rds.DescribeDBSnapshotsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&rds.DescribeDBSnapshotsOutput{DBSnapshots: []types.DBSnapshot{s}},
		nil,
	)

	var attrs []types.DBSnapshotAttribute
	if err := faker.FakeObject(&attrs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBSnapshotAttributes(
		gomock.Any(),
		&rds.DescribeDBSnapshotAttributesInput{DBSnapshotIdentifier: s.DBSnapshotIdentifier},
		gomock.Any(),
	).AnyTimes().Return(
		&rds.DescribeDBSnapshotAttributesOutput{
			DBSnapshotAttributesResult: &types.DBSnapshotAttributesResult{DBSnapshotAttributes: attrs},
		},
		nil,
	)
	return aws_client.AwsServices{RDS: mock}
}

func TestRDSDBSnapshots(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRdsDbSnapshotsGenerator{}), buildRDSClient, aws_client.TestOptions{})
}
