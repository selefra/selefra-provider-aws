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

func buildIamOpenIDConnectProviders(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIamClient(ctrl)
	l := iamTypes.OpenIDConnectProviderListEntry{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListOpenIDConnectProviders(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListOpenIDConnectProvidersOutput{
			OpenIDConnectProviderList: []iamTypes.OpenIDConnectProviderListEntry{l},
		}, nil)

	p := iam.GetOpenIDConnectProviderOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetOpenIDConnectProvider(gomock.Any(), gomock.Any()).AnyTimes().Return(&p, nil)

	return aws_client.AwsServices{
		IAM: m,
	}
}

func TestIamOpenidConnectIdentityProviders(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamOpenidConnectIdentityProvidersGenerator{}), buildIamOpenIDConnectProviders, aws_client.TestOptions{})
}
