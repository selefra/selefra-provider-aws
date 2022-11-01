package eks

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEksClusters(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEksClient(ctrl)
	l := eks.DescribeClusterOutput{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListClusters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&eks.ListClustersOutput{
			Clusters: []string{"test-cluster"},
		}, nil)
	m.EXPECT().DescribeCluster(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&l, nil)
	return aws_client.AwsServices{
		Eks: m,
	}
}

func TestEksClusters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEksClustersGenerator{}), buildEksClusters, aws_client.TestOptions{})
}
