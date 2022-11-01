package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildIotSecurityProfilesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIOTClient(ctrl)

	sp := iot.ListSecurityProfilesOutput{}
	err := faker.FakeObject(&sp)
	if err != nil {
		t.Fatal(err)
	}
	sp.NextToken = nil
	m.EXPECT().ListSecurityProfiles(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sp, nil)

	profileOutput := iot.DescribeSecurityProfileOutput{}
	err = faker.FakeObject(&profileOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeSecurityProfile(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&profileOutput, nil)

	targets := iot.ListTargetsForSecurityProfileOutput{}
	err = faker.FakeObject(&targets)
	if err != nil {
		t.Fatal(err)
	}
	targets.NextToken = nil

	m.EXPECT().ListTargetsForSecurityProfile(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&targets, nil)

	tags := iot.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&tags, nil)

	return aws_client.AwsServices{
		IOT: m,
	}
}

func TestIotSecurityProfiles(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIotSecurityProfilesGenerator{}), buildIotSecurityProfilesMock, aws_client.TestOptions{})
}
