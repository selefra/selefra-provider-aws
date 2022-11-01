package cognito

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildCognitoIdentityPools(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockCognitoIdentityPoolsClient(ctrl)

	var desc types.IdentityPoolShortDescription
	if err := faker.FakeObject(&desc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListIdentityPools(
		gomock.Any(),
		&cognitoidentity.ListIdentityPoolsInput{MaxResults: 60},
		gomock.Any(),
	).AnyTimes().Return(
		&cognitoidentity.ListIdentityPoolsOutput{IdentityPools: []types.IdentityPoolShortDescription{desc}},
		nil,
	)

	var ipo cognitoidentity.DescribeIdentityPoolOutput
	if err := faker.FakeObject(&ipo); err != nil {
		t.Fatal(err)
	}
	ipo.IdentityPoolId = desc.IdentityPoolId
	ipo.IdentityPoolId = desc.IdentityPoolName
	m.EXPECT().DescribeIdentityPool(
		gomock.Any(),
		&cognitoidentity.DescribeIdentityPoolInput{IdentityPoolId: desc.IdentityPoolId},
		gomock.Any(),
	).AnyTimes().Return(&ipo, nil)

	return aws_client.AwsServices{CognitoIdentityPools: m}
}

func TestCognitoIdentityPools(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCognitoIdentityPoolsGenerator{}), buildCognitoIdentityPools, aws_client.TestOptions{})
}
