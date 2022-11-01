package elbv1

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	elbv1Types "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildElbv1LoadBalancers(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockElbV1Client(ctrl)
	l := elbv1Types.LoadBalancerDescription{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeLoadBalancers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticloadbalancing.DescribeLoadBalancersOutput{
			LoadBalancerDescriptions: []elbv1Types.LoadBalancerDescription{l},
		}, nil)

	tag := elbv1Types.Tag{}
	err = faker.FakeObject(&tag)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticloadbalancing.DescribeTagsOutput{
			TagDescriptions: []elbv1Types.TagDescription{
				{
					LoadBalancerName:	l.LoadBalancerName,
					Tags:			[]elbv1Types.Tag{tag},
				},
			},
		}, nil)

	a := elbv1Types.LoadBalancerAttributes{}
	err = faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLoadBalancerAttributes(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticloadbalancing.DescribeLoadBalancerAttributesOutput{
			LoadBalancerAttributes: &a,
		}, nil)

	p := elbv1Types.PolicyDescription{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeLoadBalancerPolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticloadbalancing.DescribeLoadBalancerPoliciesOutput{
			PolicyDescriptions: []elbv1Types.PolicyDescription{p},
		}, nil)

	return aws_client.AwsServices{
		ELBv1: m,
	}
}

func TestElbv1LoadBalancers(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElbv1LoadBalancersGenerator{}), buildElbv1LoadBalancers, aws_client.TestOptions{})
}
