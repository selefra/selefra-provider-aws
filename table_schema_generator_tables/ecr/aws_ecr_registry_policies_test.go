package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEcrRegistryPoliciesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEcrClient(ctrl)
	var registryId string
	err := faker.FakeObject(&registryId)
	if err != nil {
		t.Fatal(err)
	}
	policyText := `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Action": [
					"ecr:GetAuthorizationToken",
					"ecr:BatchCheckLayerAvailability",
					"ecr:GetDownloadUrlForLayer",
					"ecr:GetRepositoryPolicy",
					"ecr:DescribeRepositories",
					"ecr:ListImages",
					"ecr:DescribeImages",
					"ecr:BatchGetImage",
					"ecr:GetLifecyclePolicy",
					"ecr:GetLifecyclePolicyPreview",
					"ecr:ListTagsForResource",
					"ecr:DescribeImageScanFindings"
				],
				"Resource": "*"
			}
		]
	}`
	m.EXPECT().GetRegistryPolicy(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ecr.GetRegistryPolicyOutput{
			PolicyText:	aws.String(policyText),
			RegistryId:	aws.String(registryId),
		}, nil)

	return aws_client.AwsServices{
		ECR: m,
	}
}

func TestEcrRegistryPolicies(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEcrRegistryPoliciesGenerator{}), buildEcrRegistryPoliciesMock, aws_client.TestOptions{})
}
