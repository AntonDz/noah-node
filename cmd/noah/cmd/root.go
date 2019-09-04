package cmd

import (
	"github.com/noah-blockchain/noah-node/cmd/utils"
	"github.com/noah-blockchain/noah-node/config"
	"github.com/noah-blockchain/noah-node/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg *config.Config

var RootCmd = &cobra.Command{
	Use:   "noah",
	Short: "Noah Go Node",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		v := viper.New()
		v.SetConfigFile(utils.GetNoahConfigPath())
		cfg = config.GetConfig()

		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := v.Unmarshal(cfg); err != nil {
			panic(err)
		}

		log.InitLog(cfg)
	},
}
