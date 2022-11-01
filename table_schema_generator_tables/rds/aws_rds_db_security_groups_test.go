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

func buildRdsDbSecurityGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockRdsClient(ctrl)
	var g types.DBSecurityGroup
	if err := faker.FakeObject(&g); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBSecurityGroups(
		gomock.Any(),
		&rds.DescribeDBSecurityGroupsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&rds.DescribeDBSecurityGroupsOutput{DBSecurityGroups: []types.DBSecurityGroup{g}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&rds.ListTagsForResourceInput{ResourceName: g.DBSecurityGroupArn},
		gomock.Any(),
	).AnyTimes().Return(
		&rds.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)

	return aws_client.AwsServices{RDS: mock}
}

func TestRDSDBSecurityGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRdsDbSecurityGroupsGenerator{}), buildRdsDbSecurityGroups, aws_client.TestOptions{})
}
