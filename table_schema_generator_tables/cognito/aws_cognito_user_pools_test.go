package cognito

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildCognitoUserPools(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockCognitoUserPoolsClient(ctrl)

	var desc types.UserPoolDescriptionType
	if err := faker.FakeObject(&desc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUserPools(
		gomock.Any(),
		&cognitoidentityprovider.ListUserPoolsInput{MaxResults: 60},
		gomock.Any(),
	).AnyTimes().Return(
		&cognitoidentityprovider.ListUserPoolsOutput{UserPools: []types.UserPoolDescriptionType{desc}},
		nil,
	)

	var pool types.UserPoolType
	if err := faker.FakeObject(&pool); err != nil {
		t.Fatal(err)
	}
	pool.Id = desc.Id
	m.EXPECT().DescribeUserPool(
		gomock.Any(),
		&cognitoidentityprovider.DescribeUserPoolInput{UserPoolId: desc.Id},
		gomock.Any(),
	).AnyTimes().Return(
		&cognitoidentityprovider.DescribeUserPoolOutput{UserPool: &pool},
		nil,
	)

	var providerDesc types.ProviderDescription
	if err := faker.FakeObject(&providerDesc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListIdentityProviders(
		gomock.Any(),
		&cognitoidentityprovider.ListIdentityProvidersInput{UserPoolId: pool.Id},
		gomock.Any(),
	).AnyTimes().Return(
		&cognitoidentityprovider.ListIdentityProvidersOutput{Providers: []types.ProviderDescription{providerDesc}},
		nil,
	)

	var provider types.IdentityProviderType
	if err := faker.FakeObject(&provider); err != nil {
		t.Fatal(err)
	}
	provider.ProviderName = providerDesc.ProviderName
	provider.UserPoolId = pool.Id
	m.EXPECT().DescribeIdentityProvider(
		gomock.Any(),
		&cognitoidentityprovider.DescribeIdentityProviderInput{
			ProviderName:	providerDesc.ProviderName,
			UserPoolId:	pool.Id,
		},
		gomock.Any(),
	).AnyTimes().Return(
		&cognitoidentityprovider.DescribeIdentityProviderOutput{IdentityProvider: &provider},
		nil,
	)

	return aws_client.AwsServices{CognitoUserPools: m}
}

func TestCognitoUserPools(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCognitoUserPoolsGenerator{}), buildCognitoUserPools, aws_client.TestOptions{})
}
