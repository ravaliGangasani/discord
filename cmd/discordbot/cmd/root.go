package cmd

import (
	"fmt"
	"github.com/discord_login/keys"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)



const (
	flagLogLevel  = "log-level"
	flagLogFormat = "log-format"

	logFormatJson = "json"
	logFormatText = "text"
)

func RootCmd() *cobra.Command {

	rootCmd := &cobra.Command{
		Use: keys.AppName,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat(keys.DataDir); os.IsNotExist(err) {
				_ = os.MkdirAll(keys.DataDir, os.ModePerm)
			}

			return setupLogging(cmd)
		},

		Short: "Discord bot to interact with a Autonomy chain to perform specific transactions using chat commands",
	}

	rootCmd.AddCommand(
		StartCmd(),
	)

	rootCmd.PersistentFlags().String(flagLogLevel, zerolog.DebugLevel.String(), "logging level to be used")
	rootCmd.PersistentFlags().String(flagLogFormat, logFormatText, "logging format; must be either json or text")

	return rootCmd
}

// setupLogging setups the logging for the entire project
func setupLogging(cmd *cobra.Command) error {
	// Init logging level
	value, err := cmd.Flags().GetString(flagLogLevel)
	if err != nil {
		return err
	}

	logLvl, err := zerolog.ParseLevel(value)
	if err != nil {
		return err
	}
	zerolog.SetGlobalLevel(logLvl)

	// Init logging format
	value, err = cmd.Flags().GetString(flagLogFormat)
	if err != nil {
		return err
	}
	switch value {
	case logFormatJson:
		return nil

	case logFormatText:
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		return nil

	default:
		return fmt.Errorf("invalid logging format: %s", value)
	}
}