package constants

import (
	"net/http"

	"github.com/e-mar404/sgotify/internal/config"
)

type State struct {
	Client *http.Client
	Cfg *config.File
}
