package api

import (
	"net"
	"net/rpc"

	"github.com/charmbracelet/log"
	"github.com/hashicorp/net-rpc-msgpackrpc"
)

var (
	server = rpc.NewServer()
)

func StartRPCServer() error {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Info("starting rpc server", "host", "localhost", "port", ":5000")
	for {
		log.Info("awaiting connection")
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go server.ServeCodec(msgpackrpc.NewServerCodec(conn))
	}
}
