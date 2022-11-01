package backup

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildBackupRegionSettingsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockBackupClient(ctrl)

	var settings backup.DescribeRegionSettingsOutput
	if err := faker.FakeObject(&settings); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeRegionSettings(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(
		&settings,
		nil,
	)

	return aws_client.AwsServices{
		Backup: m,
	}
}

func TestRegionSettings(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsBackupRegionSettingsGenerator{}), buildBackupRegionSettingsMock, aws_client.TestOptions{})
}
