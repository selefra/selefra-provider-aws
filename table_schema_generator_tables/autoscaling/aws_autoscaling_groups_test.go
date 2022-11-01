package autoscaling

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildAutoscalingGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockAutoscalingClient(ctrl)

	groups := autoscaling.DescribeAutoScalingGroupsOutput{}
	err := faker.FakeObject(&groups)
	if err != nil {
		t.Fatal(err)
	}
	groups.NextToken = nil
	m.EXPECT().DescribeAutoScalingGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&groups, nil)

	configurations := autoscaling.DescribeNotificationConfigurationsOutput{}
	err = faker.FakeObject(&configurations)
	if err != nil {
		t.Fatal(err)
	}
	configurations.NextToken = nil
	configurations.NotificationConfigurations[0].AutoScalingGroupName = groups.AutoScalingGroups[0].AutoScalingGroupName
	m.EXPECT().DescribeNotificationConfigurations(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&configurations, nil)

	loadBalancers := autoscaling.DescribeLoadBalancersOutput{}
	err = faker.FakeObject(&loadBalancers)
	if err != nil {
		t.Fatal(err)
	}
	loadBalancers.NextToken = nil
	m.EXPECT().DescribeLoadBalancers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&loadBalancers, nil)

	loadBalancerTargetGroups := autoscaling.DescribeLoadBalancerTargetGroupsOutput{}
	err = faker.FakeObject(&loadBalancerTargetGroups)
	if err != nil {
		t.Fatal(err)
	}
	loadBalancerTargetGroups.NextToken = nil
	m.EXPECT().DescribeLoadBalancerTargetGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&loadBalancerTargetGroups, nil)

	policies := autoscaling.DescribePoliciesOutput{}
	err = faker.FakeObject(&policies)
	if err != nil {
		t.Fatal(err)
	}
	policies.NextToken = nil
	m.EXPECT().DescribePolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&policies, nil)

	lifecycleHooks := autoscaling.DescribeLifecycleHooksOutput{}
	err = faker.FakeObject(&lifecycleHooks)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLifecycleHooks(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&lifecycleHooks, nil)

	return aws_client.AwsServices{
		Autoscaling: m,
	}
}

func TestAutoscalingGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAutoscalingGroupsGenerator{}), buildAutoscalingGroups, aws_client.TestOptions{})
}
