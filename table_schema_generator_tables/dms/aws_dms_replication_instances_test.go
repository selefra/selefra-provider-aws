package dms

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildDmsReplicationInstances(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockDatabasemigrationserviceClient(ctrl)
	l := types.ReplicationInstance{}
	if err := faker.FakeObject(&l); err != nil {
		t.Fatal(err)
	}
	l.ReplicationInstancePrivateIpAddress = aws.String("1.2.3.4")
	l.ReplicationInstancePrivateIpAddresses = []string{"1.2.3.4"}
	l.ReplicationInstancePublicIpAddress = aws.String("1.2.3.4")
	l.ReplicationInstancePublicIpAddresses = []string{"1.2.3.4"}
	m.EXPECT().DescribeReplicationInstances(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&databasemigrationservice.DescribeReplicationInstancesOutput{
			ReplicationInstances: []types.ReplicationInstance{l},
		}, nil)
	lt := types.Tag{}
	if err := faker.FakeObject(&lt); err != nil {
		t.Fatal(err)
	}
	lt.ResourceArn = l.ReplicationInstanceArn
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&databasemigrationservice.ListTagsForResourceOutput{
			TagList: []types.Tag{lt},
		}, nil)
	return aws_client.AwsServices{
		DMS: m,
	}
}

func TestDmsReplicationInstances(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDmsReplicationInstancesGenerator{}), buildDmsReplicationInstances, aws_client.TestOptions{})
}
