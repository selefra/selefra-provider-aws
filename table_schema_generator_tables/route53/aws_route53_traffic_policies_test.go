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

func buildRoute53TrafficPoliciesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockRoute53Client(ctrl)
	tps := route53Types.TrafficPolicySummary{}
	if err := faker.FakeObject(&tps); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTrafficPolicies(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&route53.ListTrafficPoliciesOutput{
			TrafficPolicySummaries: []route53Types.TrafficPolicySummary{tps},
		}, nil)
	tp := route53Types.TrafficPolicy{}
	if err := faker.FakeObject(&tp); err != nil {
		t.Fatal(err)
	}
	tp.Id = tps.Id
	document := `{"AWSPolicyFormatVersion":"2015-10-01","RecordType":"A","Endpoints":{"endpoint-geoproximity-vfcf":{"Type":"value","Value":"1.0.0.1"},"endpoint-geoproximity-gPSy":{"Type":"value","Value":"1.0.0.2"}},"Rules":{"geoproximity-start-Hfni":{"RuleType":"geoproximity","GeoproximityLocations":[{"EndpointReference":"endpoint-geoproximity-vfcf","Bias":0,"Region":"aws:route53:us-east-1","EvaluateTargetHealth":true},{"Bias":0,"Region":"aws:route53:us-east-2","EvaluateTargetHealth":true,"EndpointReference":"endpoint-geoproximity-gPSy"}]}},"StartRule":"geoproximity-start-Hfni"}`
	tp.Document = &document
	m.EXPECT().ListTrafficPolicyVersions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&route53.ListTrafficPolicyVersionsOutput{
			TrafficPolicies: []route53Types.TrafficPolicy{tp},
		}, nil)
	return aws_client.AwsServices{
		Route53: m,
	}
}

func TestRoute53TrafficPolicies(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRoute53TrafficPoliciesGenerator{}), buildRoute53TrafficPoliciesMock, aws_client.TestOptions{})
}
