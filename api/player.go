package api

import (
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type PlayerClient struct {
	HTTP *http.Client
	Query map[string]string
}

type Player struct {
	DeviceID string
	DeviceName string
	Volume int
	Client *PlayerClient
	// ik there will be other stuff here
}

func NewPlayer() *Player {
	// TODO: needs assets when getting things from viper
	return &Player{
		DeviceID: viper.GetString("default_device_id"),
		DeviceName: viper.GetString("default_device_name"),
		Volume: 100, // this is the same default that spotify gives so I will just put it here, I dont want to save it to the config
		Client: &PlayerClient{
			HTTP: &http.Client {
				Timeout: 10 * time.Second,
			},
		},
	}
}
