package fsx

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildVolumesMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockFsxClient(ctrl)

	var v types.Volume
	if err := faker.FakeObject(&v); err != nil {
		t.Fatal(err)
	}
	v.Lifecycle = types.VolumeLifecycleAvailable
	v.VolumeType = types.VolumeTypeOntap
	m.EXPECT().DescribeVolumes(
		gomock.Any(),
		&fsx.DescribeVolumesInput{MaxResults: aws.Int32(1000)},
	).AnyTimes().Return(
		&fsx.DescribeVolumesOutput{Volumes: []types.Volume{v}},
		nil,
	)

	return aws_client.AwsServices{
		FSX: m,
	}
}

func TestVolumes(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsFsxVolumesGenerator{}), buildVolumesMock, aws_client.TestOptions{})
}
