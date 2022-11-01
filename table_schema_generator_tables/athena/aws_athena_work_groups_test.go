package athena

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildWorkGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockAthenaClient(ctrl)

	workGroupsOutput := athena.ListWorkGroupsOutput{}
	err := faker.FakeObject(&workGroupsOutput)
	if err != nil {
		t.Fatal(err)
	}
	workGroupsOutput.NextToken = nil
	m.EXPECT().ListWorkGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&workGroupsOutput, nil)

	workGroup := athena.GetWorkGroupOutput{}
	err = faker.FakeObject(&workGroup)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetWorkGroup(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&workGroup, nil)

	namedQueriesOutput := athena.ListNamedQueriesOutput{}
	err = faker.FakeObject(&namedQueriesOutput)
	if err != nil {
		t.Fatal(err)
	}
	namedQueriesOutput.NextToken = nil
	m.EXPECT().ListNamedQueries(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&namedQueriesOutput, nil)

	queryExecutionsOutput := athena.ListQueryExecutionsOutput{}
	err = faker.FakeObject(&queryExecutionsOutput)
	if err != nil {
		t.Fatal(err)
	}
	queryExecutionsOutput.NextToken = nil
	m.EXPECT().ListQueryExecutions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&queryExecutionsOutput, nil)

	preparedStatementsOutput := athena.ListPreparedStatementsOutput{}
	err = faker.FakeObject(&preparedStatementsOutput)
	if err != nil {
		t.Fatal(err)
	}
	preparedStatementsOutput.NextToken = nil
	m.EXPECT().ListPreparedStatements(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&preparedStatementsOutput, nil)

	tags := athena.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tags, nil)

	preparedStatement := athena.GetPreparedStatementOutput{}
	err = faker.FakeObject(&preparedStatement)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetPreparedStatement(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&preparedStatement, nil)

	namedQuery := athena.GetNamedQueryOutput{}
	err = faker.FakeObject(&namedQuery)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetNamedQuery(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&namedQuery, nil)
	queryExecution := athena.GetQueryExecutionOutput{}
	err = faker.FakeObject(&queryExecution)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetQueryExecution(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&queryExecution, nil)

	return aws_client.AwsServices{
		Athena: m,
	}
}

func TestWorkGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAthenaWorkGroupsGenerator{}), buildWorkGroups, aws_client.TestOptions{})
}
