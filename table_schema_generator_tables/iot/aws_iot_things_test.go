package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildIotThingsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIOTClient(ctrl)

	thing := types.ThingAttribute{}
	err := faker.FakeObject(&thing)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListThings(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iot.ListThingsOutput{Things: []types.ThingAttribute{thing}}, nil)

	lp := iot.ListThingPrincipalsOutput{}
	err = faker.FakeObject(&lp)
	if err != nil {
		t.Fatal(err)
	}
	lp.NextToken = nil
	m.EXPECT().ListThingPrincipals(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lp, nil)

	return aws_client.AwsServices{
		IOT: m,
	}
}

func TestIotThings(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIotThingsGenerator{}), buildIotThingsMock, aws_client.TestOptions{})
}
