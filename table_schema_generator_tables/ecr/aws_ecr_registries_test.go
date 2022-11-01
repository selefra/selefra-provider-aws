package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEcrRegistriesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEcrClient(ctrl)
	var registryId string
	err := faker.FakeObject(&registryId)
	if err != nil {
		t.Fatal(err)
	}
	rcs := types.ReplicationConfiguration{}
	err = faker.FakeObject(&rcs)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeRegistry(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ecr.DescribeRegistryOutput{
			ReplicationConfiguration:	&rcs,
			RegistryId:			aws.String(registryId),
		}, nil)

	return aws_client.AwsServices{
		ECR: m,
	}
}

func TestEcrRegistries(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEcrRegistriesGenerator{}), buildEcrRegistriesMock, aws_client.TestOptions{})
}
