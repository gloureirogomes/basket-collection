package cmd

import (
	"github.com/GabrielLoureiroGomes/basket-collection/api"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "A command to run the API",
	Run: func(_ *cobra.Command, _ []string) {
		server := api.NewServer()
		server.StartServer()
	},
}
