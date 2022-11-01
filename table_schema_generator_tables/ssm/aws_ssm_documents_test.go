package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildSSMDocuments(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockSSMClient(ctrl)

	docName := "testDocName"
	mock.EXPECT().ListDocuments(
		gomock.Any(),
		&ssm.ListDocumentsInput{Filters: []types.DocumentKeyValuesFilter{{Key: aws.String("Owner"), Values: []string{"Self"}}}},
		gomock.Any(),
	).AnyTimes().Return(
		&ssm.ListDocumentsOutput{DocumentIdentifiers: []types.DocumentIdentifier{{Name: &docName}}},
		nil,
	)

	var d types.DocumentDescription
	if err := faker.FakeObject(&d); err != nil {
		t.Fatal(err)
	}
	d.Name = &docName
	mock.EXPECT().DescribeDocument(
		gomock.Any(),
		&ssm.DescribeDocumentInput{Name: &docName},
		gomock.Any(),
	).AnyTimes().Return(
		&ssm.DescribeDocumentOutput{Document: &d},
		nil,
	)

	mock.EXPECT().DescribeDocumentPermission(
		gomock.Any(),
		&ssm.DescribeDocumentPermissionInput{
			Name:		&docName,
			PermissionType:	types.DocumentPermissionTypeShare,
		},
		gomock.Any(),
	).AnyTimes().Return(
		&ssm.DescribeDocumentPermissionOutput{
			AccountIds:		[]string{"some"},
			AccountSharingInfoList:	[]types.AccountSharingInfo{{AccountId: aws.String("other"), SharedDocumentVersion: aws.String("version")}},
		},
		nil,
	)
	return aws_client.AwsServices{SSM: mock}
}

func TestSSMDocuments(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSsmDocumentsGenerator{}), buildSSMDocuments, aws_client.TestOptions{})
}
