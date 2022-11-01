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

func buildAutoscalingSheduledActionMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockAutoscalingClient(ctrl)
	services := aws_client.AwsServices{
		Autoscaling: m,
	}
	l := types.ScheduledUpdateGroupAction{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	autoscalingLaunchConfigurations := &autoscaling.DescribeScheduledActionsOutput{
		ScheduledUpdateGroupActions: []types.ScheduledUpdateGroupAction{l},
	}
	m.EXPECT().DescribeScheduledActions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(autoscalingLaunchConfigurations, nil)
	return services
}

func TestAutoscalingSheduledActions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAutoscalingScheduledActionsGenerator{}), buildAutoscalingSheduledActionMock, aws_client.TestOptions{})
}
