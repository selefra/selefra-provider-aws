package kinesis

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

type customKinesisClient struct {
	ConsumerCount	*int32

	KeyId			*string
	OpenShardCount		*int32
	RetentionPeriodHours	*int32
	StreamARN		*string
	StreamCreationTimestamp	*time.Time
	StreamModeDetails	*types.StreamModeDetails
	StreamName		*string
	StreamStatus		types.StreamStatus
}

func buildKinesisStreams(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	k := mocks.NewMockKinesisClient(ctrl)

	streams := kinesis.ListStreamsOutput{}
	err := faker.FakeObject(&streams)
	if err != nil {
		t.Fatal(err)
	}
	streams.HasMoreStreams = aws.Bool(false)
	k.EXPECT().ListStreams(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&streams, nil)

	stream := kinesis.DescribeStreamSummaryOutput{
		StreamDescriptionSummary: &types.StreamDescriptionSummary{
			EnhancedMonitoring: []types.EnhancedMetrics{{
				ShardLevelMetrics: []types.MetricsName{types.MetricsNameAll},
			}}},
	}
	customKinesisClient := customKinesisClient{}
	err = faker.FakeObject(&customKinesisClient)
	if err != nil {
		t.Fatal(err)
	}

	stream.StreamDescriptionSummary.ConsumerCount = customKinesisClient.ConsumerCount
	stream.StreamDescriptionSummary.KeyId = customKinesisClient.KeyId
	stream.StreamDescriptionSummary.OpenShardCount = customKinesisClient.OpenShardCount
	stream.StreamDescriptionSummary.RetentionPeriodHours = customKinesisClient.RetentionPeriodHours
	stream.StreamDescriptionSummary.StreamARN = customKinesisClient.StreamARN
	stream.StreamDescriptionSummary.StreamCreationTimestamp = customKinesisClient.StreamCreationTimestamp
	stream.StreamDescriptionSummary.StreamModeDetails = customKinesisClient.StreamModeDetails
	stream.StreamDescriptionSummary.StreamName = customKinesisClient.StreamName
	stream.StreamDescriptionSummary.StreamStatus = customKinesisClient.StreamStatus
	k.EXPECT().DescribeStreamSummary(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).AnyTimes().Return(&stream, nil)

	tags := kinesis.ListTagsForStreamOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.HasMoreTags = aws.Bool(false)
	k.EXPECT().ListTagsForStream(gomock.Any(), gomock.Any(), gomock.Any()).MinTimes(1).AnyTimes().Return(&tags, nil)

	return aws_client.AwsServices{
		Kinesis: k,
	}
}

func TestStreams(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsKinesisStreamsGenerator{}), buildKinesisStreams, aws_client.TestOptions{})
}
