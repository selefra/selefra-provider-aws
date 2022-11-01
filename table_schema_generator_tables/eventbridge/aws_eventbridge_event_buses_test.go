package eventbridge

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEventBridgeEventBusesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEventBridgeClient(ctrl)
	bus := types.EventBus{}
	err := faker.FakeObject(&bus)
	if err != nil {
		t.Fatal(err)
	}

	rule := types.Rule{}
	err = faker.FakeObject(&rule)
	if err != nil {
		t.Fatal(err)
	}

	tags := eventbridge.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListEventBuses(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&eventbridge.ListEventBusesOutput{
			EventBuses: []types.EventBus{bus},
		}, nil)
	m.EXPECT().ListRules(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&eventbridge.ListRulesOutput{
			Rules: []types.Rule{rule},
		}, nil)
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(2).AnyTimes().Return(
		&tags, nil)

	return aws_client.AwsServices{
		EventBridge: m,
	}
}

func TestEventBridgeEventBuses(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEventbridgeEventBusesGenerator{}), buildEventBridgeEventBusesMock, aws_client.TestOptions{})
}
