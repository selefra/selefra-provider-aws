package dynamodb

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildDynamodbTablesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockDynamoDBClient(ctrl)
	services := aws_client.AwsServices{
		DynamoDB: m,
	}
	var tableName string
	if err := faker.FakeObject(&tableName); err != nil {
		t.Fatal(err)
	}
	listOutput := &dynamodb.ListTablesOutput{
		TableNames: []string{tableName},
	}
	m.EXPECT().ListTables(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		listOutput,
		nil,
	)

	descOutput := &dynamodb.DescribeTableOutput{
		Table: &types.TableDescription{
			TableName: &tableName,
		},
	}
	if err := faker.FakeObject(descOutput.Table); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTable(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		descOutput,
		nil,
	)

	repOutput := &dynamodb.DescribeTableReplicaAutoScalingOutput{
		TableAutoScalingDescription: &types.TableAutoScalingDescription{
			TableName:	&tableName,
			TableStatus:	types.TableStatusActive,
		},
	}
	if err := faker.FakeObject(&repOutput.TableAutoScalingDescription.Replicas); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeTableReplicaAutoScaling(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		repOutput,
		nil,
	)

	cbOutput := &dynamodb.DescribeContinuousBackupsOutput{
		ContinuousBackupsDescription: &types.ContinuousBackupsDescription{},
	}
	if err := faker.FakeObject(&cbOutput.ContinuousBackupsDescription); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeContinuousBackups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		cbOutput,
		nil,
	)

	tags := &dynamodb.ListTagsOfResourceOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsOfResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		tags,
		nil,
	)
	return services
}

func TestDynamodbTables(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDynamodbTablesGenerator{}), buildDynamodbTablesMock, aws_client.TestOptions{})
}
