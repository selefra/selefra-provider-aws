package redshift

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildRedshiftSubnetGroupsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockRedshiftClient(ctrl)

	g := types.ClusterSubnetGroup{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeClusterSubnetGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&redshift.DescribeClusterSubnetGroupsOutput{
			ClusterSubnetGroups: []types.ClusterSubnetGroup{g},
		}, nil)
	return aws_client.AwsServices{
		Redshift: m,
	}
}

func TestRedshiftSubnetGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRedshiftSubnetGroupsGenerator{}), buildRedshiftSubnetGroupsMock, aws_client.TestOptions{})
}
