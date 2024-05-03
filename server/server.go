package main

import (
	"github.com/frpc/fnet"
	"github.com/frpc/frouter"
)

func main() {
	server := fnet.NewServer("frpc")
	server.AddRouter(1007, &frouter.PingRouter{})
	server.AddRouter(1009, &frouter.BuyRouter{})
	server.Serve()
}
