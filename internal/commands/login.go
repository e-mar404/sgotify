package command

import (
	// "net/http"

	"github.com/charmbracelet/log"
	// "github.com/e-mar404/sgotify/internal/auth"
	"github.com/e-mar404/sgotify/internal/config"
)


func Login(cfg config.Config) error {
	log.Info("logging in to spotify")
	log.Info("starting tui, need client id and secret")
	// start charm app to get client id and secret

	// router := auth.NewAuthRouter(cfg)	
	// if err := http.ListenAndServe(":8080", router); err != nil {
	// 	log.Fatal("Auth server crashed", "error", err)
	// }

	return nil
}
