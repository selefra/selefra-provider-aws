package provider

import (
	"context"

	"github.com/selefra/selefra-provider-aws/aws_client"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

const Version = "v0.0.5"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "aws",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var awsConfig aws_client.AwsProviderConfig
				err := config.Unmarshal(&awsConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				clients, err := aws_client.NewClients(awsConfig)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("account information not found")
				}

				hash := make(map[string]bool)
				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					if hash[clients[i].GetAccount()] {
						continue
					}
					res = append(res, clients[i])
					hash[clients[i].GetAccount()] = true
				}
				return res, nil
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
				var awsConfig aws_client.AwsProviderConfig
				err := config.Unmarshal(&awsConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorAll},
		},
	}
}
