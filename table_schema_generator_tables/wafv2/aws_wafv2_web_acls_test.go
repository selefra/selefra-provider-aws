package wafv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cftypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildWAFV2WebACLMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafV2Client(ctrl)
	cfm := mocks.NewMockCloudfrontClient(ctrl)

	tempWebACLSum := types.WebACLSummary{}
	if err := faker.FakeObject(&tempWebACLSum); err != nil {
		t.Fatal(err)
	}
	var tempResourceArns []string
	if err := faker.FakeObject(&tempResourceArns); err != nil {
		t.Fatal(err)
	}
	var tempTags []types.Tag
	if err := faker.FakeObject(&tempTags); err != nil {
		t.Fatal(err)
	}
	var loggingConfiguration types.LoggingConfiguration
	if err := faker.FakeObject(&loggingConfiguration); err != nil {
		t.Fatal(err)
	}
	rule := types.Rule{}
	if err := faker.FakeObject(&rule); err != nil {
		t.Fatal(err)
	}

	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {
		tempWebACL := types.WebACL{}
		if err := faker.FakeObject(&tempWebACL); err != nil {
			t.Fatal(err)
		}
		m.EXPECT().ListWebACLs(gomock.Any(), &wafv2.ListWebACLsInput{
			Scope:	scope,
			Limit:	aws.Int32(100),
		}, gomock.Any()).AnyTimes().Return(&wafv2.ListWebACLsOutput{
			WebACLs: []types.WebACLSummary{tempWebACLSum},
		}, nil)
		m.EXPECT().GetWebACL(gomock.Any(), &wafv2.GetWebACLInput{
			Id:	tempWebACLSum.Id,
			Name:	tempWebACLSum.Name,
			Scope:	scope,
		}, gomock.Any()).AnyTimes().Return(&wafv2.GetWebACLOutput{
			WebACL: &tempWebACL,
		}, nil)
		m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&wafv2.ListTagsForResourceOutput{
			TagInfoForResource: &types.TagInfoForResource{TagList: tempTags},
		}, nil)
		m.EXPECT().GetLoggingConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&wafv2.GetLoggingConfigurationOutput{
			LoggingConfiguration: &loggingConfiguration,
		}, nil)
	}
	m.EXPECT().ListResourcesForWebACL(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&wafv2.ListResourcesForWebACLOutput{
		ResourceArns: tempResourceArns,
	}, nil)

	distributionList := cftypes.DistributionList{}
	if err := faker.FakeObject(&distributionList); err != nil {
		t.Fatal(err)
	}
	distributionList.NextMarker = nil
	cfm.EXPECT().ListDistributionsByWebACLId(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&cloudfront.ListDistributionsByWebACLIdOutput{
		DistributionList: &distributionList,
	}, nil)

	return aws_client.AwsServices{WafV2: m, Cloudfront: cfm}
}

func TestWafV2WebACL(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafv2WebAclsGenerator{}), buildWAFV2WebACLMock, aws_client.TestOptions{})
}
