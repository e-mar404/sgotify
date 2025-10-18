package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/internal/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Verbose bool 
	Debug bool

	rootCmd = &cobra.Command{
		Use: "sgotify",
		Short: "Spotify for the cli/tui",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("starting tui")

			if err := tui.Run(); err != nil {
				log.Error("could not run tui", "error", err)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initRoot)
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "will show all logs except debug")
	rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "will show all logs")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initRoot() {
	switch {
		case Verbose:
			// will show all logs except debug
			log.SetLevel(log.Level(0))
		case Debug:
			// debug level is -4
			log.SetLevel(log.Level(-5))
		default:
			// highest level is 12, so this will hide all logs
			log.SetLevel(log.Level(13))
	}
	
	home, err := os.UserHomeDir()
	if err != nil {
		log.Error("could not get home dir", "error", err)
		cobra.CheckErr(err)
	}

	viper.AddConfigPath(home)
	viper.SetConfigType("json")
	viper.SetConfigName(".sgotify.json")
	
	if err := viper.ReadInConfig(); err != nil {
		log.Error("could not read config file", "error", err)
		cobra.CheckErr(err)
	}

	log.Info("config file loaded", "path", viper.ConfigFileUsed())
}
