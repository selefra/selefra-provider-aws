package shield

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildProtectionGroups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockShieldClient(ctrl)
	pp := shield.ListProtectionGroupsOutput{}
	err := faker.FakeObject(&pp)
	if err != nil {
		t.Fatal(err)
	}
	pp.NextToken = nil
	m.EXPECT().ListProtectionGroups(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&pp, nil)

	tags := shield.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tags, nil)

	return aws_client.AwsServices{
		Shield: m,
	}
}

func TestProtectionGroups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsShieldProtectionGroupsGenerator{}), buildProtectionGroups, aws_client.TestOptions{})
}
