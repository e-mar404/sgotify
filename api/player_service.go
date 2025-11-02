package api

import "github.com/charmbracelet/log"

type Player struct {
	Client *PlayerClient
}

type PlayerArgs struct {
	BaseUrl string
	AccessToken string
}

type AvailableDevicesReply struct {
	Devices []struct {
		ID string `json:"id"`
		Name string `json:"name"`	
		Volume int `json:"volume"`
	} `json:"devices"`
}

func init() {
	server.Register(&Player{})
}

func NewPlayerService() *Player {
	return &Player{
		Client: NewPlayerClient(),
	}
}

func (p *Player) AvailableDevices(args *PlayerArgs, reply *AvailableDevicesReply) error {
	p.Client.args = *args

	url := args.BaseUrl + "/me/player/devices"
	availableDevices, err := do[AvailableDevicesReply](p.Client, "GET", url, nil)
	if err != nil {
		return err
	}
	
	log.Info("got reply", "available devices", availableDevices)

	*reply = *availableDevices

	return nil
}

