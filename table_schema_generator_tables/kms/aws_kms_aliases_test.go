package kms

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildKmsAliases(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockKmsClient(ctrl)

	aliases := kms.ListAliasesOutput{}
	err := faker.FakeObject(&aliases)
	if err != nil {
		t.Fatal(err)
	}
	aliases.NextMarker = nil
	m.EXPECT().ListAliases(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&aliases, nil)

	return aws_client.AwsServices{
		KMS: m,
	}
}

func TestKmsAliases(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsKmsAliasesGenerator{}), buildKmsAliases, aws_client.TestOptions{})
}
