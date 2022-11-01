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

func buildRdsClusterParameterGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockRdsClient(ctrl)
	var g types.DBClusterParameterGroup
	if err := faker.FakeObject(&g); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBClusterParameterGroups(
		gomock.Any(),
		&rds.DescribeDBClusterParameterGroupsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&rds.DescribeDBClusterParameterGroupsOutput{DBClusterParameterGroups: []types.DBClusterParameterGroup{g}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&rds.ListTagsForResourceInput{ResourceName: g.DBClusterParameterGroupArn},
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
	mock.EXPECT().DescribeDBClusterParameters(
		gomock.Any(),
		&rds.DescribeDBClusterParametersInput{DBClusterParameterGroupName: g.DBClusterParameterGroupName},
		gomock.Any(),
	).AnyTimes().Return(
		&rds.DescribeDBClusterParametersOutput{Parameters: []types.Parameter{p}},
		nil,
	)
	return aws_client.AwsServices{RDS: mock}
}

func TestRdsClusterParameterGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRdsClusterParameterGroupsGenerator{}), buildRdsClusterParameterGroups, aws_client.TestOptions{})
}
