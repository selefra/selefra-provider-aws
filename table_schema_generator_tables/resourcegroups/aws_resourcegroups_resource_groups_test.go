package resourcegroups

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildResourceGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockResourceGroupsClient(ctrl)
	gId := types.GroupIdentifier{}
	err := faker.FakeObject(&gId)
	if err != nil {
		t.Fatal(err)
	}

	groupResponse := types.Group{}
	err = faker.FakeObject(&groupResponse)
	if err != nil {
		t.Fatal(err)
	}

	tagsResponse := resourcegroups.GetTagsOutput{}
	err = faker.FakeObject(&tagsResponse)
	if err != nil {
		t.Fatal(err)
	}

	query := types.GroupQuery{}
	err = faker.FakeObject(&query)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&resourcegroups.ListGroupsOutput{
			GroupIdentifiers: []types.GroupIdentifier{gId},
		}, nil)
	m.EXPECT().GetGroup(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&resourcegroups.GetGroupOutput{
			Group: &groupResponse,
		}, nil)
	m.EXPECT().GetTags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tagsResponse, nil)
	m.EXPECT().GetGroupQuery(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&resourcegroups.GetGroupQueryOutput{
			GroupQuery: &query,
		}, nil)

	return aws_client.AwsServices{
		ResourceGroups: m,
	}
}

func TestResourceGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsResourcegroupsResourceGroupsGenerator{}), buildResourceGroupsMock, aws_client.TestOptions{})
}
