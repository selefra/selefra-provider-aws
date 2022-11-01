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

func buildIamPasswordPolicies(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.PasswordPolicy{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetAccountPasswordPolicy(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.GetAccountPasswordPolicyOutput{
			PasswordPolicy: &g,
		}, nil)
	return aws_client.AwsServices{
		IAM: m,
	}
}

func TestIamPasswordPolicies(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamPasswordPoliciesGenerator{}), buildIamPasswordPolicies, aws_client.TestOptions{})
}
