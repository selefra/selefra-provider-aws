package cloudhsmv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildHSMClusters(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockCloudHSMV2Client(ctrl)

	var clusters []types.Cluster
	if err := faker.FakeObject(&clusters); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeClusters(
		gomock.Any(),
		&cloudhsmv2.DescribeClustersInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&cloudhsmv2.DescribeClustersOutput{Clusters: clusters},
		nil,
	)

	return aws_client.AwsServices{CloudHSMV2: mock}
}

func TestClusters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCloudhsmv2ClustersGenerator{}), buildHSMClusters, aws_client.TestOptions{})
}
