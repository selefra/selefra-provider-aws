package firehose

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/firehose/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildKinesisFirehoses(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	f := mocks.NewMockFirehoseClient(ctrl)

	streams := firehose.ListDeliveryStreamsOutput{}
	err := faker.FakeObject(&streams)
	if err != nil {
		t.Fatal(err)
	}
	streams.HasMoreDeliveryStreams = aws.Bool(false)
	streams.DeliveryStreamNames = []string{"test-stream"}
	f.EXPECT().ListDeliveryStreams(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&streams, nil)

	stream := firehose.DescribeDeliveryStreamOutput{}

	err = faker.FakeObject(&stream)
	if err != nil {
		t.Fatal(err)
	}
	stream.DeliveryStreamDescription.Destinations = []types.DestinationDescription{stream.DeliveryStreamDescription.Destinations[0]}

	f.EXPECT().DescribeDeliveryStream(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).AnyTimes().Return(&stream, nil)

	tags := firehose.ListTagsForDeliveryStreamOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.HasMoreTags = aws.Bool(false)
	f.EXPECT().ListTagsForDeliveryStream(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).AnyTimes().Return(&tags, nil)

	return aws_client.AwsServices{
		Firehose: f,
	}
}

func TestFirehoses(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFirehoseDeliveryStreamsGenerator{}), buildKinesisFirehoses, aws_client.TestOptions{})
}
