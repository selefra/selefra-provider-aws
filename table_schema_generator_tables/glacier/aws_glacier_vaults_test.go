package glacier

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func buildVaultsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlacierClient(ctrl)

	v := glacier.ListVaultsOutput{}
	require.NoError(t, faker.FakeObject(&v))
	v.Marker = nil
	m.EXPECT().ListVaults(gomock.Any(), gomock.Any()).AnyTimes().Return(&v, nil)

	ap := glacier.GetVaultAccessPolicyOutput{}
	require.NoError(t, faker.FakeObject(&ap))
	ap.Policy.Policy = aws.String(`{"some":"policy"}`)
	m.EXPECT().GetVaultAccessPolicy(gomock.Any(), gomock.Any()).AnyTimes().Return(&ap, nil)

	lp := glacier.GetVaultLockOutput{}
	require.NoError(t, faker.FakeObject(&lp))
	lp.Policy = aws.String(`{"some":"policy"}`)
	m.EXPECT().GetVaultLock(gomock.Any(), gomock.Any()).AnyTimes().Return(&lp, nil)

	vn := glacier.GetVaultNotificationsOutput{}
	require.NoError(t, faker.FakeObject(&vn))
	m.EXPECT().GetVaultNotifications(gomock.Any(), gomock.Any()).AnyTimes().Return(&vn, nil)

	tags := glacier.ListTagsForVaultOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	m.EXPECT().ListTagsForVault(gomock.Any(), gomock.Any()).AnyTimes().Return(&tags, nil)

	return aws_client.AwsServices{
		Glacier: m,
	}
}

func TestVaults(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlacierVaultsGenerator{}), buildVaultsMock, aws_client.TestOptions{})
}
