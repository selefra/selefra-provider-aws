package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildSnapshotsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockFsxClient(ctrl)

	var s types.Snapshot
	err := faker.FakeObject(&s)
	if err != nil {
		t.Fatalf("FakeObject returned error: %v", err)
	}
	s.Lifecycle = types.SnapshotLifecycleAvailable
	m.EXPECT().DescribeSnapshots(
		gomock.Any(),
		&fsx.DescribeSnapshotsInput{MaxResults: aws.Int32(1000)},
	).AnyTimes().Return(
		&fsx.DescribeSnapshotsOutput{Snapshots: []types.Snapshot{s}},
		nil,
	)

	return aws_client.AwsServices{
		FSX: m,
	}
}

func TestSnapshots(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFsxSnapshotsGenerator{}), buildSnapshotsMock, aws_client.TestOptions{})
}
