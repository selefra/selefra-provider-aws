package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildRule() (*iot.GetTopicRuleOutput, error) {
	p := types.TopicRule{}
	if err := faker.FakeObject(&p); err != nil {
		return nil, err
	}
	a := types.Action{}
	if err := faker.FakeObject(&a); err != nil {
		return nil, err
	}
	p.Actions = []types.Action{
		a,
	}
	p.ErrorAction = &a
	o := iot.GetTopicRuleOutput{}
	if err := faker.FakeObject(&o); err != nil {
		return nil, err
	}
	o.Rule = &p
	return &o, nil
}

func buildIotTopicRules(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIOTClient(ctrl)

	lp := iot.ListTopicRulesOutput{}
	if err := faker.FakeObject(&lp); err != nil {
		t.Fatal(err)
	}
	lp.NextToken = nil
	m.EXPECT().ListTopicRules(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lp, nil)

	p, err := buildRule()
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetTopicRule(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		p, nil)

	tags := iot.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&tags, nil)

	return aws_client.AwsServices{
		IOT: m,
	}
}

func TestIotTopicRules(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIotTopicRulesGenerator{}), buildIotTopicRules, aws_client.TestOptions{})
}
