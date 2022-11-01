package sns

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

type Subscription struct {
	Endpoint	*string

	Owner	*string

	Protocol	*string

	SubscriptionArn	*string

	TopicArn	*string

	ConfirmationWasAuthenticated	*bool

	DeliveryPolicy	*string

	EffectiveDeliveryPolicy	*string

	FilterPolicy	*string

	PendingConfirmation	*bool

	RawMessageDelivery	*bool

	RedrivePolicy	*string

	SubscriptionRoleArn	*string

	UnknownFields	map[string]interface{}	`mapstructure:",remain"`
}

func buildSnsSubscriptions(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockSnsClient(ctrl)
	sub := types.Subscription{}
	err := faker.FakeObject(&sub)
	if err != nil {
		t.Fatal(err)
	}

	subTemp := types.Subscription{}
	err = faker.FakeObject(&subTemp)
	if err != nil {
		t.Fatal(err)
	}
	emptySub := types.Subscription{
		SubscriptionArn:	aws.String("PendingConfirmation"),
		Owner:			subTemp.Owner,
		Protocol:		subTemp.Protocol,
		TopicArn:		subTemp.TopicArn,
		Endpoint:		subTemp.Endpoint,
	}

	m.EXPECT().ListSubscriptions(
		gomock.Any(),
		&sns.ListSubscriptionsInput{},
	).AnyTimes().Return(
		&sns.ListSubscriptionsOutput{
			Subscriptions: []types.Subscription{sub, emptySub},
		}, nil)

	m.EXPECT().GetSubscriptionAttributes(
		gomock.Any(),
		&sns.GetSubscriptionAttributesInput{SubscriptionArn: sub.SubscriptionArn},
	).AnyTimes().Return(
		&sns.GetSubscriptionAttributesOutput{Attributes: map[string]string{
			"ConfirmationWasAuthenticated":	"true",
			"DeliveryPolicy":		"{}",
			"EffectiveDeliveryPolicy":	"{}",
			"FilterPolicy":			"{}",
			"PendingConfirmation":		"true",
			"RawMessageDelivery":		"true",
			"RedrivePolicy":		`{"deadLetterTargetArn": "test"}`,
			"SubscriptionRoleArn":		"some",
			"WeirdAndUnexpectedField":	"needs updating",
		}},
		nil,
	)

	return aws_client.AwsServices{
		SNS: m,
	}
}

func TestSnsSubscriptions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSnsSubscriptionsGenerator{}), buildSnsSubscriptions, aws_client.TestOptions{})
}
