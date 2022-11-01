package neptune

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildNeptuneDBParameterGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockNeptuneClient(ctrl)
	var g types.DBParameterGroup
	if err := faker.FakeObject(&g); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBParameterGroups(
		gomock.Any(),
		&neptune.DescribeDBParameterGroupsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&neptune.DescribeDBParameterGroupsOutput{DBParameterGroups: []types.DBParameterGroup{g}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: g.DBParameterGroupArn},
		gomock.Any(),
	).AnyTimes().Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)

	var p types.Parameter
	if err := faker.FakeObject(&p); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBParameters(
		gomock.Any(),
		&neptune.DescribeDBParametersInput{DBParameterGroupName: g.DBParameterGroupName},
		gomock.Any(),
	).AnyTimes().Return(
		&neptune.DescribeDBParametersOutput{Parameters: []types.Parameter{p}},
		nil,
	)
	return aws_client.AwsServices{Neptune: mock}
}

func TestNeptuneDBParameterGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsNeptuneDbParameterGroupsGenerator{}), buildNeptuneDBParameterGroups, aws_client.TestOptions{})
}
