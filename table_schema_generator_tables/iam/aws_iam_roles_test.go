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

func buildRoles(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIamClient(ctrl)
	r := iamTypes.Role{}
	err := faker.FakeObject(&r)
	if err != nil {
		t.Fatal(err)
	}

	p := iamTypes.AttachedPolicy{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}

	document := `{"stuff": 3}`
	r.AssumeRolePolicyDocument = &document

	m.EXPECT().GetRole(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.GetRoleOutput{
			Role: &r,
		}, nil)

	m.EXPECT().ListRoles(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListRolesOutput{
			Roles: []iamTypes.Role{r},
		}, nil)
	m.EXPECT().ListAttachedRolePolicies(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListAttachedRolePoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{p},
		}, nil)

	var l []string
	err = faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRolePolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListRolePoliciesOutput{
			PolicyNames: l,
		}, nil)

	pd := iam.GetRolePolicyOutput{}
	err = faker.FakeObject(&pd)
	if err != nil {
		t.Fatal(err)
	}
	pd.PolicyDocument = &document
	m.EXPECT().GetRolePolicy(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&pd, nil)

	tag := iamTypes.Tag{}
	err = faker.FakeObject(&tag)
	if err != nil {
		t.Fatal(err)
	}

	return aws_client.AwsServices{
		IAM: m,
	}
}

func TestIamRoles(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamRolesGenerator{}), buildRoles, aws_client.TestOptions{})
}
