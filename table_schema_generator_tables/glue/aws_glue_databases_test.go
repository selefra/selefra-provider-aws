package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func buildDatabasesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlueClient(ctrl)

	db := glue.GetDatabasesOutput{}
	require.NoError(t, faker.FakeObject(&db))
	db.NextToken = nil
	m.EXPECT().GetDatabases(gomock.Any(), gomock.Any()).AnyTimes().Return(&db, nil)

	tb := glue.GetTablesOutput{}
	require.NoError(t, faker.FakeObject(&tb))
	tb.NextToken = nil
	m.EXPECT().GetTables(gomock.Any(), gomock.Any()).AnyTimes().Return(&tb, nil)

	i := glue.GetPartitionIndexesOutput{}
	require.NoError(t, faker.FakeObject(&i))
	i.NextToken = nil
	m.EXPECT().GetPartitionIndexes(gomock.Any(), gomock.Any()).AnyTimes().Return(&i, nil)

	tags := glue.GetTagsOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	m.EXPECT().GetTags(gomock.Any(), gomock.Any()).AnyTimes().Return(&tags, nil)

	return aws_client.AwsServices{
		Glue: m,
	}
}

func TestDatabases(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlueDatabasesGenerator{}), buildDatabasesMock, aws_client.TestOptions{})
}
