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

func buildWAFWebACLMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockWafClient(ctrl)
	tempWebACLSum := types.WebACLSummary{}
	if err := faker.FakeObject(&tempWebACLSum); err != nil {
		t.Fatal(err)
	}
	tempWebACL := types.WebACL{}
	if err := faker.FakeObject(&tempWebACL); err != nil {
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
	m.EXPECT().ListWebACLs(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.ListWebACLsOutput{
		WebACLs: []types.WebACLSummary{tempWebACLSum},
	}, nil)
	m.EXPECT().GetWebACL(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.GetWebACLOutput{
		WebACL: &tempWebACL,
	}, nil)
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.ListTagsForResourceOutput{
		TagInfoForResource: &types.TagInfoForResource{TagList: tempTags},
	}, nil)
	m.EXPECT().GetLoggingConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&waf.GetLoggingConfigurationOutput{
		LoggingConfiguration: &loggingConfiguration,
	}, nil)

	return aws_client.AwsServices{Waf: m}
}

func TestWafWebACL(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsWafWebAclsGenerator{}), buildWAFWebACLMock, aws_client.TestOptions{})
}
