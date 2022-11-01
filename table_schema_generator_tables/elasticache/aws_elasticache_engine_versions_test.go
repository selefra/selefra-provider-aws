package elasticache

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildElasticacheEngineVersions(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mockElasticache := mocks.NewMockElastiCache(ctrl)
	output := elasticache.DescribeCacheEngineVersionsOutput{}
	err := faker.FakeObject(&output)
	output.Marker = nil
	if err != nil {
		t.Fatal(err)
	}

	mockElasticache.EXPECT().DescribeCacheEngineVersions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&output, nil)

	return aws_client.AwsServices{
		ElastiCache: mockElasticache,
	}
}

func TestElasticacheEngineVersions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElasticacheEngineVersionsGenerator{}), buildElasticacheEngineVersions, aws_client.TestOptions{})
}
