package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildConfigConformancePack(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockConfigServiceClient(ctrl)

	var cpd types.ConformancePackDetail
	if err := faker.FakeObject(&cpd); err != nil {
		t.Fatal(err)
	}
	var cprc types.ConformancePackRuleCompliance
	if err := faker.FakeObject(&cprc); err != nil {
		t.Fatal(err)
	}
	var cpre types.ConformancePackEvaluationResult
	if err := faker.FakeObject(&cpre); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeConformancePacks(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&configservice.DescribeConformancePacksOutput{
			ConformancePackDetails: []types.ConformancePackDetail{cpd},
		}, nil)
	m.EXPECT().DescribeConformancePackCompliance(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&configservice.DescribeConformancePackComplianceOutput{
			ConformancePackRuleComplianceList: []types.ConformancePackRuleCompliance{cprc},
		}, nil)
	m.EXPECT().GetConformancePackComplianceDetails(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&configservice.GetConformancePackComplianceDetailsOutput{
			ConformancePackRuleEvaluationResults: []types.ConformancePackEvaluationResult{cpre},
		}, nil)

	return aws_client.AwsServices{
		ConfigService: m,
	}
}

func TestConfigConformancePack(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsConfigConformancePacksGenerator{}), buildConfigConformancePack, aws_client.TestOptions{})
}
