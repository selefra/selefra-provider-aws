package redshift

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildRedshiftClustersMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockRedshiftClient(ctrl)
	g := types.Cluster{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}
	p := types.Parameter{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	logging := redshift.DescribeLoggingStatusOutput{}
	err = faker.FakeObject(&logging)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeClusters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&redshift.DescribeClustersOutput{
			Clusters: []types.Cluster{g},
		}, nil)
	m.EXPECT().DescribeClusterParameters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&redshift.DescribeClusterParametersOutput{
			Parameters: []types.Parameter{p},
		}, nil)
	m.EXPECT().DescribeLoggingStatus(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&logging, nil)

	var snap types.Snapshot
	if err := faker.FakeObject(&snap); err != nil {
		t.Fatal(err)
	}
	snap.ClusterIdentifier = g.ClusterIdentifier
	snap.ClusterCreateTime = g.ClusterCreateTime
	m.EXPECT().DescribeClusterSnapshots(
		gomock.Any(),
		&redshift.DescribeClusterSnapshotsInput{
			ClusterExists:		aws.Bool(true),
			ClusterIdentifier:	g.ClusterIdentifier,
			MaxRecords:		aws.Int32(100),
		},
		gomock.Any(),
	).AnyTimes().Return(
		&redshift.DescribeClusterSnapshotsOutput{Snapshots: []types.Snapshot{snap}},
		nil,
	)

	return aws_client.AwsServices{
		Redshift: m,
	}
}

func TestRedshiftClusters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRedshiftClustersGenerator{}), buildRedshiftClustersMock, aws_client.TestOptions{})
}
