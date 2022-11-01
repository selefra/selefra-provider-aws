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

func buildRdsDBClusters(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockRdsClient(ctrl)
	l := rdsTypes.DBCluster{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeDBClusters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&rds.DescribeDBClustersOutput{
			DBClusters: []rdsTypes.DBCluster{l},
		}, nil)
	return aws_client.AwsServices{
		RDS: m,
	}
}

func TestRdsClusters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRdsClustersGenerator{}), buildRdsDBClusters, aws_client.TestOptions{})
}
