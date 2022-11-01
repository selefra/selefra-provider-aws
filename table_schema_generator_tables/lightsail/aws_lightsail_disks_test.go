package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildDisks(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockLightsailClient(ctrl)

	var disks lightsail.GetDisksOutput
	if err := faker.FakeObject(&disks); err != nil {
		t.Fatal(err)
	}
	disks.NextPageToken = nil
	mock.EXPECT().GetDisks(
		gomock.Any(),
		&lightsail.GetDisksInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&disks,
		nil,
	)

	var diskSnapshots lightsail.GetDiskSnapshotsOutput
	if err := faker.FakeObject(&diskSnapshots); err != nil {
		t.Fatal(err)
	}
	diskSnapshots.NextPageToken = nil
	mock.EXPECT().GetDiskSnapshots(
		gomock.Any(),
		&lightsail.GetDiskSnapshotsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&diskSnapshots,
		nil,
	)

	return aws_client.AwsServices{Lightsail: mock}
}

func TestLightsailDisks(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLightsailDisksGenerator{}), buildDisks, aws_client.TestOptions{})
}
