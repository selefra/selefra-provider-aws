package sns

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

type Topic struct {
	DeliveryPolicy	*string

	DisplayName	*string

	Owner	*string

	Policy	*string

	SubscriptionsConfirmed	*int

	SubscriptionsDeleted	*int

	SubscriptionsPending	*int

	Arn	*string	`mapstructure:"TopicArn"`

	EffectiveDeliveryPolicy	*string

	KmsMasterKeyId	*string

	FifoTopic	*bool

	ContentBasedDeduplication	*bool

	UnknownFields	map[string]interface{}	`mapstructure:",remain"`
}

func buildSnsTopics(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockSnsClient(ctrl)
	topic := types.Topic{}
	tag := types.Tag{}
	err := faker.FakeObject(&topic)
	if err != nil {
		t.Fatal(err)
	}
	tagerr := faker.FakeObject(&tag)
	if tagerr != nil {
		t.Fatal(tagerr)
	}

	m.EXPECT().ListTopics(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sns.ListTopicsOutput{
			Topics: []types.Topic{topic},
		}, nil)
	m.EXPECT().GetTopicAttributes(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sns.GetTopicAttributesOutput{
			Attributes: map[string]string{
				"SubscriptionsConfirmed":	"5",
				"SubscriptionsDeleted":		"3",
				"SubscriptionsPending":		"0",
				"FifoTopic":			"false",
				"ContentBasedDeduplication":	"true",
				"DisplayName":			"selefra",
				"KmsMasterKeyId":		"test/key",
				"Owner":			"owner",
				"Policy":			`{"stuff": 3}`,
				"DeliveryPolicy":		`{"stuff": 3}`,
				"EffectiveDeliveryPolicy":	`{"stuff": 3}`,
				"WeirdAndUnexpectedField":	"needs updating",
			},
		}, nil)
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&sns.ListTagsForResourceOutput{
			Tags: []types.Tag{tag},
		}, nil)
	return aws_client.AwsServices{
		SNS: m,
	}
}

func TestSnsTopics(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSnsTopicsGenerator{}), buildSnsTopics, aws_client.TestOptions{})
}
