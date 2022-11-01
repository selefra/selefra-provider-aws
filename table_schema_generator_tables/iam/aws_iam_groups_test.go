package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildIamGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.Group{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}

	p := iamTypes.AttachedPolicy{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListGroupsOutput{
			Groups: []iamTypes.Group{g},
		}, nil)
	m.EXPECT().ListAttachedGroupPolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListAttachedGroupPoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{p},
		}, nil)

	var l []string
	err = faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListGroupPolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListGroupPoliciesOutput{
			PolicyNames: l,
		}, nil)

	gp := iam.GetGroupPolicyOutput{}
	err = faker.FakeObject(&gp)
	if err != nil {
		t.Fatal(err)
	}
	document := "{\"test\": {\"t1\":1}}"
	gp.PolicyDocument = &document
	m.EXPECT().GetGroupPolicy(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&gp, nil)
	return aws_client.AwsServices{
		IAM: m,
	}
}

func TestIamGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamGroupsGenerator{}), buildIamGroups, aws_client.TestOptions{})
}
