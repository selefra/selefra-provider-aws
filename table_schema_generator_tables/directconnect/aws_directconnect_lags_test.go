package directconnect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildDirectconnectLag(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockDirectconnectClient(ctrl)
	lag := types.Lag{}
	err := faker.FakeObject(&lag)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeLags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&directconnect.DescribeLagsOutput{
			Lags: []types.Lag{lag},
		}, nil)
	return aws_client.AwsServices{
		Directconnect: m,
	}
}

func TestDirectconnectLag(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDirectconnectLagsGenerator{}), buildDirectconnectLag, aws_client.TestOptions{})
}
