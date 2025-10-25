package api

import (
	"net/http"
	"time"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

type DeviceItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Volume int    `json:"volume_percent"`
}

type DeviceList struct {
	Devcies []DeviceItem `json:"devices"`
}

type DevicesClient struct {
	HTTP  *http.Client
	Query map[string]string
}

func (dc *DevicesClient) prep(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+viper.GetString("access_token"))
}

func (dc *DevicesClient) do(req *http.Request) (*http.Response, error) {
	return dc.HTTP.Do(req)
}

func NewDevicesClient() *DevicesClient {
	return &DevicesClient{
		HTTP: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (dc *DevicesClient) AvailableDevices() (*DeviceList, error) {
	url := viper.GetString("spotify_api_url") + "/me/player/devices"
	var devicesRes *DeviceList
	devicesRes, err := do[DeviceList](dc, "GET", url, nil)
	if err != nil {
		log.Error("could not complete available device list", "error", err)
		return nil, err
	}
	return devicesRes, nil
}
