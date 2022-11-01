package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildIamVirtualMfaDevices(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.VirtualMFADevice{}
	err := faker.FakeObject(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListVirtualMFADevices(gomock.Any(), gomock.Any()).AnyTimes().Return(
		&iam.ListVirtualMFADevicesOutput{
			VirtualMFADevices: []iamTypes.VirtualMFADevice{g},
		}, nil)
	return aws_client.AwsServices{
		IAM: m,
	}
}

func TestIAMVirtualMfaDevices(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsIamVirtualMfaDevicesGenerator{}), buildIamVirtualMfaDevices, aws_client.TestOptions{})
}
