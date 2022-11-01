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

func buildIamPolicies(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.ManagedPolicyDetail{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}
	document := `{"stuff": 3}`

	for i := range g.PolicyVersionList {
		g.PolicyVersionList[i].Document = &document
	}

	m.EXPECT().GetAccountAuthorizationDetails(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.GetAccountAuthorizationDetailsOutput{
			Policies: []iamTypes.ManagedPolicyDetail{g},
		}, nil)

	tag := iamTypes.Tag{}
	err = faker.FakeObject(&tag)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListPolicyTags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListPolicyTagsOutput{
			Tags: []iamTypes.Tag{
				tag,
			},
		}, nil)
	return aws_client.AwsServices{
		IAM: m,
	}
}

func TestIamPolicies(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamPoliciesGenerator{}), buildIamPolicies, aws_client.TestOptions{})
}
