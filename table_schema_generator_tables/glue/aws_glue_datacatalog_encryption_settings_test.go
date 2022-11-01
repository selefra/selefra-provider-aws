package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func buildDatacatalogEncryptionSettingsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlueClient(ctrl)

	var s glue.GetDataCatalogEncryptionSettingsOutput
	require.NoError(t, faker.FakeObject(&s))
	m.EXPECT().GetDataCatalogEncryptionSettings(
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(&s, nil)

	return aws_client.AwsServices{
		Glue: m,
	}
}

func TestDatacatalogEncryptionSettings(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlueDatacatalogEncryptionSettingsGenerator{}), buildDatacatalogEncryptionSettingsMock, aws_client.TestOptions{})
}
