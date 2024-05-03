package main

import (
	"github.com/frpc/fnet"
	"github.com/frpc/frouter"
)

func main() {
	server := fnet.NewServer("frpc")
	server.AddRouter(&frouter.PingRouter{})
	server.Serve()
}
