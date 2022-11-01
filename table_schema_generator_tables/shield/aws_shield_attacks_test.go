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

func buildAttacks(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockShieldClient(ctrl)
	protection := shield.ListAttacksOutput{}
	err := faker.FakeObject(&protection)
	if err != nil {
		t.Fatal(err)
	}
	protection.NextToken = nil
	m.EXPECT().ListAttacks(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&protection, nil)

	tags := shield.DescribeAttackOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeAttack(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tags, nil)
	return aws_client.AwsServices{
		Shield: m,
	}
}

func TestAttacks(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsShieldAttacksGenerator{}), buildAttacks, aws_client.TestOptions{})
}
