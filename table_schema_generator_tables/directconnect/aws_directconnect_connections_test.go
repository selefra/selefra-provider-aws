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

func buildDirectconnectConnection(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockDirectconnectClient(ctrl)
	conn := types.Connection{}
	err := faker.FakeObject(&conn)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConnections(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&directconnect.DescribeConnectionsOutput{
			Connections: []types.Connection{conn},
		}, nil)
	return aws_client.AwsServices{
		Directconnect: m,
	}
}

func TestDirectconnectConnection(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDirectconnectConnectionsGenerator{}), buildDirectconnectConnection, aws_client.TestOptions{})
}
