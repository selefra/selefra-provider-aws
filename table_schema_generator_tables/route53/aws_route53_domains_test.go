package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildRoute53Domains(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockRoute53DomainsClient(ctrl)

	var ds types.DomainSummary
	if err := faker.FakeObject(&ds.DomainName); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListDomains(gomock.Any(), &route53domains.ListDomainsInput{}, gomock.Any()).AnyTimes().Return(
		&route53domains.ListDomainsOutput{Domains: []types.DomainSummary{ds}},
		nil,
	)

	var detail route53domains.GetDomainDetailOutput
	if err := faker.FakeObject(&detail); err != nil {
		t.Fatal(err)
	}
	detail.DomainName = ds.DomainName
	mock.EXPECT().GetDomainDetail(gomock.Any(), &route53domains.GetDomainDetailInput{DomainName: ds.DomainName}, gomock.Any()).AnyTimes().Return(
		&detail, nil,
	)

	var tagsOut route53domains.ListTagsForDomainOutput
	if err := faker.FakeObject(&tagsOut); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListTagsForDomain(gomock.Any(), &route53domains.ListTagsForDomainInput{DomainName: ds.DomainName}, gomock.Any()).AnyTimes().Return(
		&tagsOut, nil,
	)

	return aws_client.AwsServices{
		Route53Domains: mock,
	}
}

func TestRoute53Domains(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRoute53DomainsGenerator{}), buildRoute53Domains, aws_client.TestOptions{})
}
