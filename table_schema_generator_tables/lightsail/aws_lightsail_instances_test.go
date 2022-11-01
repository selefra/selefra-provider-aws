package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildInstances(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockLightsailClient(ctrl)

	var instances []types.Instance
	if err := faker.FakeObject(&instances); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().GetInstances(
		gomock.Any(),
		&lightsail.GetInstancesInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&lightsail.GetInstancesOutput{
			Instances: instances,
		},
		nil,
	)

	var p lightsail.GetInstancePortStatesOutput
	if err := faker.FakeObject(&p); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetInstancePortStates(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(&p, nil)

	var a lightsail.GetInstanceAccessDetailsOutput
	if err := faker.FakeObject(&a); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().GetInstanceAccessDetails(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(&a, nil)

	return aws_client.AwsServices{Lightsail: mock}
}

func TestLightsailInstances(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLightsailInstancesGenerator{}), buildInstances, aws_client.TestOptions{})
}
