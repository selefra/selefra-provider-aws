package appsync

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildAppsyncGraphqlApisMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockAppSyncClient(ctrl)
	l := types.GraphqlApi{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListGraphqlApis(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&appsync.ListGraphqlApisOutput{
			GraphqlApis: []types.GraphqlApi{l},
		}, nil)

	return aws_client.AwsServices{
		AppSync: m,
	}
}

func TestAppSyncGraphqlApis(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAppsyncGraphqlApisGenerator{}), buildAppsyncGraphqlApisMock, aws_client.TestOptions{})
}
