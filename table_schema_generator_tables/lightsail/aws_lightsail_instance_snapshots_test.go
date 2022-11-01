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

func buildInstanceSnapshots(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockLightsailClient(ctrl)

	var is lightsail.GetInstanceSnapshotsOutput
	if err := faker.FakeObject(&is); err != nil {
		t.Fatal(err)
	}
	is.NextPageToken = nil

	mock.EXPECT().GetInstanceSnapshots(
		gomock.Any(),
		&lightsail.GetInstanceSnapshotsInput{},
		gomock.Any(),
	).AnyTimes().Return(&is, nil)

	return aws_client.AwsServices{Lightsail: mock}
}

func TestLightsailInstanceSnapshots(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLightsailInstanceSnapshotsGenerator{}), buildInstanceSnapshots, aws_client.TestOptions{})
}
