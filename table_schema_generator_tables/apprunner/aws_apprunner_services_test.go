package apprunner

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/apprunner"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildApprunnerGraphqlApisMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockAppRunnerClient(ctrl)
	s := types.Service{}
	err := faker.FakeObject(&s)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListServices(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apprunner.ListServicesOutput{
			ServiceSummaryList: []types.ServiceSummary{
				{ServiceArn: s.ServiceArn},
			},
		}, nil)

	m.EXPECT().DescribeService(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&apprunner.DescribeServiceOutput{
			Service: &s,
		}, nil)

	return aws_client.AwsServices{
		Apprunner: m,
	}
}

func TestAppSyncGraphqlApis(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsApprunnerServicesGenerator{}), buildApprunnerGraphqlApisMock, aws_client.TestOptions{})
}
