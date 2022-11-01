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

func buildIamUsers(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIamClient(ctrl)
	u := iamTypes.User{}
	err := faker.FakeObject(&u)
	if err != nil {
		t.Fatal(err)
	}
	g := iamTypes.Group{}
	err = faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}
	km := iamTypes.AccessKeyMetadata{}
	err = faker.FakeObject(&km)
	if err != nil {
		t.Fatal(err)
	}
	aup := iamTypes.AttachedPolicy{}
	err = faker.FakeObject(&aup)
	if err != nil {
		t.Fatal(err)
	}
	akl := iam.GetAccessKeyLastUsedOutput{}
	err = faker.FakeObject(&akl)
	if err != nil {
		t.Fatal(err)
	}

	var tags []iamTypes.Tag
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListUsers(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListUsersOutput{
			Users: []iamTypes.User{u},
		}, nil)
	m.EXPECT().GetUser(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.GetUserOutput{
			User: &u,
		}, nil)
	m.EXPECT().ListGroupsForUser(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListGroupsForUserOutput{
			Groups: []iamTypes.Group{g},
		}, nil)
	m.EXPECT().ListAccessKeys(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListAccessKeysOutput{
			AccessKeyMetadata: []iamTypes.AccessKeyMetadata{km},
		}, nil)
	m.EXPECT().ListAttachedUserPolicies(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListAttachedUserPoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{aup},
		}, nil)
	m.EXPECT().GetAccessKeyLastUsed(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&akl, nil)

	var l []string
	err = faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUserPolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListUserPoliciesOutput{
			PolicyNames: l,
		}, nil)

	p := iam.GetUserPolicyOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	document := "{\"test\": {\"t1\":1}}"
	p.PolicyDocument = &document
	m.EXPECT().GetUserPolicy(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&p, nil)

	return aws_client.AwsServices{
		IAM: m,
	}
}

func TestIamUsers(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamUsersGenerator{}), buildIamUsers, aws_client.TestOptions{})
}
