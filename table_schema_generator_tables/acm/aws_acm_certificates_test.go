package acm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildACMCertificates(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockACMClient(ctrl)

	var cs types.CertificateSummary
	if err := faker.FakeObject(&cs); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListCertificates(
		gomock.Any(),
		&acm.ListCertificatesInput{},
		gomock.Any(),
	).AnyTimes().Return(
		&acm.ListCertificatesOutput{CertificateSummaryList: []types.CertificateSummary{cs}},
		nil,
	)

	var cert types.CertificateDetail
	if err := faker.FakeObject(&cert); err != nil {
		t.Fatal(err)
	}
	cert.CertificateArn = cs.CertificateArn
	mock.EXPECT().DescribeCertificate(
		gomock.Any(),
		&acm.DescribeCertificateInput{CertificateArn: cs.CertificateArn},
		gomock.Any(),
	).AnyTimes().Return(
		&acm.DescribeCertificateOutput{Certificate: &cert},
		nil,
	)

	mock.EXPECT().ListTagsForCertificate(
		gomock.Any(),
		&acm.ListTagsForCertificateInput{CertificateArn: cert.CertificateArn},
	).AnyTimes().Return(
		&acm.ListTagsForCertificateOutput{
			Tags: []types.Tag{
				{Key: aws.String("key"), Value: aws.String("value")},
			},
		},
		nil,
	)
	return aws_client.AwsServices{ACM: mock}
}

func TestACMCertificates(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsAcmCertificatesGenerator{}), buildACMCertificates, aws_client.TestOptions{})
}
