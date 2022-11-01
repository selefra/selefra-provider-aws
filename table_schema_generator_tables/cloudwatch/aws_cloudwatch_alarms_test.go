package cloudwatch

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildCloudWatchAlarmsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockCloudwatchClient(ctrl)
	services := aws_client.AwsServices{
		Cloudwatch: m,
	}
	a := types.MetricAlarm{}
	err := faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeAlarms(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&cloudwatch.DescribeAlarmsOutput{
			MetricAlarms: []types.MetricAlarm{a},
		}, nil)

	tagResponse := cloudwatch.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tagResponse)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tagResponse, nil)

	return services
}

func TestCloudwatchAlarms(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCloudwatchAlarmsGenerator{}), buildCloudWatchAlarmsMock, aws_client.TestOptions{})
}
