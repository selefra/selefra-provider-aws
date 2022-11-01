package s3

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildS3Buckets(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	mgr := mocks.NewMockS3ManagerClient(ctrl)
	m := mocks.NewMockS3Client(ctrl)
	b := s3Types.Bucket{}
	err := faker.FakeObject(&b)
	if err != nil {
		t.Fatal(err)
	}
	bloc := s3.GetBucketLocationOutput{}
	err = faker.FakeObject(&bloc)
	if err != nil {
		t.Fatal(err)
	}
	blog := s3.GetBucketLoggingOutput{}
	err = faker.FakeObject(&blog)
	if err != nil {
		t.Fatal(err)
	}
	bpol := s3.GetBucketPolicyOutput{}
	err = faker.FakeObject(&bpol)
	if err != nil {
		t.Fatal(err)
	}
	jsonDoc := `{"stuff": 3}`
	bpol.Policy = &jsonDoc
	bver := s3.GetBucketVersioningOutput{}
	err = faker.FakeObject(&bver)
	if err != nil {
		t.Fatal(err)
	}
	bgrant := s3Types.Grant{}
	err = faker.FakeObject(&bgrant)
	if err != nil {
		t.Fatal(err)
	}
	bcors := s3Types.CORSRule{}
	err = faker.FakeObject(&bcors)
	if err != nil {
		t.Fatal(err)
	}
	bencryption := s3.GetBucketEncryptionOutput{}
	err = faker.FakeObject(&bencryption)
	if err != nil {
		t.Fatal(err)
	}

	bpba := s3.GetPublicAccessBlockOutput{}
	err = faker.FakeObject(&bpba)
	if err != nil {
		t.Fatal(err)
	}
	btag := s3.GetBucketTaggingOutput{}
	err = faker.FakeObject(&btag)
	if err != nil {
		t.Fatal(err)
	}
	bownershipcontrols := s3.GetBucketOwnershipControlsOutput{}
	err = faker.FakeObject(&bownershipcontrols)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListBuckets(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&s3.ListBucketsOutput{
			Buckets: []s3Types.Bucket{b},
		}, nil)
	m.EXPECT().GetBucketLogging(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&blog, nil)
	m.EXPECT().GetBucketPolicy(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&bpol, nil)
	m.EXPECT().GetBucketVersioning(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&bver, nil)
	m.EXPECT().GetBucketAcl(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&s3.GetBucketAclOutput{
			Grants: []s3Types.Grant{bgrant},
		}, nil)
	m.EXPECT().GetBucketCors(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&s3.GetBucketCorsOutput{
			CORSRules: []s3Types.CORSRule{bcors},
		}, nil)
	m.EXPECT().GetBucketEncryption(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&bencryption, nil)

	m.EXPECT().GetPublicAccessBlock(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&bpba, nil)
	m.EXPECT().GetBucketOwnershipControls(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&bownershipcontrols, nil)

	ro := s3.GetBucketReplicationOutput{}
	if err := faker.FakeObject(&ro); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetBucketReplication(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&ro, nil)
	m.EXPECT().GetBucketTagging(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&btag, nil)
	tt := s3Types.Transition{}
	if err := faker.FakeObject(&tt); err != nil {
		t.Fatal(err)
	}
	glco := s3.GetBucketLifecycleConfigurationOutput{}
	if err := faker.FakeObject(&glco); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetBucketLifecycleConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&glco, nil)
	mgr.EXPECT().GetBucketRegion(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		"us-east-1", nil)
	return aws_client.AwsServices{
		S3:		m,
		S3Manager:	mgr,
	}
}

func TestS3Buckets(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsS3BucketsGenerator{}), buildS3Buckets, aws_client.TestOptions{})
}
