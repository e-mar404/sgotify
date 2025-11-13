package cmd

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/e-mar404/sgotify/api"
	msgpackrpc "github.com/hashicorp/net-rpc-msgpackrpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func batch(cmds ...func(*cobra.Command, []string)) func(*cobra.Command, []string) {
	return func(c *cobra.Command, s []string) {
		for _, cmd := range cmds {
			cmd(c, s)
		}
	}
}

func prepLogs(cmd *cobra.Command, args []string) {
	var level log.Level
	baseLevel := log.Level(13)
	verboseLevel := log.Level(0)
	if cmd.Use == "server" { // only the server cmd will have logs by default
		baseLevel = log.Level(0)
		verboseLevel = log.Level(-5)
	}

	if verbose {
		level = verboseLevel
	} else {
		level = baseLevel
	}

	// TODO: should expand the title on the log to have a max width of 5 on the logs that get cut off (Fatal, Debug, Error)
	logger := log.NewWithOptions(os.Stderr, log.Options{
		Level:           level,
		ReportCaller:    true,
		ReportTimestamp: true,
		Formatter:       log.TextFormatter,
	})

	log.SetDefault(logger)
}

func startClient(cmd *cobra.Command, _ []string) {
	if cmd.Use == "server" {
		log.Info("this command does not require rpc client", "cmd", cmd.Use)
		return
	}

	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Printf("couldn't find rpc server\n")
		log.Fatal("unable to connect to server", "error", err)
	}
	client = rpc.NewClientWithCodec(msgpackrpc.NewClientCodec(conn))
}

func requireAuth(cmd *cobra.Command, args []string) {
	if cmd.Use == "login" || cmd.Use == "logout" || cmd.Use == "server" {
		log.Info("this command does not require auth", "cmd", cmd.Use)
		return
	}

	log.Info("checking access token status")
	assert := func(condition bool) {
		if condition {
			fmt.Fprintln(os.Stderr, "Please run `sgotify login` first.")
			log.Fatal("not signed in or wrong state returned")
		}
	}

	accessToken := viper.GetString("access_token")
	assert(accessToken == "")

	last_refresh := viper.GetInt64("last_refresh")
	if time.Now().Add(-time.Minute*55).Unix() <= last_refresh {
		log.Info("Access token is still good, not refreshing")
		return
	}

	log.Info("asking for a new access token")

	refreshArgs := api.RefreshArgs{
		RefreshToken: viper.GetString("refresh_token"),
		ClientID:     viper.GetString("client_id"),
		ClientSecret: viper.GetString("client_secret"),
	}
	reply := api.CredentialsReply{}
	err := client.Call("Auth.RefreshAccessToken", &refreshArgs, &reply)
	assert(err != nil)
	assert(reply.AccessToken == "")

	viper.Set("access_token", reply.AccessToken)
	viper.Set("last_refresh", time.Now().Unix())
	if reply.RefreshToken != "" {
		viper.Set("refresh_token", reply.RefreshToken)
	}

	err = viper.WriteConfig()
	assert(err != nil)
}
