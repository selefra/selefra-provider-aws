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

func buildIotPolicies(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIOTClient(ctrl)

	lp := iot.ListPoliciesOutput{}
	err := faker.FakeObject(&lp)
	if err != nil {
		t.Fatal(err)
	}
	lp.NextMarker = nil
	m.EXPECT().ListPolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lp, nil)

	p := iot.GetPolicyOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetPolicy(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
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

func TestIotPolicies(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIotPoliciesGenerator{}), buildIotPolicies, aws_client.TestOptions{})
}
