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

func buildProtections(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockShieldClient(ctrl)
	protection := shield.ListProtectionsOutput{}
	err := faker.FakeObject(&protection)
	if err != nil {
		t.Fatal(err)
	}
	protection.NextToken = nil
	m.EXPECT().ListProtections(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&protection, nil)

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

func TestProtections(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsShieldProtectionsGenerator{}), buildProtections, aws_client.TestOptions{})
}
