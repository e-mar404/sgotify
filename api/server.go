package api

import (
	"encoding/json"
	"net"
	"net/http"
	// "net/http/cookiejar"
	"net/rpc"
	"os"

	"github.com/charmbracelet/log"
	"github.com/hashicorp/net-rpc-msgpackrpc"
)

var (
	server = rpc.NewServer()
	// jar *cookiejar.Jar
)

func StartRPCServer() error {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		return err
	}
	defer listener.Close()
	
	// need to init cookiejar here so it can be used across all diff services
	// jar, _ = cookiejar.New(nil)
		
	log.Info("starting rpc server", "host", "localhost", "port", ":5000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go server.ServeCodec(msgpackrpc.NewServerCodec(conn))
	}
}

func saveCookies(cookies []*http.Cookie) error {
	data, err := json.Marshal(cookies)
	if err != nil {
		log.Error("unable to json marshal cookes", err)
		return err
	}
	return os.WriteFile("/home/emar/.config/sgotify/cookies.json",data, 0600)
}
