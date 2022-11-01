package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildParameters(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mock := mocks.NewMockSSMClient(ctrl)
	var pm types.ParameterMetadata
	if err := faker.FakeObject(&pm); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeParameters(
		gomock.Any(),
		&ssm.DescribeParametersInput{},
	).AnyTimes().Return(
		&ssm.DescribeParametersOutput{Parameters: []types.ParameterMetadata{pm}},
		nil,
	)
	return aws_client.AwsServices{SSM: mock}
}

func TestParameters(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsSsmParametersGenerator{}), buildParameters, aws_client.TestOptions{})
}
