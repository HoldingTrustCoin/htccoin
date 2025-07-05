package main

import (
	"github.com/spf13/cobra"

	"github.com/HoldingTrustCoin/htccoin/app"
	"github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/types/module"
)

// NewRootCmd cria o comando raiz para htcd
func NewRootCmd() (*cobra.Command, module.BasicManager) {
	encodingConfig := app.MakeEncodingConfig()

	rootCmd := &cobra.Command{
		Use:   "htcd",
		Short: "HTC Coin Daemon (full-node) da blockchain",
	}

	initRootCmd(rootCmd, encodingConfig)

	return rootCmd, app.ModuleBasics
}

// initRootCmd adiciona subcomandos ao rootCmd
func initRootCmd(rootCmd *cobra.Command, encodingConfig app.EncodingConfig) {
	// Adiciona comandos padrões do Cosmos SDK
	rootCmd.AddCommand(
		cmd.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		cmd.CollectGenTxsCmd(app.DefaultNodeHome),
		cmd.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, app.DefaultNodeHome),
		cmd.ValidateGenesisCmd(app.ModuleBasics),
		cmd.AddGenesisAccountCmd(app.DefaultNodeHome),
		cmd.NewCompletionCmd(rootCmd, true),
	)
}
