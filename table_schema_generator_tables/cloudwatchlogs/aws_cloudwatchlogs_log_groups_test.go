package cloudwatchlogs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildCloudwatchLogsLogGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockCloudwatchLogsClient(ctrl)
	l := types.LogGroup{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLogGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&cloudwatchlogs.DescribeLogGroupsOutput{
			LogGroups: []types.LogGroup{l},
		}, nil)

	tags := &cloudwatchlogs.ListTagsLogGroupOutput{}
	err = faker.FakeObject(tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsLogGroup(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(tags, nil)

	return aws_client.AwsServices{
		CloudwatchLogs: m,
	}
}

func TestCloudwatchlogsLogGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCloudwatchlogsLogGroupsGenerator{}), buildCloudwatchLogsLogGroupsMock, aws_client.TestOptions{})
}
