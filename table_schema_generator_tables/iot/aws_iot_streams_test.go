package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildIotStreamsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIOTClient(ctrl)

	streams := iot.ListStreamsOutput{}
	err := faker.FakeObject(&streams)
	if err != nil {
		t.Fatal(err)
	}
	streams.NextToken = nil
	m.EXPECT().ListStreams(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&streams, nil)

	streamOutput := iot.DescribeStreamOutput{}
	err = faker.FakeObject(&streamOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeStream(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&streamOutput, nil)

	return aws_client.AwsServices{
		IOT: m,
	}
}

func TestIotStreams(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIotStreamsGenerator{}), buildIotStreamsMock, aws_client.TestOptions{})
}
