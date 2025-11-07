package api

import "github.com/charmbracelet/log"

type Player struct {
	Client *PlayerClient
}

type PlayerArgs struct {
	BaseURL     string
	AccessToken string
}

type AvailableDevicesReply struct {
	Devices []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Volume int    `json:"volume"`
	} `json:"devices"`
}

func init() {
	server.Register(NewPlayerService())
}

func NewPlayerService() *Player {
	return &Player{
		Client: NewPlayerClient(),
	}
}

func (p *Player) AvailableDevices(args *PlayerArgs, reply *AvailableDevicesReply) error {
	log.Info("called Player.AvailableDevices")

	p.Client.args.BaseURL = args.BaseURL
	p.Client.args.AccessToken = args.AccessToken

	url := args.BaseURL + "/me/player/devices"
	availableDevices, err := do[AvailableDevicesReply](p.Client, "GET", url, nil)
	if err != nil {
		return err
	}

	log.Debug("got reply", "available devices", availableDevices)

	*reply = *availableDevices

	return nil
}
