package xray

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEncryptionConfig(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockXrayClient(ctrl)

	var config types.EncryptionConfig
	if err := faker.FakeObject(&config); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().GetEncryptionConfig(
		gomock.Any(),
		&xray.GetEncryptionConfigInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&xray.GetEncryptionConfigOutput{
			EncryptionConfig: &config,
		},
		nil,
	)

	return aws_client.AwsServices{Xray: mock}
}

func TestXrayEncryptionConfig(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsXrayEncryptionConfigGenerator{}), buildEncryptionConfig, aws_client.TestOptions{})
}
