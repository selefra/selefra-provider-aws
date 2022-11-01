package ecs

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildEcsClusterMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockEcsClient(ctrl)
	services := aws_client.AwsServices{
		ECS: m,
	}
	c := ecsTypes.Cluster{}
	err := faker.FakeObject(&c)
	if err != nil {
		t.Fatal(err)
	}
	ecsOutput := &ecs.DescribeClustersOutput{
		Clusters: []ecsTypes.Cluster{c},
	}
	m.EXPECT().DescribeClusters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(ecsOutput, nil)
	ecsListOutput := &ecs.ListClustersOutput{
		ClusterArns: []string{"randomClusteArn"},
	}
	m.EXPECT().ListClusters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(ecsListOutput, nil)

	servicesList := ecs.ListServicesOutput{
		ServiceArns: []string{"test"},
	}
	m.EXPECT().ListServices(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&servicesList, nil)

	svcs := ecs.DescribeServicesOutput{}
	err = faker.FakeObject(&svcs)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeServices(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&svcs, nil)

	instancesList := ecs.ListContainerInstancesOutput{
		ContainerInstanceArns: []string{"test"},
	}
	m.EXPECT().ListContainerInstances(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&instancesList, nil)

	instances := ecs.DescribeContainerInstancesOutput{}
	err = faker.FakeObject(&instances)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeContainerInstances(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&instances, nil)

	listTasks := ecs.ListTasksOutput{}
	err = faker.FakeObject(&listTasks)
	if err != nil {
		t.Fatal(err)
	}
	listTasks.NextToken = nil
	m.EXPECT().ListTasks(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&listTasks, nil)

	tasks := ecs.DescribeTasksOutput{}
	err = faker.FakeObject(&tasks)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTasks(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tasks, nil)

	return services
}

func TestEcsClusters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsEcsClustersGenerator{}), buildEcsClusterMock, aws_client.TestOptions{})
}
