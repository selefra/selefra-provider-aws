package wafregional

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildWebACLsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafRegionalClient(ctrl)

	var acl types.WebACL
	if err := faker.FakeObject(&acl); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListWebACLs(
		gomock.Any(),
		&wafregional.ListWebACLsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&wafregional.ListWebACLsOutput{
			WebACLs: []types.WebACLSummary{{WebACLId: acl.WebACLId}},
		},
		nil,
	)

	m.EXPECT().GetWebACL(
		gomock.Any(),
		&wafregional.GetWebACLInput{WebACLId: acl.WebACLId},
		gomock.Any(),
	).AnyTimes().Return(
		&wafregional.GetWebACLOutput{WebACL: &acl},
		nil,
	)

	m.EXPECT().ListTagsForResource(
		gomock.Any(),
		&wafregional.ListTagsForResourceInput{
			ResourceARN: acl.WebACLArn,
		},
		gomock.Any(),
	).AnyTimes().Return(
		&wafregional.ListTagsForResourceOutput{},
		nil,
	)

	m.EXPECT().ListResourcesForWebACL(
		gomock.Any(),
		&wafregional.ListResourcesForWebACLInput{
			WebACLId: acl.WebACLId,
		},
		gomock.Any(),
	).AnyTimes().Return(
		&wafregional.ListResourcesForWebACLOutput{
			ResourceArns: []string{"arn:aws:cloudfront::123456789012:distribution/EDFDVBD6EXAMPLE"},
		},
		nil,
	)

	return aws_client.AwsServices{WafRegional: m}
}

func TestWebACLs(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafregionalWebAclsGenerator{}), buildWebACLsMock, aws_client.TestOptions{})
}
