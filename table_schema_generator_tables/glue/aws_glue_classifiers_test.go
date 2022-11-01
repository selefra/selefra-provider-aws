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

func buildClassifiers(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockGlueClient(ctrl)

	var c glue.GetClassifiersOutput
	if err := faker.FakeObject(&c); err != nil {
		t.Fatal(err)
	}
	c.NextToken = nil
	m.EXPECT().GetClassifiers(gomock.Any(), gomock.Any()).AnyTimes().Return(&c, nil)

	return aws_client.AwsServices{
		Glue: m,
	}
}

func TestClassifiers(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsGlueClassifiersGenerator{}), buildClassifiers, aws_client.TestOptions{})
}
