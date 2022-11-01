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

func buildWAFRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafClient(ctrl)
	tempRuleGroupSum := types.RuleGroupSummary{}
	if err := faker.FakeObject(&tempRuleGroupSum); err != nil {
		t.Fatal(err)
	}
	tempRuleGroup := types.RuleGroup{}
	if err := faker.FakeObject(&tempRuleGroup); err != nil {
		t.Fatal(err)
	}
	tempRule := types.ActivatedRule{}
	if err := faker.FakeObject(&tempRule); err != nil {
		t.Fatal(err)
	}
	var tempTags []types.Tag
	if err := faker.FakeObject(&tempTags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRuleGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.ListRuleGroupsOutput{
		RuleGroups: []types.RuleGroupSummary{tempRuleGroupSum},
	}, nil)
	m.EXPECT().GetRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.GetRuleGroupOutput{
		RuleGroup: &tempRuleGroup,
	}, nil)
	m.EXPECT().ListActivatedRulesInRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.ListActivatedRulesInRuleGroupOutput{
		ActivatedRules: []types.ActivatedRule{tempRule},
	}, nil)
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.ListTagsForResourceOutput{
		TagInfoForResource: &types.TagInfoForResource{TagList: tempTags},
	}, nil)

	return aws_client.AwsServices{Waf: m}
}

func TestWafRuleGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafRuleGroupsGenerator{}), buildWAFRuleGroupsMock, aws_client.TestOptions{})
}
