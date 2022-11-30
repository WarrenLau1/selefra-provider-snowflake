package provider

import (
	"context"
	"os"
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
	os.Setenv("SELEFRA_DATABASE_DSN", "host=127.0.0.1 user=postgres password=password port=5432 dbname=postgres sslmode=disable")
	wk := "."
	config := `
`
	myProvider := GetProvider()
	Pull(myProvider, config, wk, "*")
}

func Pull(myProvider *provider.Provider, config, workspace string, pullTables ...string) {

	diagnostics := schema.NewDiagnostics()

	// init Provider
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
		MaxGoroutines: 100,
		Timeout:       1000 * 60 * 60,
	}, shard.NewFakeProviderServerSender())
	if err != nil {
		panic(diagnostics.AddFatal("provider pull table error: %s", err.Error()).ToString())
	}
}
