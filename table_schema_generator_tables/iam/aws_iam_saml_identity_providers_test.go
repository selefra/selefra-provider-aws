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

func buildIamSAMLProviders(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIamClient(ctrl)
	l := iamTypes.SAMLProviderListEntry{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListSAMLProviders(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListSAMLProvidersOutput{
			SAMLProviderList: []iamTypes.SAMLProviderListEntry{l},
		}, nil)

	p := iam.GetSAMLProviderOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetSAMLProvider(gomock.Any(), gomock.Any()).AnyTimes().Return(&p, nil)

	return aws_client.AwsServices{
		IAM: m,
	}
}

func TestIAMSamlIdentityProviders(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamSamlIdentityProvidersGenerator{}), buildIamSAMLProviders, aws_client.TestOptions{})
}
