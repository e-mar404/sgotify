package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

const (
	RedirectURI = "http://127.0.0.1:8080/callback"
	AuthorizeURL = "https://accounts.spotify.com/authorize"
	TokenURL = "https://accounts.spotify.com/api/token"
)

type File struct {
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AuthToken string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
}

func Load() (*File, error) {
	conf := &File{}
	
	f, err := os.Open(path())
	if err != nil {
		return conf, err
	}

	content, err := io.ReadAll(f)
	if err != nil {
		return conf, err
	}

	err = json.Unmarshal(content, &conf)
	if err != nil {
		return conf, err
	}
	
	return conf, nil
}

func (cfg *File) Save() error {
	log.Info("Saving conf file", "cfg", cfg)

	jsonConf, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(path(), jsonConf, 0644)
	if err != nil {
		return err
	}

	return nil
}

func path() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".sgotify.json")
}
