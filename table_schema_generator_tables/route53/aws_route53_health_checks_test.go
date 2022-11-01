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

func buildRoute53HealthChecksMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockRoute53Client(ctrl)
	hc := route53Types.HealthCheck{}
	if err := faker.FakeObject(&hc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListHealthChecks(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&route53.ListHealthChecksOutput{
			HealthChecks: []route53Types.HealthCheck{hc},
		}, nil)
	tag := route53Types.Tag{}
	if err := faker.FakeObject(&tag); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListTagsForResources(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&route53.ListTagsForResourcesOutput{
			ResourceTagSets: []route53Types.ResourceTagSet{
				{
					ResourceId:	hc.Id,
					Tags:		[]route53Types.Tag{tag},
				},
			},
		}, nil)
	return aws_client.AwsServices{
		Route53: m,
	}
}

func TestRoute53HealthCheck(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRoute53HealthChecksGenerator{}), buildRoute53HealthChecksMock, aws_client.TestOptions{})
}
