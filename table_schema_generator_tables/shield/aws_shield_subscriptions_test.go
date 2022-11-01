package shield

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildSubscriptions(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockShieldClient(ctrl)
	subscription := shield.DescribeSubscriptionOutput{}
	err := faker.FakeObject(&subscription)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeSubscription(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&subscription, nil)

	return aws_client.AwsServices{
		Shield: m,
	}
}

func TestSubscriptions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsShieldSubscriptionsGenerator{}), buildSubscriptions, aws_client.TestOptions{})
}
