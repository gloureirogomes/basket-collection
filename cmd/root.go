package cmd

import (
	"github.com/GabrielLoureiroGomes/basket-collection/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "A main command to start API",
	Long:  "Command used to start the basket collection API",
}

// Execute executes the root command.
func Execute() error {
	cobra.OnInitialize(config.InitConfig)
	rootCmd.AddCommand(apiCmd)

	return rootCmd.Execute()
}
