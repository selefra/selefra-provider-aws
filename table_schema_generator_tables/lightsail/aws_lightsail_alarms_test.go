package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildAlarmsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockLightsailClient(ctrl)

	b := lightsail.GetAlarmsOutput{}
	err := faker.FakeObject(&b)
	if err != nil {
		t.Fatal(err)
	}
	b.NextPageToken = nil
	m.EXPECT().GetAlarms(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&b, nil)

	return aws_client.AwsServices{
		Lightsail: m,
	}
}

func TestAlarms(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLightsailAlarmsGenerator{}), buildAlarmsMock, aws_client.TestOptions{})
}
