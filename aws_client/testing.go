package aws_client

import (
	"context"
	"github.com/selefra/selefra-provider-sdk/test_helper"
	"github.com/spf13/viper"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TestOptions struct{}

func MockTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) AwsServices, _ TestOptions) {
	ctrl := gomock.NewController(t)
	testProvider := newTestProvider(t, ctrl, table, builder)
	config := "test : test"
	test_helper.RunProviderPullTables(testProvider, config, "./", "*")
}

func newTestProvider(t *testing.T, ctrl *gomock.Controller, table *schema.Table, builder func(*testing.T, *gomock.Controller) AwsServices) *provider.Provider {
	return &provider.Provider{
		Name:		"aws",
		Version:	"v0.0.1",
		TableList:	[]*schema.Table{table},
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {

				partition := "selefra-mock-aws"
				region := "us-east-1"
				awsServices := builder(t, ctrl)
				accountAwsServiceManager := NewAwsServiceCache(nil)
				accountAwsServiceManager.AwsServicesManagerMap[partition] = make(map[string]*AwsServices, 0)
				accountAwsServiceManager.AwsServicesManagerMap[partition][region] = &awsServices
				client := &Client{
					Partition:			partition,
					AccountID:			"testAccount",
					Region:				region,
					accountAwsServiceManager:	accountAwsServiceManager,
				}
				return []any{client}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `##  Optional, Repeated. Add an accounts block for every account you want to assume-role into and fetch data from.
#accounts:
#    #     Optional. User identification
#  - account_name: <UNIQUE ACCOUNT IDENTIFIER>
#    #    Optional. Named profile in config or credential file from where Selefra should grab credentials
#    shared_config_profile: < PROFILE_NAME >
#    #    Optional. Location of shared configuration files
#    shared_config_files:
#      - <FILE_PATH>
#    #   Optional. Location of shared credentials files
#    shared_credentials_files:
#      - <FILE_PATH>
#    #    Optional. Role ARN we want to assume when accessing this account
#    role_arn: < YOUR_ROLE_ARN >
#    #    Optional. Named role session to grab specific operation under the assumed role
#    role_session_name: <SESSION_NAME>
#    #    Optional. Any outside of the org account id that has additional control
#    external_id: <ID>
#    #    Optional. Designated region of servers
#    default_region: <REGION_CODE>
#    #    Optional. by default assumes all regions
#    regions:
#      - us-east-1
#      - us-west-2
##    The maximum number of times that a request will be retried for failures. Defaults to 10 retry attempts.
#max_attempts: 10
##    The maximum back off delay between attempts. The backoff delays exponentially with a jitter based on the number of attempts. Defaults to 30 seconds.
#max_backoff: 30`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand:	true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
