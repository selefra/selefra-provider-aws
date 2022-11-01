package wafv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildWAFV2RuleGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafV2Client(ctrl)
	visibilityConfig := types.VisibilityConfig{}
	if err := faker.FakeObject(&visibilityConfig); err != nil {
		t.Fatal(err)
	}
	customRespBody := map[string]types.CustomResponseBody{}
	if err := faker.FakeObject(&customRespBody); err != nil {
		t.Fatal(err)
	}
	var labelSummaries []types.LabelSummary
	if err := faker.FakeObject(&labelSummaries); err != nil {
		t.Fatal(err)
	}
	overrideAction := types.OverrideAction{}
	if err := faker.FakeObject(&overrideAction); err != nil {
		t.Fatal(err)
	}
	action := types.RuleAction{}
	if err := faker.FakeObject(&action); err != nil {
		t.Fatal(err)
	}
	var labels []types.Label
	if err := faker.FakeObject(&labelSummaries); err != nil {
		t.Fatal(err)
	}
	rule := types.Rule{}
	if err := faker.FakeObject(&rule); err != nil {
		t.Fatal(err)
	}
	rule.VisibilityConfig = &visibilityConfig
	rule.Action = &action
	rule.OverrideAction = &overrideAction
	rule.RuleLabels = labels
	var tempPolicyOutput wafv2.GetPermissionPolicyOutput
	if err := faker.FakeObject(&tempPolicyOutput); err != nil {
		t.Fatal(err)
	}
	tempPolicyOutput.Policy = aws.String(`{"test": 1}`)
	var tempTags []types.Tag
	if err := faker.FakeObject(&tempTags); err != nil {
		t.Fatal(err)
	}
	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {
		tempRuleGroupSum := types.RuleGroupSummary{}
		if err := faker.FakeObject(&tempRuleGroupSum); err != nil {
			t.Fatal(err)
		}
		m.EXPECT().ListRuleGroups(gomock.Any(), &wafv2.ListRuleGroupsInput{Scope: scope}, gomock.Any()).AnyTimes().Return(&wafv2.ListRuleGroupsOutput{
			RuleGroups: []types.RuleGroupSummary{tempRuleGroupSum},
		}, nil)

		tempRuleGroup := types.RuleGroup{}
		if err := faker.FakeObject(&tempRuleGroup); err != nil {
			t.Fatal(err)
		}
		m.EXPECT().GetRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&wafv2.GetRuleGroupOutput{
			RuleGroup: &tempRuleGroup,
		}, nil)
		m.EXPECT().GetPermissionPolicy(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tempPolicyOutput, nil)
		m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&wafv2.ListTagsForResourceOutput{
			TagInfoForResource: &types.TagInfoForResource{TagList: tempTags},
		}, nil)
	}

	return aws_client.AwsServices{WafV2: m}
}

func TestWafV2RuleGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafv2RuleGroupsGenerator{}), buildWAFV2RuleGroupsMock, aws_client.TestOptions{})
}
