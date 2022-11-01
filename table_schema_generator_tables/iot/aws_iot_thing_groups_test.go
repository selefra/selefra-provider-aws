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

func buildIotThingGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIOTClient(ctrl)

	groupsOutput := iot.ListThingGroupsOutput{}
	err := faker.FakeObject(&groupsOutput)
	if err != nil {
		t.Fatal(err)
	}
	groupsOutput.NextToken = nil
	m.EXPECT().ListThingGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&groupsOutput, nil)

	groupOutput := iot.DescribeThingGroupOutput{}
	err = faker.FakeObject(&groupOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeThingGroup(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&groupOutput, nil)

	thingsInThingGroupOutput := iot.ListThingsInThingGroupOutput{}
	err = faker.FakeObject(&thingsInThingGroupOutput)
	if err != nil {
		t.Fatal(err)
	}
	thingsInThingGroupOutput.NextToken = nil
	m.EXPECT().ListThingsInThingGroup(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&thingsInThingGroupOutput, nil)

	p := iot.ListAttachedPoliciesOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	p.NextMarker = nil
	m.EXPECT().ListAttachedPolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&p, nil)

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

func TestIotThingGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIotThingGroupsGenerator{}), buildIotThingGroupsMock, aws_client.TestOptions{})
}
