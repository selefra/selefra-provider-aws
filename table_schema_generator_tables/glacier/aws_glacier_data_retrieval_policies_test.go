package glacier

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func buildDRPMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlacierClient(ctrl)

	p := glacier.GetDataRetrievalPolicyOutput{}
	require.NoError(t, faker.FakeObject(&p))
	m.EXPECT().GetDataRetrievalPolicy(gomock.Any(), gomock.Any()).AnyTimes().Return(&p, nil)

	return aws_client.AwsServices{
		Glacier: m,
	}
}

func TestDataRetrievalPolicies(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlacierDataRetrievalPoliciesGenerator{}), buildDRPMock, aws_client.TestOptions{})
}
