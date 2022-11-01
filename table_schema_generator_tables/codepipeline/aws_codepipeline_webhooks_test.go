package codepipeline

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildWebhooks(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockCodePipelineClient(ctrl)

	var webhook types.ListWebhookItem
	if err := faker.FakeObject(&webhook); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListWebhooks(
		gomock.Any(),
		&codepipeline.ListWebhooksInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&codepipeline.ListWebhooksOutput{Webhooks: []types.ListWebhookItem{webhook}},
		nil,
	)

	return aws_client.AwsServices{CodePipeline: mock}
}

func TestCodePipelineWebhooks(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCodepipelineWebhooksGenerator{}), buildWebhooks, aws_client.TestOptions{})
}
