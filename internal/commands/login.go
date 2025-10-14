package command

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

func Login() error {
	path := "~/.sgotify.json"
	log.Info("Checking for client id & client secret", "path", path)
	
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Info("Configuration file does not exists creating new config file") 
		createConfig(path)	
	}

	return nil
}

func createConfig(path string) {
	var ClientID, ClientSecret string
	fmt.Printf("What is your ClientID?\n")
	fmt.Print("> ")
	fmt.Scan(&ClientID)


	fmt.Printf("\nWhat is your ClientSecret?\n")
	fmt.Print("> ")
	fmt.Scan(&ClientSecret)

	log.Info("Saving to config file", "ClientID", ClientID, "ClientSecret", ClientSecret, "path", path)
}
