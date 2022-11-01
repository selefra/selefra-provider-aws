package mq

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildMqBrokers(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockMQClient(ctrl)

	bs := types.BrokerSummary{}
	if err := faker.FakeObject(&bs); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListBrokers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&mq.ListBrokersOutput{
			BrokerSummaries: []types.BrokerSummary{bs},
		}, nil)

	bo := mq.DescribeBrokerOutput{}
	if err := faker.FakeObject(&bo); err != nil {
		t.Fatal(err)
	}
	bo.BrokerId = bs.BrokerId
	username := "test_username"
	bo.Users = []types.UserSummary{{Username: &username}}
	var cfgID types.ConfigurationId
	if err := faker.FakeObject(&cfgID); err != nil {
		t.Fatal(err)
	}
	bo.Configurations.Current = &cfgID
	bo.Configurations.History = []types.ConfigurationId{cfgID}
	m.EXPECT().DescribeBroker(gomock.Any(), &mq.DescribeBrokerInput{BrokerId: bs.BrokerId}, gomock.Any()).AnyTimes().Return(&bo, nil)

	uo := mq.DescribeUserOutput{}
	if err := faker.FakeObject(&uo); err != nil {
		t.Fatal(err)
	}
	uo.Username = &username
	uo.BrokerId = bo.BrokerId
	m.EXPECT().DescribeUser(gomock.Any(), &mq.DescribeUserInput{BrokerId: bo.BrokerId, Username: &username}, gomock.Any()).AnyTimes().Return(&uo, nil)

	var co mq.DescribeConfigurationOutput
	if err := faker.FakeObject(&co); err != nil {
		t.Fatal(err)
	}
	co.Id = cfgID.Id
	m.EXPECT().DescribeConfiguration(gomock.Any(), &mq.DescribeConfigurationInput{ConfigurationId: cfgID.Id}, gomock.Any()).AnyTimes().Return(&co, nil)

	revisions := mq.ListConfigurationRevisionsOutput{}
	if err := faker.FakeObject(&revisions); err != nil {
		t.Fatal(err)
	}
	revisions.NextToken = nil
	m.EXPECT().ListConfigurationRevisions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&revisions, nil)

	revision := mq.DescribeConfigurationRevisionOutput{}
	if err := faker.FakeObject(&revision); err != nil {
		t.Fatal(err)
	}
	revision.Data = aws.String("PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz48aGVsbG8+d29ybGQ8L2hlbGxvPg==")
	m.EXPECT().DescribeConfigurationRevision(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&revision, nil)

	return aws_client.AwsServices{MQ: m}
}

func TestMqBrokers(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsMqBrokersGenerator{}), buildMqBrokers, aws_client.TestOptions{})
}
