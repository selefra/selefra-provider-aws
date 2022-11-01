package ses

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildSESTemplates(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	sesClient := mocks.NewMockSESClient(ctrl)

	tplMeta := types.EmailTemplateMetadata{}
	err := faker.FakeObject(&tplMeta)
	if err != nil {
		t.Fatal(err)
	}

	tpl := new(types.EmailTemplateContent)
	err = faker.FakeObject(tpl)
	if err != nil {
		t.Fatal(err)
	}

	sesClient.EXPECT().ListEmailTemplates(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sesv2.ListEmailTemplatesOutput{TemplatesMetadata: []types.EmailTemplateMetadata{tplMeta}},
		nil,
	)
	sesClient.EXPECT().GetEmailTemplate(gomock.Any(), &sesv2.GetEmailTemplateInput{TemplateName: tplMeta.TemplateName}, gomock.Any()).AnyTimes().Return(
		&sesv2.GetEmailTemplateOutput{
			TemplateName:		tplMeta.TemplateName,
			TemplateContent:	tpl,
		}, nil,
	)

	return aws_client.AwsServices{
		SES: sesClient,
	}
}

func TestSESTemplates(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSesTemplatesGenerator{}), buildSESTemplates, aws_client.TestOptions{})
}
