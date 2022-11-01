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

func buildElasticbeanstalkEnvironments(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockElasticbeanstalkClient(ctrl)

	la := elasticbeanstalkTypes.ApplicationDescription{}
	err := faker.FakeObject(&la)
	if err != nil {
		t.Fatal(err)
	}

	l := elasticbeanstalkTypes.EnvironmentDescription{
		ApplicationName: la.ApplicationName,
	}
	err = faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeEnvironments(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&elasticbeanstalk.DescribeEnvironmentsOutput{
			Environments: []elasticbeanstalkTypes.EnvironmentDescription{l},
		}, nil)

	tags := elasticbeanstalk.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&tags, nil)

	configSettingsOutput := elasticbeanstalk.DescribeConfigurationSettingsOutput{}
	err = faker.FakeObject(&configSettingsOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConfigurationSettings(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&configSettingsOutput, nil)

	configOptsOutput := elasticbeanstalk.DescribeConfigurationOptionsOutput{}
	err = faker.FakeObject(&configOptsOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConfigurationOptions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&configOptsOutput, nil)

	return aws_client.AwsServices{
		ElasticBeanstalk: m,
	}
}

func TestElasticbeanstalkEnvironments(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsElasticbeanstalkEnvironmentsGenerator{}), buildElasticbeanstalkEnvironments, aws_client.TestOptions{})
}
