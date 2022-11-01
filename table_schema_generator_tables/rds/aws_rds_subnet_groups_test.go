package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	rdsTypes "github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildRdsDBSubnetGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBSubnetGroup{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&rds.DescribeDBSubnetGroupsOutput{
			DBSubnetGroups: []rdsTypes.DBSubnetGroup{l},
		}, nil)
	return aws_client.AwsServices{
		RDS: m,
	}
}

func TestRdsSubnetGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRdsSubnetGroupsGenerator{}), buildRdsDBSubnetGroups, aws_client.TestOptions{})
}
