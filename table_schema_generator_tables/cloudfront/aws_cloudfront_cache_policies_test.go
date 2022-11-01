package cloudfront

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildCloudfrontCachePoliciesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := aws_client.AwsServices{
		Cloudfront: m,
	}
	cp := cloudfrontTypes.CachePolicySummary{}
	if err := faker.FakeObject(&cp); err != nil {
		t.Fatal(err)
	}

	cloudfrontOutput := &cloudfront.ListCachePoliciesOutput{
		CachePolicyList: &cloudfrontTypes.CachePolicyList{
			Items: []cloudfrontTypes.CachePolicySummary{cp},
		},
	}
	m.EXPECT().ListCachePolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		cloudfrontOutput,
		nil,
	)
	return services
}

func TestCloudfrontCachePolicies(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCloudfrontCachePoliciesGenerator{}), buildCloudfrontCachePoliciesMock, aws_client.TestOptions{})
}
