package lightsail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildDatabasesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockLightsailClient(ctrl)

	b := lightsail.GetRelationalDatabasesOutput{}
	err := faker.FakeObject(&b)
	if err != nil {
		t.Fatal(err)
	}
	b.NextPageToken = nil
	m.EXPECT().GetRelationalDatabases(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&b, nil)

	ac := lightsail.GetRelationalDatabaseParametersOutput{}
	err = faker.FakeObject(&ac)
	if err != nil {
		t.Fatal(err)
	}
	ac.NextPageToken = nil

	m.EXPECT().GetRelationalDatabaseParameters(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ac, nil)

	e := lightsail.GetRelationalDatabaseEventsOutput{}
	err = faker.FakeObject(&e)
	if err != nil {
		t.Fatal(err)
	}
	e.NextPageToken = nil

	m.EXPECT().GetRelationalDatabaseEvents(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&e, nil)
	ls := lightsail.GetRelationalDatabaseLogStreamsOutput{}
	err = faker.FakeObject(&ls)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetRelationalDatabaseLogStreams(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ls, nil)

	le := lightsail.GetRelationalDatabaseLogEventsOutput{}
	err = faker.FakeObject(&le)
	if err != nil {
		t.Fatal(err)
	}
	le.NextForwardToken = nil
	m.EXPECT().GetRelationalDatabaseLogEvents(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&le, nil)

	return aws_client.AwsServices{
		Lightsail: m,
	}
}

func TestDatabases(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLightsailDatabasesGenerator{}), buildDatabasesMock, aws_client.TestOptions{})
}
