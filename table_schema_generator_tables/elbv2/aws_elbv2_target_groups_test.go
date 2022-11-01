package elbv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	elbv2Types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildElbv2TargetGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockElbV2Client(ctrl)
	l := elbv2Types.TargetGroup{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTargetGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticloadbalancingv2.DescribeTargetGroupsOutput{
			TargetGroups: []elbv2Types.TargetGroup{l},
		}, nil)

	tags := elasticloadbalancingv2.DescribeTagsOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tags, nil)

	th := elasticloadbalancingv2.DescribeTargetHealthOutput{}
	err = faker.FakeObject(&th)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTargetHealth(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&th, nil)
	return aws_client.AwsServices{
		ELBv2: m,
	}
}

func TestElbv2TargetGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElbv2TargetGroupsGenerator{}), buildElbv2TargetGroups, aws_client.TestOptions{})
}
