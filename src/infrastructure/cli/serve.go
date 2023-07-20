package cli

import (
	"github.com/spf13/cobra"
	"url-shortener/src/infrastructure/server"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serving url shortener service.",
	Long:  `Serving url shortener service.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve() {
	server.RunServer()
}
