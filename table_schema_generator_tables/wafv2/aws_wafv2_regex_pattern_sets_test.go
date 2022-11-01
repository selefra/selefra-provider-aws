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

func buildRegexPatternSetsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafV2Client(ctrl)

	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {
		var s types.RegexPatternSet
		if err := faker.FakeObject(&s); err != nil {
			t.Fatal(err)
		}
		m.EXPECT().ListRegexPatternSets(
			gomock.Any(),
			&wafv2.ListRegexPatternSetsInput{Scope: scope, Limit: aws.Int32(100)},
			gomock.Any(),
		).AnyTimes().Return(
			&wafv2.ListRegexPatternSetsOutput{
				RegexPatternSets: []types.RegexPatternSetSummary{{Id: s.Id, Name: s.Name}},
			},
			nil,
		)

		m.EXPECT().GetRegexPatternSet(
			gomock.Any(),
			&wafv2.GetRegexPatternSetInput{Id: s.Id, Name: s.Name, Scope: scope},
			gomock.Any(),
		).AnyTimes().Return(
			&wafv2.GetRegexPatternSetOutput{RegexPatternSet: &s},
			nil,
		)

		m.EXPECT().ListTagsForResource(
			gomock.Any(),
			&wafv2.ListTagsForResourceInput{ResourceARN: s.ARN},
			gomock.Any(),
		).AnyTimes().Return(
			&wafv2.ListTagsForResourceOutput{
				TagInfoForResource: &types.TagInfoForResource{
					TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
				},
			},
			nil,
		)
	}

	return aws_client.AwsServices{WafV2: m}
}

func TestWafV2RegexPatternSets(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafv2RegexPatternSetsGenerator{}), buildRegexPatternSetsMock, aws_client.TestOptions{})
}
