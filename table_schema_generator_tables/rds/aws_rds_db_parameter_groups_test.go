package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildRDSDBParameterGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockRdsClient(ctrl)
	var g types.DBParameterGroup
	if err := faker.FakeObject(&g); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBParameterGroups(
		gomock.Any(),
		&rds.DescribeDBParameterGroupsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&rds.DescribeDBParameterGroupsOutput{DBParameterGroups: []types.DBParameterGroup{g}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&rds.ListTagsForResourceInput{ResourceName: g.DBParameterGroupArn},
		gomock.Any(),
	).AnyTimes().Return(
		&rds.ListTagsForResourceOutput{
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
		&rds.DescribeDBParametersInput{DBParameterGroupName: g.DBParameterGroupName},
		gomock.Any(),
	).AnyTimes().Return(
		&rds.DescribeDBParametersOutput{Parameters: []types.Parameter{p}},
		nil,
	)
	return aws_client.AwsServices{RDS: mock}
}

func TestRDSDBParameterGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRdsDbParameterGroupsGenerator{}), buildRDSDBParameterGroups, aws_client.TestOptions{})
}
