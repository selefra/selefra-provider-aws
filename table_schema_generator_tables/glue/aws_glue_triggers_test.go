package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func buildTriggersMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlueClient(ctrl)

	var name string
	require.NoError(t, faker.FakeObject(&name))
	m.EXPECT().ListTriggers(
		gomock.Any(),
		&glue.ListTriggersInput{MaxResults: aws.Int32(200)},
	).AnyTimes().Return(
		&glue.ListTriggersOutput{TriggerNames: []string{name}},
		nil,
	)

	var tr types.Trigger
	require.NoError(t, faker.FakeObject(&tr))
	tr.Name = &name
	m.EXPECT().GetTrigger(
		gomock.Any(),
		&glue.GetTriggerInput{Name: aws.String(name)},
	).AnyTimes().Return(
		&glue.GetTriggerOutput{Trigger: &tr},
		nil,
	)

	m.EXPECT().GetTags(
		gomock.Any(),
		&glue.GetTagsInput{ResourceArn: aws.String("arn:aws:glue:us-east-1:testAccount:trigger/" + name)},
	).AnyTimes().Return(
		&glue.GetTagsOutput{Tags: map[string]string{"key": "value"}},
		nil,
	)

	return aws_client.AwsServices{
		Glue: m,
	}
}

func TestTriggers(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlueTriggersGenerator{}), buildTriggersMock, aws_client.TestOptions{})
}
