package wafv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildWAFV2ManagedRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafV2Client(ctrl)
	var tempDescribeManagedRuleGroup wafv2.DescribeManagedRuleGroupOutput
	if err := faker.FakeObject(&tempDescribeManagedRuleGroup); err != nil {
		t.Fatal(err)
	}
	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {
		tempManagedRuleGroupSum := types.ManagedRuleGroupSummary{}
		if err := faker.FakeObject(&tempManagedRuleGroupSum); err != nil {
			t.Fatal(err)
		}
		m.EXPECT().ListAvailableManagedRuleGroups(gomock.Any(), &wafv2.ListAvailableManagedRuleGroupsInput{
			Scope: scope,
		}, gomock.Any()).AnyTimes().Return(&wafv2.ListAvailableManagedRuleGroupsOutput{
			ManagedRuleGroups: []types.ManagedRuleGroupSummary{tempManagedRuleGroupSum},
		}, nil)
		m.EXPECT().DescribeManagedRuleGroup(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tempDescribeManagedRuleGroup, nil)
	}

	return aws_client.AwsServices{WafV2: m}
}

func TestWafV2ManagedRuleGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafv2ManagedRuleGroupsGenerator{}), buildWAFV2ManagedRuleGroupsMock, aws_client.TestOptions{})
}
