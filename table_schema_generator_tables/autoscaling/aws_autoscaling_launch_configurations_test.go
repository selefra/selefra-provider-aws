package autoscaling

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildAutoscalingLaunchConfigurationsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockAutoscalingClient(ctrl)
	services := aws_client.AwsServices{
		Autoscaling: m,
	}
	l := types.LaunchConfiguration{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	autoscalingLaunchConfigurations := &autoscaling.DescribeLaunchConfigurationsOutput{
		LaunchConfigurations: []types.LaunchConfiguration{l},
	}
	m.EXPECT().DescribeLaunchConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(autoscalingLaunchConfigurations, nil)
	return services
}

func TestAutoscalingLaunchConfigurations(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAutoscalingLaunchConfigurationsGenerator{}), buildAutoscalingLaunchConfigurationsMock, aws_client.TestOptions{})
}
