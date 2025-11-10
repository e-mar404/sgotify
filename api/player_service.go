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
	BaseURL         string
	AccessToken     string
	DeviceID        string
	PlayRequestBody PlayRequestBody
}

type AvailableDevicesReply struct {
	Devices []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Volume int    `json:"volume"`
	} `json:"devices"`
}

type PlayerReply struct {
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"error,omitzero"`
}

type Offset struct {
	Position int `json:"position,omitempty"`
}
type PlayRequestBody struct {
	ContextURI string   `json:"context_uri,omitempty"`
	URIS       []string `json:"uris,omitempty"`
	Offset     Offset   `json:"offset,omitzero"`
	PositionMs int      `json:"position_ms,omitempty"`
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

func (p *Player) Play(args *PlayerArgs, reply *PlayerReply) error {
	log.Info("called Player.Play")

	p.Client.args = *args

	q := map[string]string{
		"device_id": args.DeviceID,
	}
	u := args.BaseURL + "/me/player/play"

	body := args.PlayRequestBody

	rawBody, _ := json.Marshal(body)
	playReply, err := do[PlayerReply](p.Client, "PUT", u, q, bytes.NewReader(rawBody))
	if err != nil {
		return err
	}

	// if reply is nil then there is no meaningful return, check spotify api for
	// context
	if playReply == nil {
		return nil
	}

	*reply = *playReply

	return nil
}

func (p *Player) Pause(args *PlayerArgs, reply *PlayerReply) error {
	log.Info("called Play.Pause")

	p.Client.args = *args

	q := map[string]string{
		"device_id": args.DeviceID,
	}
	u := args.BaseURL + "/me/player/pause"

	pauseReply, err := do[PlayerReply](p.Client, "PUT", u, q, nil)
	if err != nil {
		return err
	}

	// if there is no meaningful return
	if pauseReply == nil {
		return nil
	}

	*reply = *pauseReply

	return nil
}
