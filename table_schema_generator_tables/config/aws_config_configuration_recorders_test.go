package config

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildConfigConfigurationRecorders(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockConfigServiceClient(ctrl)
	l := types.ConfigurationRecorder{}
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	sl := types.ConfigurationRecorderStatus{}
	if err := faker.FakeObject(&sl); err != nil {
		t.Fatal(err)
	}
	sl.Name = l.Name
	m.EXPECT().DescribeConfigurationRecorderStatus(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&configservice.DescribeConfigurationRecorderStatusOutput{
			ConfigurationRecordersStatus: []types.ConfigurationRecorderStatus{sl},
		}, nil)
	m.EXPECT().DescribeConfigurationRecorders(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&configservice.DescribeConfigurationRecordersOutput{
			ConfigurationRecorders: []types.ConfigurationRecorder{l},
		}, nil)
	return aws_client.AwsServices{
		ConfigService: m,
	}
}

func TestConfigConfigurationRecorders(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsConfigConfigurationRecordersGenerator{}), buildConfigConfigurationRecorders, aws_client.TestOptions{})
}
