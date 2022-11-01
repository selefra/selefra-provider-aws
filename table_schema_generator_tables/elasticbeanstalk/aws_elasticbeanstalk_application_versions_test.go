package elasticbeanstalk

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	elasticbeanstalkTypes "github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildElasticbeanstalkApplicationVersions(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationVersionDescription{}
	err := faker.FakeObject(&la)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeApplicationVersions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticbeanstalk.DescribeApplicationVersionsOutput{
			ApplicationVersions: []elasticbeanstalkTypes.ApplicationVersionDescription{la},
		}, nil)

	return aws_client.AwsServices{
		ElasticBeanstalk: m,
	}
}

func TestElasticbeanstalkApplicationVersions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElasticbeanstalkApplicationVersionsGenerator{}), buildElasticbeanstalkApplicationVersions, aws_client.TestOptions{})
}
