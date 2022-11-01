package sqs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildSQSQueues(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	sqsMock := mocks.NewMockSQSClient(ctrl)

	var queueURL = "https://url1"
	sqsMock.EXPECT().ListQueues(
		gomock.Any(),
		&sqs.ListQueuesInput{},
	).AnyTimes().Return(
		&sqs.ListQueuesOutput{QueueUrls: []string{queueURL}},
		nil,
	)

	sqsMock.EXPECT().GetQueueAttributes(
		gomock.Any(),
		&sqs.GetQueueAttributesInput{QueueUrl: &queueURL, AttributeNames: []types.QueueAttributeName{types.QueueAttributeNameAll}},
	).AnyTimes().Return(
		&sqs.GetQueueAttributesOutput{
			Attributes: map[string]string{
				"Policy":					`{"field1":1}`,
				"VisibilityTimeout":				"3600",
				"MaximumMessageSize":				"1000",
				"MessageRetentionPeriod":			"7200",
				"ApproximateNumberOfMessages":			"5",
				"ApproximateNumberOfMessagesNotVisible":	"10",
				"CreatedTimestamp":				"1633416468",
				"LastModifiedTimestamp":			"1633416468",
				"QueueArn":					"arn:aws:sqs:us-east-1:704956590351:terraform-example-queue",
				"ApproximateNumberOfMessagesDelayed":		"6",
				"DelaySeconds":					"7",
				"ReceiveMessageWaitTimeSeconds":		"8",
				"RedrivePolicy":				`{"field2":2}`,
				"FifoQueue":					"true",
				"ContentBasedDeduplication":			"false",
				"KmsMasterKeyId":				"key",
				"KmsDataKeyReusePeriodSeconds":			"9",
				"SqsManagedSseEnabled":				"true",
				"DeduplicationScope":				"messageGroup",
				"FifoThroughputLimit":				"queue",
				"RedriveAllowPolicy":				`{"field3":3}`,

				"UnexpectedField":	"someValue",
			},
		},
		nil,
	)

	sqsMock.EXPECT().ListQueueTags(
		gomock.Any(),
		&sqs.ListQueueTagsInput{QueueUrl: &queueURL},
	).AnyTimes().Return(
		&sqs.ListQueueTagsOutput{Tags: map[string]string{"tag": "value"}},
		nil,
	)
	return aws_client.AwsServices{SQS: sqsMock}
}

func TestQueues(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSqsQueuesGenerator{}), buildSQSQueues, aws_client.TestOptions{})
}
