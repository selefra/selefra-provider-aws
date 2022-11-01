package lambda

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-aws/aws_client/mocks"
	"github.com/selefra/selefra-provider-aws/faker"
	"github.com/selefra/selefra-provider-aws/table_schema_generator"
)

func buildLambdaFunctionsMock(t *testing.T, ctrl *gomock.Controller) aws_client.AwsServices {
	m := mocks.NewMockLambdaClient(ctrl)

	lastModified := "1994-11-05T08:15:30.000+0500"

	f := lambda.GetFunctionOutput{}
	err := faker.FakeObject(&f)
	if err != nil {
		t.Fatal(err)
	}
	f.Configuration.LastModified = &lastModified
	m.EXPECT().GetFunction(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&f, nil)

	fc := types.FunctionConfiguration{}
	err = faker.FakeObject(&fc)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListFunctions(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lambda.ListFunctionsOutput{
			Functions: []types.FunctionConfiguration{fc},
		}, nil)

	a := types.AliasConfiguration{}
	err = faker.FakeObject(&a)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListAliases(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lambda.ListAliasesOutput{
			Aliases: []types.AliasConfiguration{a},
		}, nil)

	i := types.FunctionEventInvokeConfig{}
	err = faker.FakeObject(&i)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListFunctionEventInvokeConfigs(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lambda.ListFunctionEventInvokeConfigsOutput{
			FunctionEventInvokeConfigs: []types.FunctionEventInvokeConfig{i},
		}, nil)

	cc := types.ProvisionedConcurrencyConfigListItem{}
	err = faker.FakeObject(&cc)
	if err != nil {
		t.Fatal(err)
	}
	cc.LastModified = &lastModified
	m.EXPECT().ListProvisionedConcurrencyConfigs(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lambda.ListProvisionedConcurrencyConfigsOutput{
			ProvisionedConcurrencyConfigs: []types.ProvisionedConcurrencyConfigListItem{cc},
		}, nil)

	esm := types.EventSourceMappingConfiguration{}
	err = faker.FakeObject(&esm)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListEventSourceMappings(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lambda.ListEventSourceMappingsOutput{
			EventSourceMappings: []types.EventSourceMappingConfiguration{esm},
		}, nil)

	fp := lambda.GetPolicyOutput{}
	err = faker.FakeObject(&fp)
	if err != nil {
		t.Fatal(err)
	}
	document := "{\"test\":1}"
	fp.Policy = &document
	m.EXPECT().GetPolicy(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&fp, nil)

	csco := lambda.GetFunctionCodeSigningConfigOutput{}
	err = faker.FakeObject(&csco)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetFunctionCodeSigningConfig(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&csco, nil)

	csc := types.CodeSigningConfig{}
	err = faker.FakeObject(&csc)
	if err != nil {
		t.Fatal(err)
	}
	isoDate := "2011-10-05T14:48:00.000Z"
	csc.LastModified = &isoDate
	m.EXPECT().GetCodeSigningConfig(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lambda.GetCodeSigningConfigOutput{
			CodeSigningConfig: &csc,
		}, nil)
	fc.LastModified = &lastModified
	m.EXPECT().ListVersionsByFunction(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&lambda.ListVersionsByFunctionOutput{
			Versions: []types.FunctionConfiguration{fc},
		}, nil)

	urlConfig := lambda.GetFunctionUrlConfigOutput{}
	err = faker.FakeObject(&urlConfig)
	if err != nil {
		t.Fatal(err)
	}
	urlConfig.CreationTime = aws.String("2012-07-14T01:00:00+01:00")
	urlConfig.LastModifiedTime = aws.String("2012-07-14T01:00:00+01:00")
	m.EXPECT().GetFunctionUrlConfig(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().Return(
		&urlConfig, nil)

	return aws_client.AwsServices{
		Lambda: m,
	}
}

func TestLambdaFunctions(t *testing.T) {
	aws_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAwsLambdaFunctionsGenerator{}), buildLambdaFunctionsMock, aws_client.TestOptions{})
}
