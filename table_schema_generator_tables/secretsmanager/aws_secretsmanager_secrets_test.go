package secretsmanager

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildSecretsmanagerModels(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockSecretsManagerClient(ctrl)

	secret := types.SecretListEntry{}
	if err := faker.FakeObject(&secret); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListSecrets(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&secretsmanager.ListSecretsOutput{SecretList: []types.SecretListEntry{secret}},
		nil,
	)

	dsecret := secretsmanager.DescribeSecretOutput{}
	if err := faker.FakeObject(&dsecret); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeSecret(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&dsecret,
		nil,
	)

	var policy secretsmanager.GetResourcePolicyOutput
	if err := faker.FakeObject(&policy); err != nil {
		t.Fatal(err)
	}
	p := `{"key":"value"}`
	policy.ResourcePolicy = &p
	m.EXPECT().GetResourcePolicy(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&policy,
		nil,
	)

	return aws_client.AwsServices{
		SecretsManager: m,
	}
}

func TestSecretsManagerModels(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSecretsmanagerSecretsGenerator{}), buildSecretsmanagerModels, aws_client.TestOptions{})
}
