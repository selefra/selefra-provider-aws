package inspector

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildInspectorFindings(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	inspectorClient := mocks.NewMockInspectorClient(ctrl)

	finding := types.Finding{}
	err := faker.FakeObject(&finding)
	if err != nil {
		t.Fatal(err)
	}

	inspectorClient.EXPECT().ListFindings(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&inspector.ListFindingsOutput{FindingArns: []string{aws.ToString(finding.Arn)}},
		nil,
	)
	inspectorClient.EXPECT().DescribeFindings(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&inspector.DescribeFindingsOutput{Findings: []types.Finding{finding}},
		nil,
	)

	return aws_client.AwsServices{
		Inspector: inspectorClient,
	}
}

func TestInspectorFindings(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsInspectorFindingsGenerator{}), buildInspectorFindings, aws_client.TestOptions{})
}
