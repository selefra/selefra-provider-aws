package route53

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildRoute53HostedZonesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockRoute53Client(ctrl)
	h := types.HostedZone{}
	if err := faker.FakeObject(&h); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListHostedZones(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&route53.ListHostedZonesOutput{
			HostedZones: []types.HostedZone{h},
		}, nil)
	tag := types.Tag{}
	if err := faker.FakeObject(&tag); err != nil {
		t.Fatal(err)
	}

	hzId := *h.Id
	newId := fmt.Sprintf("/%s/%s", types.TagResourceTypeHostedzone, *h.Id)
	h.Id = &newId
	m.EXPECT().ListTagsForResources(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&route53.ListTagsForResourcesOutput{
			ResourceTagSets: []types.ResourceTagSet{
				{
					ResourceId:	&hzId,
					Tags:		[]types.Tag{tag},
				},
			},
		}, nil)
	qlc := types.QueryLoggingConfig{}
	if err := faker.FakeObject(&qlc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListQueryLoggingConfigs(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&route53.ListQueryLoggingConfigsOutput{
			QueryLoggingConfigs: []types.QueryLoggingConfig{qlc},
		}, nil)
	rrs := types.ResourceRecordSet{}
	if err := faker.FakeObject(&rrs); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListResourceRecordSets(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&route53.ListResourceRecordSetsOutput{
			ResourceRecordSets: []types.ResourceRecordSet{rrs},
		}, nil)
	tpi := types.TrafficPolicyInstance{}
	if err := faker.FakeObject(&tpi); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTrafficPolicyInstancesByHostedZone(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&route53.ListTrafficPolicyInstancesByHostedZoneOutput{
			TrafficPolicyInstances: []types.TrafficPolicyInstance{tpi},
		}, nil)
	vpc := types.VPC{}
	if err := faker.FakeObject(&vpc); err != nil {
		t.Fatal(err)
	}
	ds := types.DelegationSet{}
	if err := faker.FakeObject(&ds); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetHostedZone(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&route53.GetHostedZoneOutput{
			HostedZone:	&h,
			DelegationSet:	&ds,
			VPCs:		[]types.VPC{vpc},
		}, nil)
	return aws_client.AwsServices{
		Route53: m,
	}
}

func TestRoute53HostedZones(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsRoute53HostedZonesGenerator{}), buildRoute53HostedZonesMock, aws_client.TestOptions{})
}
