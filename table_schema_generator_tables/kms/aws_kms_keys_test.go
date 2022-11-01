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

func buildKmsKeys(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockKmsClient(ctrl)

	keys := kms.ListKeysOutput{}
	err := faker.FakeObject(&keys)
	if err != nil {
		t.Fatal(err)
	}
	keys.NextMarker = nil
	m.EXPECT().ListKeys(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&keys, nil)

	tags := kms.ListResourceTagsOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextMarker = nil
	m.EXPECT().ListResourceTags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&tags, nil)

	key := kms.DescribeKeyOutput{}
	err = faker.FakeObject(&key)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeKey(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&key, nil)

	rotation := kms.GetKeyRotationStatusOutput{}
	err = faker.FakeObject(&rotation)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetKeyRotationStatus(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&rotation, nil)

	return aws_client.AwsServices{
		KMS: m,
	}
}

func TestKmsKeys(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsKmsKeysGenerator{}), buildKmsKeys, aws_client.TestOptions{})
}
