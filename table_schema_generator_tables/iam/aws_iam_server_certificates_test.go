package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildIamServerCerts(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIamClient(ctrl)
	u := iamTypes.ServerCertificateMetadata{}
	err := faker.FakeObject(&u)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListServerCertificates(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListServerCertificatesOutput{
			ServerCertificateMetadataList: []iamTypes.ServerCertificateMetadata{u},
		}, nil)

	return aws_client.AwsServices{
		IAM: m,
	}
}

func TestIamServerCertificates(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamServerCertificatesGenerator{}), buildIamServerCerts, aws_client.TestOptions{})
}
