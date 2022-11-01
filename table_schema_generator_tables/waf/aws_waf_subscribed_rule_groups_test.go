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

func buildWAFSubscribedRuleGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafClient(ctrl)
	tempSubscrRuleGroupSum := types.SubscribedRuleGroupSummary{}
	if err := faker.FakeObject(&tempSubscrRuleGroupSum); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListSubscribedRuleGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.ListSubscribedRuleGroupsOutput{
		RuleGroups: []types.SubscribedRuleGroupSummary{tempSubscrRuleGroupSum},
	}, nil)

	return aws_client.AwsServices{Waf: m}
}

func TestWafSubscribedRuleGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafSubscribedRuleGroupsGenerator{}), buildWAFSubscribedRuleGroupsMock, aws_client.TestOptions{})
}
