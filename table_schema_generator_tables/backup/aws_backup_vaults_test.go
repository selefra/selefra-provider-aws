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

func buildBackupVaultsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockBackupClient(ctrl)

	var vault types.BackupVaultListMember
	if err := faker.FakeObject(&vault); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListBackupVaults(
		gomock.Any(),
		&backup.ListBackupVaultsInput{MaxResults: aws.Int32(1000)},
		gomock.Any(),
	).AnyTimes().Return(
		&backup.ListBackupVaultsOutput{BackupVaultList: []types.BackupVaultListMember{vault}},
		nil,
	)

	m.EXPECT().ListTags(
		gomock.Any(),
		&backup.ListTagsInput{ResourceArn: vault.BackupVaultArn},
		gomock.Any(),
	).AnyTimes().Return(
		&backup.ListTagsOutput{
			Tags: map[string]string{"tag1": "value1"},
		},
		nil,
	)

	m.EXPECT().GetBackupVaultAccessPolicy(
		gomock.Any(),
		&backup.GetBackupVaultAccessPolicyInput{BackupVaultName: vault.BackupVaultName},
		gomock.Any(),
	).AnyTimes().Return(
		&backup.GetBackupVaultAccessPolicyOutput{
			Policy: aws.String(`{"key":"value"}`),
		},
		nil,
	)

	m.EXPECT().GetBackupVaultNotifications(
		gomock.Any(),
		&backup.GetBackupVaultNotificationsInput{BackupVaultName: vault.BackupVaultName},
		gomock.Any(),
	).AnyTimes().Return(
		&backup.GetBackupVaultNotificationsOutput{
			BackupVaultEvents:	[]types.BackupVaultEvent{types.BackupVaultEventBackupJobFailed},
			SNSTopicArn:		aws.String("not really an ARN"),
		},
		nil,
	)

	var rp types.RecoveryPointByBackupVault
	if err := faker.FakeObject(&rp); err != nil {
		t.Fatal(err)
	}
	rp.ResourceArn = aws.String("arn:aws:s3:eu-central-1:testAccount:resource/id")

	m.EXPECT().ListRecoveryPointsByBackupVault(
		gomock.Any(),
		&backup.ListRecoveryPointsByBackupVaultInput{BackupVaultName: vault.BackupVaultName, MaxResults: aws.Int32(100)},
		gomock.Any(),
	).AnyTimes().Return(
		&backup.ListRecoveryPointsByBackupVaultOutput{RecoveryPoints: []types.RecoveryPointByBackupVault{rp}},
		nil,
	)

	m.EXPECT().ListTags(
		gomock.Any(),
		&backup.ListTagsInput{ResourceArn: rp.RecoveryPointArn},
		gomock.Any(),
	).AnyTimes().Return(
		&backup.ListTagsOutput{
			Tags: map[string]string{"tag1": "value1"},
		},
		nil,
	)

	return aws_client.AwsServices{
		Backup: m,
	}
}

func TestVaults(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsBackupVaultsGenerator{}), buildBackupVaultsMock, aws_client.TestOptions{})
}
