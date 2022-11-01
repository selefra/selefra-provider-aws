package emr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEMRClusters(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockEmrClient(ctrl)
	var summary types.ClusterSummary
	if err := faker.FakeObject(&summary); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListClusters(gomock.Any(), &emr.ListClustersInput{ClusterStates: []types.ClusterState{
		types.ClusterStateRunning,
		types.ClusterStateStarting,
		types.ClusterStateBootstrapping,
		types.ClusterStateWaiting,
	}}, gomock.Any()).AnyTimes().Return(
		&emr.ListClustersOutput{Clusters: []types.ClusterSummary{summary}},
		nil,
	)

	var cluster types.Cluster
	if err := faker.FakeObject(&cluster); err != nil {
		t.Fatal(err)
	}
	cluster.InstanceCollectionType = types.InstanceCollectionTypeInstanceFleet
	cluster.RepoUpgradeOnBoot = types.RepoUpgradeOnBootNone
	cluster.ScaleDownBehavior = types.ScaleDownBehaviorTerminateAtInstanceHour
	var config types.Configuration
	if err := faker.FakeObject(&config); err != nil {
		t.Fatal(err)
	}
	config.Configurations = []types.Configuration{}
	cluster.Configurations = []types.Configuration{config}
	mock.EXPECT().DescribeCluster(gomock.Any(), &emr.DescribeClusterInput{ClusterId: summary.Id}).AnyTimes().Return(
		&emr.DescribeClusterOutput{Cluster: &cluster},
		nil,
	)
	return aws_client.AwsServices{EMR: mock}
}

func TestEMRClusters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEmrClustersGenerator{}), buildEMRClusters, aws_client.TestOptions{})
}
