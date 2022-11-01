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

func buildIotCaCertificatesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIOTClient(ctrl)

	ca := iot.ListCACertificatesOutput{}
	err := faker.FakeObject(&ca)
	if err != nil {
		t.Fatal(err)
	}
	ca.NextMarker = nil
	m.EXPECT().ListCACertificates(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ca, nil)

	cd := iot.DescribeCACertificateOutput{}
	err = faker.FakeObject(&cd)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeCACertificate(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&cd, nil)

	ct := iot.ListCertificatesByCAOutput{}
	err = faker.FakeObject(&ct)
	if err != nil {
		t.Fatal(err)
	}
	ct.NextMarker = nil
	m.EXPECT().ListCertificatesByCA(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&ct, nil)

	return aws_client.AwsServices{
		IOT: m,
	}
}

func TestIotCaCertificates(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIotCaCertificatesGenerator{}), buildIotCaCertificatesMock, aws_client.TestOptions{})
}
