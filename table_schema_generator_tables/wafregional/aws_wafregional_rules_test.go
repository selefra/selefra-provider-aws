package wafregional

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildRulesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafRegionalClient(ctrl)

	var r types.Rule
	if err := faker.FakeObject(&r); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRules(
		gomock.Any(),
		&wafregional.ListRulesInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&wafregional.ListRulesOutput{
			Rules: []types.RuleSummary{{RuleId: r.RuleId}},
		},
		nil,
	)

	m.EXPECT().GetRule(
		gomock.Any(),
		&wafregional.GetRuleInput{RuleId: r.RuleId},
		gomock.Any(),
	).AnyTimes().Return(
		&wafregional.GetRuleOutput{Rule: &r},
		nil,
	)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&wafregional.ListTagsForResourceInput{
			ResourceARN: aws.String(fmt.Sprintf("arn:aws:waf-regional:us-east-1:testAccount:rule/%v", *r.RuleId)),
		},
		gomock.Any(),
	).AnyTimes().Return(
		&wafregional.ListTagsForResourceOutput{},
		nil,
	)

	return aws_client.AwsServices{WafRegional: m}
}

func TestRules(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafregionalRulesGenerator{}), buildRulesMock, aws_client.TestOptions{})
}
