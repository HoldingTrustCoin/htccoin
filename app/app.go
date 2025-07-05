package app

import (
	wasm "github.com/CosmWasm/wasmd/x/wasm"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp"
)

// NewHTCCoinApp cria a aplicação da blockchain HTC Coin
func NewHTCCoinApp() *baseapp.BaseApp {
	// Aqui estamos usando simapp apenas como exemplo base
	app := simapp.NewSimApp(
		nil,              // logger
		nil,              // db
		nil,              // traceStore
		true,             // loadLatest
		map[int64]bool{}, // skipUpgradeHeights
		"",               // homePath
		0,                // invCheckPeriod
		simapp.MakeEncodingConfig(),
		simapp.EmptyAppOptions{},
	)

	// Registra o módulo CosmWasm
	wasmModule := wasm.NewAppModule(app.AppCodec(), &app.WasmKeeper, app.AccountKeeper, app.BankKeeper)
	app.ModuleManager.RegisterModules(wasmModule)

	return app.BaseApp
}
