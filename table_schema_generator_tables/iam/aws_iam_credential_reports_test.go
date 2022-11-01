package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

var exampleReport = `user,arn,user_creation_time,password_enabled,password_last_used,password_last_changed,password_next_rotation,mfa_active,access_key_1_active,access_key_1_last_rotated,access_key_1_last_used_date,access_key_1_last_used_region,access_key_1_last_used_service,access_key_2_active,access_key_2_last_rotated,access_key_2_last_used_date,access_key_2_last_used_region,access_key_2_last_used_service,cert_1_active,cert_1_last_rotated,cert_2_active,cert_2_last_rotated
user-readonly,arn:aws:iam::123456789012:user/user-readonly,2022-08-31T11:10:33+00:00,false,2022-08-30T11:10:33+00:00,2022-08-31T11:10:33+00:00,2023-08-31T11:10:33+00:00,false,true,2022-08-31T11:10:34+00:00,2022-08-31T11:23:00+00:00,us-east-1,iam,true,2022-08-31T11:10:33+00:00,2022-08-31T11:10:33+00:00,N/A,N/A,false,2022-08-31T11:10:33+00:00,false,2022-08-31T11:10:33+00:00`

func buildCredentialReports(_ *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIamClient(ctrl)
	m.EXPECT().GetCredentialReport(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.GetCredentialReportOutput{
			Content: []byte(exampleReport),
		}, nil)

	return aws_client.AwsServices{
		IAM: m,
	}
}

func TestCredentialReports(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamCredentialReportsGenerator{}), buildCredentialReports, aws_client.TestOptions{})
}
