package inspector2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildInspectorV2Findings(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	inspectorClient := mocks.NewMockInspectorV2Client(ctrl)

	finding := types.Finding{}
	err := faker.FakeObject(&finding)
	if err != nil {
		t.Fatal(err)
	}

	inspectorClient.EXPECT().ListFindings(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&inspector2.ListFindingsOutput{Findings: []types.Finding{finding}},
		nil,
	)

	return aws_client.AwsServices{
		InspectorV2: inspectorClient,
	}
}

func TestInspectorV2Findings(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsInspector2FindingsGenerator{}), buildInspectorV2Findings, aws_client.TestOptions{})
}
