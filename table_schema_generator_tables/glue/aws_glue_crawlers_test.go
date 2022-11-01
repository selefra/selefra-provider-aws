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

func buildCrawlers(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlueClient(ctrl)

	var crawler glue.GetCrawlersOutput
	if err := faker.FakeObject(&crawler); err != nil {
		t.Fatal(err)
	}
	crawler.NextToken = nil
	m.EXPECT().GetCrawlers(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&crawler, nil)

	var tags glue.GetTagsOutput
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetTags(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(&tags, nil)

	return aws_client.AwsServices{
		Glue: m,
	}
}

func TestCrawlers(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlueCrawlersGenerator{}), buildCrawlers, aws_client.TestOptions{})
}
