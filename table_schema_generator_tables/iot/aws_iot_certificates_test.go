package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildIotCertificatesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIOTClient(ctrl)

	certs := iot.ListCertificatesOutput{}
	err := faker.FakeObject(&certs)
	if err != nil {
		t.Fatal(err)
	}
	certs.NextMarker = nil
	m.EXPECT().ListCertificates(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&certs, nil)

	cd := iot.DescribeCertificateOutput{}
	err = faker.FakeObject(&cd)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeCertificate(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&cd, nil)

	p := iot.ListAttachedPoliciesOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	p.NextMarker = nil
	m.EXPECT().ListAttachedPolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&p, nil)

	return aws_client.AwsServices{
		IOT: m,
	}
}

func TestIotCertificates(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIotCertificatesGenerator{}), buildIotCertificatesMock, aws_client.TestOptions{})
}
