package waf

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildWAFRulesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafClient(ctrl)
	tempRuleSum := types.RuleSummary{}
	if err := faker.FakeObject(&tempRuleSum); err != nil {
		t.Fatal(err)
	}
	tempRule := types.Rule{}
	if err := faker.FakeObject(&tempRule); err != nil {
		t.Fatal(err)
	}
	var tempTags []types.Tag
	if err := faker.FakeObject(&tempTags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRules(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.ListRulesOutput{
		Rules: []types.RuleSummary{tempRuleSum},
	}, nil)
	m.EXPECT().GetRule(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.GetRuleOutput{
		Rule: &tempRule,
	}, nil)
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.ListTagsForResourceOutput{
		TagInfoForResource: &types.TagInfoForResource{TagList: tempTags},
	}, nil)

	return aws_client.AwsServices{Waf: m}
}

func TestWafRules(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafRulesGenerator{}), buildWAFRulesMock, aws_client.TestOptions{})
}
