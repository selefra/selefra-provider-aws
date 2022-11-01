package guardduty

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	gdTypes "github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildGuardDutyDetectors(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGuardDutyClient(ctrl)

	var d guardduty.GetDetectorOutput
	if err := faker.FakeObject(&d); err != nil {
		t.Fatal(err)
	}
	d.CreatedAt = aws.String(time.Now().Format(time.RFC3339))
	d.UpdatedAt = aws.String(time.Now().Format(time.RFC3339))

	m.EXPECT().ListDetectors(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&guardduty.ListDetectorsOutput{
			DetectorIds: []string{""},
		}, nil)

	m.EXPECT().GetDetector(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&d, nil)

	var member gdTypes.Member
	if err := faker.FakeObject(&member); err != nil {
		t.Fatal(err)
	}
	member.UpdatedAt = aws.String(time.Now().Format(time.RFC3339))
	member.InvitedAt = aws.String(time.Now().Format(time.RFC3339))

	m.EXPECT().ListMembers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&guardduty.ListMembersOutput{Members: []gdTypes.Member{member}}, nil,
	)
	return aws_client.AwsServices{
		GuardDuty: m,
	}
}

func TestGuarddutyDetectors(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGuarddutyDetectorsGenerator{}), buildGuardDutyDetectors, aws_client.TestOptions{})
}
