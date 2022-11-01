package organizations

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/organizations"
	organizationsTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildOrganizationsAccounts(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockOrganizationsClient(ctrl)
	g := organizationsTypes.Account{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListAccounts(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&organizations.ListAccountsOutput{
			Accounts: []organizationsTypes.Account{g},
		}, nil)

	tt := make([]organizationsTypes.Tag, 3)
	if err := faker.FakeObject(&tt); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&organizations.ListTagsForResourceOutput{
			Tags: tt,
		}, nil)
	return aws_client.AwsServices{
		Organizations: m,
	}
}

func TestOrganizationsAccounts(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsOrganizationsAccountsGenerator{}), buildOrganizationsAccounts, aws_client.TestOptions{})
}
