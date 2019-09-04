package main

import (
	cmd2 "github.com/noah-blockchain/noah-node/cmd/noah/cmd"
	"github.com/noah-blockchain/noah-node/cmd/utils"
	"github.com/noah-blockchain/noah-node/config"
)

func main() {

	//box := packr.New("myBox", "testnet/noah-testnet-1")
	//_, err := box.FindString("genesis.json")
	//if err != nil {
	//	panic(err)
	//}
	//panic(err)
	rootCmd := cmd2.RootCmd

	rootCmd.AddCommand(
		cmd2.RunNode,
		cmd2.ShowNodeId,
		cmd2.ShowValidator,
		cmd2.Version,
	)

	rootCmd.PersistentFlags().StringVar(&utils.NoahHome, "home-dir", "", "base dir (default is $HOME/noah)")
	rootCmd.PersistentFlags().StringVar(&utils.NoahConfig, "config", "", "path to config (default is $(home-dir)/config/config.toml)")
	rootCmd.PersistentFlags().StringVar(&config.NetworkId, "network-id", config.DefaultNetworkId, "network id")

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
