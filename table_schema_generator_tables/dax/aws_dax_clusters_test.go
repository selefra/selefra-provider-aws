package dax

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/dax/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildDAXClustersMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockDAXClient(ctrl)
	services := aws_client.AwsServices{
		DAX: m,
	}
	c := types.Cluster{}
	if err := faker.FakeObject(&c); err != nil {
		t.Fatal(err)
	}
	daxOutput := &dax.DescribeClustersOutput{
		Clusters: []types.Cluster{c},
	}
	m.EXPECT().DescribeClusters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		daxOutput,
		nil,
	)

	tags := &dax.ListTagsOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		tags,
		nil,
	)
	return services
}

func TestDAXClusters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDaxClustersGenerator{}), buildDAXClustersMock, aws_client.TestOptions{})
}
