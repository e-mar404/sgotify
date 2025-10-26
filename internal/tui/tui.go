package tui

import (
	"log"
	"net"
	"net/rpc"

	// "os"

	tea "github.com/charmbracelet/bubbletea"
	constants "github.com/e-mar404/sgotify/internal/const"
	msgpackrpc "github.com/hashicorp/net-rpc-msgpackrpc"
)

var client *rpc.Client

func Run() error {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		return err
	}
	client = rpc.NewClientWithCodec(msgpackrpc.NewClientCodec(conn))
	defer client.Close()

	// f, err := tea.LogToFile("debug.log", "debug")
	// if err != nil {
	// 	log.Fatal("error setting log to file", err)
	// 	return err
	// }
	// log.SetOutput(os.Stderr)
	// defer f.Close()

	m := newHomeUI()
	constants.P = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := constants.P.Run(); err != nil {
		log.Fatal("error running program", err)
		return err
	}

	return nil
}
