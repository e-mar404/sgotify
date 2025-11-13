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
	log.Info("Player.AvailableDevices called")

	p.Client.args.AccessToken = args.AccessToken

	url := apiBaseURL + "/me/player/devices"
	availableDevices, err := do[AvailableDevicesReply](p.Client, "GET", url, nil, nil)
	if err != nil {
		return err
	}

	*reply = *availableDevices

	jsonReply, _ := json.MarshalIndent(*reply, "", " ")
	log.Debug("sending reply", "AvailableDevicesReply", string(jsonReply))
	log.Info("Player.AvailableDevices sent reply")

	return nil
}

func (p *Player) Play(args *PlayerArgs, reply *PlayerReply) error {
	log.Info("Player.Play called")

	p.Client.args = *args

	q := map[string]string{
		"device_id": args.DeviceID,
	}
	u := apiBaseURL + "/me/player/play"

	body := args.PlayRequestBody

	rawBody, _ := json.Marshal(body)
	playReply, err := do[PlayerReply](p.Client, "PUT", u, q, bytes.NewReader(rawBody))
	if err != nil {
		return err
	}

	// if reply is nil then there is no meaningful return, check spotify api for
	// context
	if playReply == nil {
		log.Debug("no content to reply")
		return nil
	}

	*reply = *playReply

	jsonReply, _ := json.MarshalIndent(*reply, "", " ")
	log.Debug("sending reply", "PlayerReply", string(jsonReply))
	log.Info("Player.Play sent reply")

	return nil
}

func (p *Player) Pause(args *PlayerArgs, reply *PlayerReply) error {
	log.Info("Play.Pause called")

	p.Client.args = *args

	q := map[string]string{
		"device_id": args.DeviceID,
	}
	u := apiBaseURL + "/me/player/pause"

	pauseReply, err := do[PlayerReply](p.Client, "PUT", u, q, nil)
	if err != nil {
		return err
	}

	// if there is no meaningful return
	if pauseReply == nil {
		log.Debug("no content to reply")
		return nil
	}

	*reply = *pauseReply

	jsonReply, _ := json.MarshalIndent(*reply, "", " ")
	log.Debug("sending reply", "PlayerReply", string(jsonReply))
	log.Info("Player.Pause sent reply")

	return nil
}

func (p *Player) Next(args *PlayerArgs, reply *PlayerReply) error {
	log.Info("Player.Next called")

	p.Client.args = *args

	u := apiBaseURL + "/me/player/next"

	nextReply, err := do[PlayerReply](p.Client, "POST", u, nil, nil)
	if err != nil {
		return err
	}

	if nextReply == nil {
		log.Debug("no content to reply")
		return nil
	}

	*reply = *nextReply

	jsonReply, _ := json.MarshalIndent(*reply, "", " ")
	log.Debug("sending reply", "PlayerReply", string(jsonReply))
	log.Info("Player.Next sent reply")

	return nil
}

func (p *Player) Prev(args *PlayerArgs, reply *PlayerReply) error {
	log.Info("Player.Prev called")

	p.Client.args = *args

	u := apiBaseURL + "/me/player/previous"

	nextReply, err := do[PlayerReply](p.Client, "POST", u, nil, nil)
	if err != nil {
		return err
	}

	if nextReply == nil {
		log.Debug("no content to reply")
		return nil
	}

	*reply = *nextReply

	jsonReply, _ := json.MarshalIndent(*reply, "", " ")
	log.Debug("sending reply", "PlayerReply", string(jsonReply))
	log.Info("Player.Prev sent reply")

	return nil
}
