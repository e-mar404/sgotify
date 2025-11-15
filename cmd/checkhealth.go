package cmd

import (
	"fmt"
	"net"
	"net/rpc"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	msgpackrpc "github.com/hashicorp/net-rpc-msgpackrpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	checkhealthCmd = &cobra.Command{
		Use:   "checkhealth",
		Short: "verify that all the services and resources are in working condition",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("checking health...\n")

			// represents: map[nameOfTest string]hasItRan bool
			statusChecks := map[string]bool{
				"configuration file":    false,
				"accessible rpc server": false,
				"good login info":       false,
			}

			defer printLeftoverChecks(statusChecks)

			name := "configuration file"
			statusChecks[name] = true
			if correctConfig() {
				fmt.Printf("✓ %s\n", name)
			} else {
				fmt.Printf("x %s\n", name)
				return
			}

			name = "accessible rpc server"
			statusChecks[name] = true
			if accessibleServer() {
				fmt.Printf("✓ %s\n", name)
			} else {
				fmt.Printf("x %s\n", name)
				return
			}

			name = "good login info"
			statusChecks[name] = true
			if goodLoginInfo() {
				fmt.Printf("✓ %s\n", name)
			} else {
				fmt.Printf("x %s\n", name)
				return
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(checkhealthCmd)
}

func correctConfig() bool {
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("could not read/find config file", "error", err)
		return false
	}
	return true
}

func accessibleServer() bool {
	_, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		log.Error("could not dial into server", "error", err)
		return false
	}
	return true
}

func goodLoginInfo() bool {
	conn, err := net.Dial("tcp", "localhost:5000")
	client = rpc.NewClientWithCodec(msgpackrpc.NewClientCodec(conn))
	refreshArgs := api.RefreshArgs{
		RefreshToken: viper.GetString("refresh_token"),
		ClientID:     viper.GetString("client_id"),
		ClientSecret: viper.GetString("client_secret"),
	}
	reply := api.CredentialsReply{}
	err = client.Call("Auth.RefreshAccessToken", &refreshArgs, &reply)
	if err != nil {
		log.Error("no proper credentials set", "error", err)
		return false
	}
	return true
}

func printLeftoverChecks(s map[string]bool) {
	for name, ran := range s {
		if !ran {
			fmt.Printf("- %s\n", name)
		}
	}
}
