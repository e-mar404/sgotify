package api

import (
	"net"
	"net/rpc"

	"github.com/charmbracelet/log"
	"github.com/hashicorp/go-msgpack/codec"
)

var (
	mh     codec.MsgpackHandle
	server = rpc.NewServer()
)

func StartRPCServer() error {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Info("starting rpc server", "host", "localhost", "port", ":5000")
	exit := make(chan error)
	go func() {
		for {
			log.Info("awaiting connection")
			conn, err := listener.Accept()
			if err != nil {
				exit <- err
			}
			codec := codec.GoRpc.ServerCodec(conn, &mh)
			server.ServeCodec(codec)
		}
	}()

	exitErr := <-exit

	log.Error("shutting down rpc server", "error", exitErr)

	return err
}
