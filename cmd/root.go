package cmd

import (
	"example/config"
	"fmt"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var Verbose bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "example",
	Short: "This is a short description",
	Long:  `This is a longer description.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := envconfig.Process("", &config.AppConfiguration)
		if err != nil {
			log.Fatal().Err(err).Msg("error starting app config is missing")
		}

		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		if Verbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
		output := zerolog.ConsoleWriter{Out: os.Stderr}
		output.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		}
		log.Logger = log.Output(output)
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Enable verbose logging")
}
