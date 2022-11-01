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

func buildDataCatalogs(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockAthenaClient(ctrl)

	catalogs := athena.ListDataCatalogsOutput{}
	err := faker.FakeObject(&catalogs)
	if err != nil {
		t.Fatal(err)
	}
	catalogs.NextToken = nil
	m.EXPECT().ListDataCatalogs(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&catalogs, nil)

	catalog := athena.GetDataCatalogOutput{}
	err = faker.FakeObject(&catalog)
	if err != nil {
		t.Fatal(err)
	}
	catalogs.NextToken = nil
	m.EXPECT().GetDataCatalog(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&catalog, nil)

	databases := athena.ListDatabasesOutput{}
	err = faker.FakeObject(&databases)
	if err != nil {
		t.Fatal(err)
	}
	databases.NextToken = nil
	m.EXPECT().ListDatabases(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&databases, nil)

	tags := athena.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tags, nil)

	tables := athena.ListTableMetadataOutput{}
	err = faker.FakeObject(&tables)
	if err != nil {
		t.Fatal(err)
	}
	tables.NextToken = nil
	m.EXPECT().ListTableMetadata(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tables, nil)

	return aws_client.AwsServices{
		Athena: m,
	}
}

func TestDataCatalogs(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAthenaDataCatalogsGenerator{}), buildDataCatalogs, aws_client.TestOptions{})
}
