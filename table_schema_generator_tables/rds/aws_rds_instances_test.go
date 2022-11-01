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

func buildRdsDBInstances(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBInstance{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBInstances(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&rds.DescribeDBInstancesOutput{
			DBInstances: []rdsTypes.DBInstance{l},
		}, nil)
	return aws_client.AwsServices{
		RDS: m,
	}
}

func TestRdsInstances(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRdsInstancesGenerator{}), buildRdsDBInstances, aws_client.TestOptions{})
}
