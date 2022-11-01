package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53Types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildRoute53DelegationSetsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockRoute53Client(ctrl)
	ds := route53Types.DelegationSet{}
	if err := faker.FakeObject(&ds); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListReusableDelegationSets(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&route53.ListReusableDelegationSetsOutput{
			DelegationSets: []route53Types.DelegationSet{ds},
		}, nil)
	return aws_client.AwsServices{
		Route53: m,
	}
}

func TestRoute53DelegationSets(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRoute53DelegationSetsGenerator{}), buildRoute53DelegationSetsMock, aws_client.TestOptions{})
}
