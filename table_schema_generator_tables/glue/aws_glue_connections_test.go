package glue

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildConnections(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlueClient(ctrl)

	var connecions glue.GetConnectionsOutput
	if err := faker.FakeObject(&connecions); err != nil {
		t.Fatal(err)
	}
	connecions.NextToken = nil
	m.EXPECT().GetConnections(gomock.Any(), gomock.Any()).AnyTimes().Return(&connecions, nil)

	return aws_client.AwsServices{
		Glue: m,
	}
}

func TestConnections(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlueConnectionsGenerator{}), buildConnections, aws_client.TestOptions{})
}
