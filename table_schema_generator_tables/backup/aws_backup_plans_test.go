package backup

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildBackupPlansMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockBackupClient(ctrl)

	var plan backup.GetBackupPlanOutput
	if err := faker.FakeObject(&plan); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListBackupPlans(
		gomock.Any(),
		&backup.ListBackupPlansInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).AnyTimes().Return(
		&backup.ListBackupPlansOutput{BackupPlansList: []types.BackupPlansListMember{
			{
				BackupPlanId:	plan.BackupPlanId,
				VersionId:	plan.VersionId,
			},
		}},
		nil,
	)

	m.EXPECT().GetBackupPlan(
		gomock.Any(),
		&backup.GetBackupPlanInput{BackupPlanId: plan.BackupPlanId, VersionId: plan.VersionId},
		gomock.Any(),
	).AnyTimes().Return(
		&plan,
		nil,
	)

	m.EXPECT().ListTags(
		gomock.Any(),
		&backup.ListTagsInput{ResourceArn: plan.BackupPlanArn},
		gomock.Any(),
	).AnyTimes().Return(
		&backup.ListTagsOutput{
			Tags: map[string]string{"plan1": "value1"},
		},
		nil,
	)

	var selection backup.GetBackupSelectionOutput
	if err := faker.FakeObject(&selection); err != nil {
		t.Fatal(err)
	}
	selection.BackupPlanId = plan.BackupPlanId
	m.EXPECT().ListBackupSelections(
		gomock.Any(),
		&backup.ListBackupSelectionsInput{
			BackupPlanId:	plan.BackupPlanId,
			MaxResults:	aws.Int32(1000),
		},
		gomock.Any(),
	).AnyTimes().Return(
		&backup.ListBackupSelectionsOutput{
			BackupSelectionsList: []types.BackupSelectionsListMember{{SelectionId: selection.SelectionId}},
		},
		nil,
	)

	m.EXPECT().GetBackupSelection(
		gomock.Any(),
		&backup.GetBackupSelectionInput{
			BackupPlanId:	plan.BackupPlanId,
			SelectionId:	selection.SelectionId,
		},
		gomock.Any(),
	).AnyTimes().Return(&selection, nil)

	return aws_client.AwsServices{
		Backup: m,
	}
}

func TestPlans(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsBackupPlansGenerator{}), buildBackupPlansMock, aws_client.TestOptions{})
}
