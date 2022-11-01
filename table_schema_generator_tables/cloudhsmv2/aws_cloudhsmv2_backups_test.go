package cloudhsmv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildHSMBackups(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockCloudHSMV2Client(ctrl)

	var backups []types.Backup
	if err := faker.FakeObject(&backups); err != nil {
		t.Fatal(err)
	}

	mock.EXPECT().DescribeBackups(
		gomock.Any(),
		&cloudhsmv2.DescribeBackupsInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&cloudhsmv2.DescribeBackupsOutput{Backups: backups},
		nil,
	)

	return aws_client.AwsServices{CloudHSMV2: mock}
}

func TestBackups(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsCloudhsmv2BackupsGenerator{}), buildHSMBackups, aws_client.TestOptions{})
}
