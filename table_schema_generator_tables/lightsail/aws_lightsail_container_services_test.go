package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildContainerServicesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockLightsailClient(ctrl)

	dep := types.ContainerServiceDeployment{State: "test", Containers: map[string]types.Container{"test": {Image: aws.String("test")}}}
	err := faker.FakeObject(&dep)
	if err != nil {
		t.Fatal(err)
	}
	service := types.ContainerService{CurrentDeployment: &dep, NextDeployment: &dep, Power: "test", ResourceType: "test", State: "test"}
	err = faker.FakeObject(&service)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetContainerServices(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lightsail.GetContainerServicesOutput{ContainerServices: []types.ContainerService{service}}, nil)

	m.EXPECT().GetContainerServiceDeployments(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lightsail.GetContainerServiceDeploymentsOutput{Deployments: []types.ContainerServiceDeployment{dep}}, nil)

	i := lightsail.GetContainerImagesOutput{}
	err = faker.FakeObject(&i)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetContainerImages(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&i, nil)

	return aws_client.AwsServices{
		Lightsail: m,
	}
}

func TestContainerServices(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLightsailContainerServicesGenerator{}), buildContainerServicesMock, aws_client.TestOptions{})
}
