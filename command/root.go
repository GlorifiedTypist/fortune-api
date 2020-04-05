package command

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "fortune-api",
	Short: "*NIX Fortune API",
	Long: `Returns a *NIX fortune via RESTful API.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
