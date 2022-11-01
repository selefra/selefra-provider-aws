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

func buildDirectconnectVirtualInterfacesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.VirtualInterface{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVirtualInterfaces(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&directconnect.DescribeVirtualInterfacesOutput{
			VirtualInterfaces: []types.VirtualInterface{l},
		}, nil)
	return aws_client.AwsServices{
		Directconnect: m,
	}
}

func TestDirectconnecVirtualInterfaces(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsDirectconnectVirtualInterfacesGenerator{}), buildDirectconnectVirtualInterfacesMock, aws_client.TestOptions{})
}
