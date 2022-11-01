package accessanalyzer

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildAccessAnalyzer(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockAnalyzerClient(ctrl)
	u := types.AnalyzerSummary{}
	if err := faker.FakeObject(&u); err != nil {
		t.Fatal(err)
	}
	f := types.FindingSummary{}
	if err := faker.FakeObject(&f); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListAnalyzers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&accessanalyzer.ListAnalyzersOutput{
			Analyzers: []types.AnalyzerSummary{u},
		}, nil)

	m.EXPECT().ListFindings(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&accessanalyzer.ListFindingsOutput{
			Findings: []types.FindingSummary{f},
		}, nil)

	arch := types.ArchiveRuleSummary{}
	if err := faker.FakeObject(&arch); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListArchiveRules(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&accessanalyzer.ListArchiveRulesOutput{
			ArchiveRules: []types.ArchiveRuleSummary{arch},
		}, nil)

	return aws_client.AwsServices{
		Analyzer: m,
	}
}

func TestAccessAnalyzerAnalyzer(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAccessanalyzerAnalyzersGenerator{}), buildAccessAnalyzer, aws_client.TestOptions{})
}
