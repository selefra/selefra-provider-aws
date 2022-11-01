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

func buildIotBillingGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIOTClient(ctrl)

	groupsOutput := iot.ListBillingGroupsOutput{}
	err := faker.FakeObject(&groupsOutput)
	if err != nil {
		t.Fatal(err)
	}
	groupsOutput.NextToken = nil
	m.EXPECT().ListBillingGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&groupsOutput, nil)

	groupOutput := iot.DescribeBillingGroupOutput{}
	err = faker.FakeObject(&groupOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeBillingGroup(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&groupOutput, nil)

	thingsInBillingGroupOutput := iot.ListThingsInBillingGroupOutput{}
	err = faker.FakeObject(&thingsInBillingGroupOutput)
	if err != nil {
		t.Fatal(err)
	}
	thingsInBillingGroupOutput.NextToken = nil
	m.EXPECT().ListThingsInBillingGroup(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&thingsInBillingGroupOutput, nil)

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

func TestIotBillingGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIotBillingGroupsGenerator{}), buildIotBillingGroupsMock, aws_client.TestOptions{})
}
