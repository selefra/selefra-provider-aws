package xray

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/xray"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockXrayClient(ctrl)

	test := "test"

	var group types.GroupSummary
	if err := faker.FakeObject(&group); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().GetGroups(
		gomock.Any(),
		&xray.GetGroupsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&xray.GetGroupsOutput{
			Groups: []types.GroupSummary{
				group,
			},
		},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&xray.ListTagsForResourceInput{ResourceARN: group.GroupARN},
		gomock.Any(),
	).AnyTimes().Return(
		&xray.ListTagsForResourceOutput{
			Tags: []types.Tag{
				{
					Key:	&test,
					Value:	&test,
				},
			},
		},
		nil,
	)

	return aws_client.AwsServices{Xray: mock}
}

func TestXrayGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsXrayGroupsGenerator{}), buildGroups, aws_client.TestOptions{})
}
