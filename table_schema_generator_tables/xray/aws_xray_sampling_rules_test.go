package xray

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildSamplingRules(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockXrayClient(ctrl)

	test := "test"

	var rule types.SamplingRuleRecord
	if err := faker.FakeObject(&rule); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().GetSamplingRules(
		gomock.Any(),
		&xray.GetSamplingRulesInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&xray.GetSamplingRulesOutput{
			SamplingRuleRecords: []types.SamplingRuleRecord{
				rule,
			},
		},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&xray.ListTagsForResourceInput{ResourceARN: rule.SamplingRule.RuleARN},
		gomock.Any(),
	).AnyTimes().Return(
		&xray.ListTagsForResourceOutput{
			Tags: []types.Tag{
				{
					Key:	&test,
					Value:	&test,
				},
			},
		},
		nil,
	)

	return aws_client.AwsServices{Xray: mock}
}

func TestXraySamplingRules(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsXraySamplingRulesGenerator{}), buildSamplingRules, aws_client.TestOptions{})
}
