package api

import (
	"bytes"
	"encoding/json"

	"github.com/charmbracelet/log"
)

type Player struct {
	Client *PlayerClient
}

type PlayerArgs struct {
	BaseURL     string
	AccessToken string
	DeviceID    string
}

type AvailableDevicesReply struct {
	Devices []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Volume int    `json:"volume"`
	} `json:"devices"`
}

type PlayReply struct {
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"error,omitzero"`
}

type Offset struct {
	Position int `json:"position,omitempty"`
}
type PlayRequest struct {
	ContextURI string `json:"context_uri,omitempty"`
	Offset     Offset `json:"offset,omitzero"`
	PositionMs int    `json:"position_ms,omitempty"`
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
	availableDevices, err := do[AvailableDevicesReply](p.Client, "GET", url, nil, nil)
	if err != nil {
		return err
	}

	log.Debug("got reply", "available devices", availableDevices)

	*reply = *availableDevices

	return nil
}

func (p *Player) Play(args *PlayerArgs, reply *PlayReply) error {
	log.Info("called Player.Play")

	p.Client.args = *args

	q := map[string]string{
		"device_id": args.DeviceID,
	}
	u := args.BaseURL + "/me/player/play"

	// TODO: add this to PlayerArgs
	body := PlayRequest{}

	rawBody, _ := json.Marshal(body)
	playReply, err := do[PlayReply](p.Client, "PUT", u, q, bytes.NewReader(rawBody))
	if err != nil {
		return err
	}

	// if reply is nil then there is no meaningful return
	if playReply == nil {
		return nil
	}

	*reply = *playReply

	return nil
}
