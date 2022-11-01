package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func buildRegistriesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlueClient(ctrl)

	var r types.RegistryListItem
	require.NoError(t, faker.FakeObject(&r))
	m.EXPECT().ListRegistries(
		gomock.Any(),
		&glue.ListRegistriesInput{MaxResults: aws.Int32(100)},
	).AnyTimes().Return(
		&glue.ListRegistriesOutput{Registries: []types.RegistryListItem{r}},
		nil,
	)

	m.EXPECT().GetTags(
		gomock.Any(),
		&glue.GetTagsInput{ResourceArn: r.RegistryArn},
	).AnyTimes().Return(
		&glue.GetTagsOutput{Tags: map[string]string{"tag": "value"}},
		nil,
	)

	var s glue.GetSchemaOutput
	require.NoError(t, faker.FakeObject(&s))
	m.EXPECT().ListSchemas(
		gomock.Any(),
		&glue.ListSchemasInput{
			RegistryId:	&types.RegistryId{RegistryArn: r.RegistryArn},
			MaxResults:	aws.Int32(100),
		},
	).AnyTimes().Return(
		&glue.ListSchemasOutput{Schemas: []types.SchemaListItem{{SchemaArn: s.SchemaArn}}},
		nil,
	)

	m.EXPECT().GetSchema(
		gomock.Any(),
		&glue.GetSchemaInput{SchemaId: &types.SchemaId{SchemaArn: s.SchemaArn}},
	).AnyTimes().Return(&s, nil)

	m.EXPECT().GetTags(
		gomock.Any(),
		&glue.GetTagsInput{ResourceArn: s.SchemaArn},
	).AnyTimes().Return(
		&glue.GetTagsOutput{Tags: map[string]string{"tag": "value"}},
		nil,
	)

	var lsv glue.ListSchemaVersionsOutput
	require.NoError(t, faker.FakeObject(&lsv))
	lsv.NextToken = nil
	m.EXPECT().ListSchemaVersions(
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(&lsv, nil)

	var sv glue.GetSchemaVersionOutput
	require.NoError(t, faker.FakeObject(&sv))
	m.EXPECT().GetSchemaVersion(
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(&sv, nil)

	var sm glue.QuerySchemaVersionMetadataOutput
	require.NoError(t, faker.FakeObject(&sm))
	sm.NextToken = nil
	m.EXPECT().QuerySchemaVersionMetadata(
		gomock.Any(),
		gomock.Any(),
	).AnyTimes().Return(&sm, nil)

	return aws_client.AwsServices{
		Glue: m,
	}
}

func TestRegistries(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlueRegistriesGenerator{}), buildRegistriesMock, aws_client.TestOptions{})
}
