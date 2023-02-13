package main

import (
	"os"

	"github.com/SkyMack/clibase"
	"github.com/SkyMack/staledesk/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	appName        = "staledesk"
	appDescription = "A Freshdesk compatible REST API"

	flagPrefix = "STALE_"
)

var (
	rootCmd = &cobra.Command{
		Use:   appName,
		Short: appDescription,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var (
				conf *config.Data
				err  error
			)
			conf, err = config.GenerateConfigData()
			if err != nil {
				return err
			}
			config.SetConfig(conf)
			return nil
		},
	}
)

func init() {
}

func main() {
	//rootCmd := clibase.New(appName, appDescription)
	rootCmd := clibase.NewUsingCmd(rootCmd)
	config.AddConfigCmd(rootCmd)
	addServeCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("server exited with an error")
		os.Exit(1)
	}
}
