package cli

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	rootCmd = &cobra.Command{
		Use:   "url shortener",
		Short: "url shortener service.",
		Long:  ` url shortener.`,
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "url-shortener-configs", "config file")
	rootCmd.PersistentFlags().StringP("author", "a", "Ali Sadafi", "bale@gmail.com")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(setupConfig)
}
