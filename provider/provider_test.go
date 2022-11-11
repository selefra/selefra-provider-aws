package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/selefra/selefra-provider-sdk/env"
	"github.com/selefra/selefra-provider-sdk/grpc/shard"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/storage/database_storage/postgresql_storage"
	"github.com/selefra/selefra-utils/pkg/json_util"
	"github.com/selefra/selefra-utils/pkg/pointer"
)

func TestProvider_PullTable(t *testing.T) {

	fmt.Println(env.GetDatabaseDsn())

	wk := "."
	config := `providers:
    # provider configurations
    - name: aws
      #  Optional, Repeated. Add an accounts block for every account you want to assume-role into and fetch data from.
      accounts:
        #     Optional. User identification
        - account_name: t1
          #    Optional. Named profile in config or credential file from where Selefra should grab credentials
      #    shared_credentials_files: "/Users/songzhibin/go/src/aqgs/selefra-provider-aws/config.ini"
      #    shared_config_profile: "shubo"
      #    The maximum number of times that a request will be retried for failures. Defaults to 10 retry attempts.
      max_attempts: 10
      #    The maximum back off delay between attempts. The backoff delays exponentially with a jitter based on the number of attempts. Defaults to 30 seconds.
      max_backoff: 30
`
	myProvider := GetProvider()

	Pull(myProvider, config, wk, "aws_iam_credential_reports")

}

func Test(t *testing.T) {
	myProvider := GetProvider()

	for _, table := range myProvider.TableList {
		fmt.Println(table.TableName)
	}
}

func Pull(myProvider *provider.Provider, config, workspace string, pullTables ...string) {

	diagnostics := schema.NewDiagnostics()

	initProviderRequest := &shard.ProviderInitRequest{
		Storage: &shard.Storage{
			Type:           0,
			StorageOptions: json_util.ToJsonBytes(postgresql_storage.NewPostgresqlStorageOptions(env.GetDatabaseDsn())),
		},
		Workspace:      &workspace,
		IsInstallInit:  pointer.TruePointer(),
		ProviderConfig: &config,
	}

	response, err := myProvider.Init(context.Background(), initProviderRequest)
	if err != nil {
		panic(diagnostics.AddFatal("init provider error: %s", err.Error()).ToString())
	}
	if diagnostics.AddDiagnostics(response.Diagnostics).HasError() {
		panic(diagnostics.ToString())
	}

	err = myProvider.PullTables(context.Background(), &shard.PullTablesRequest{
		Tables:        pullTables,
		MaxGoroutines: 1000,
		Timeout:       1000 * 60 * 60,
	}, shard.NewFakeProviderServerSender())
	if err != nil {
		panic(diagnostics.AddFatal("provider pull table error: %s", err.Error()).ToString())
	}
}
