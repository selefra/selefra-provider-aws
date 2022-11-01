package lambda

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildLambdaRuntimesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockLambdaClient(ctrl)
	return aws_client.AwsServices{
		Lambda: m,
	}
}

func TestLambdaRuntimes(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLambdaRuntimesGenerator{}), buildLambdaRuntimesMock, aws_client.TestOptions{})
}
