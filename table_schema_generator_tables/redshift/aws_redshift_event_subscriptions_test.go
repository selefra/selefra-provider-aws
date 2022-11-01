package redshift

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildRedshiftEventSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockRedshiftClient(ctrl)

	var s types.EventSubscription
	if err := faker.FakeObject(&s); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeEventSubscriptions(
		gomock.Any(),
		&redshift.DescribeEventSubscriptionsInput{MaxRecords: aws.Int32(100)},
		gomock.Any(),
	).AnyTimes().Return(
		&redshift.DescribeEventSubscriptionsOutput{
			EventSubscriptionsList: []types.EventSubscription{s},
		},
		nil,
	)

	return aws_client.AwsServices{
		Redshift: m,
	}
}

func TestRedshiftEventSubscriptions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRedshiftEventSubscriptionsGenerator{}), buildRedshiftEventSubscriptionsMock, aws_client.TestOptions{})
}
