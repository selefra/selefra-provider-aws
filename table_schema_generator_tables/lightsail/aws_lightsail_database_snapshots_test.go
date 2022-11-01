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

func buildDatabaseSnapshotsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockLightsailClient(ctrl)

	s := lightsail.GetRelationalDatabaseSnapshotsOutput{}
	err := faker.FakeObject(&s)
	if err != nil {
		t.Fatal(err)
	}
	s.NextPageToken = nil
	m.EXPECT().GetRelationalDatabaseSnapshots(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&s, nil)

	return aws_client.AwsServices{
		Lightsail: m,
	}
}

func TestDatabaseSnapshots(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLightsailDatabaseSnapshotsGenerator{}), buildDatabaseSnapshotsMock, aws_client.TestOptions{})
}
