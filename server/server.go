package main

import (
	"github.com/frpc/fhook_conn"
	"github.com/frpc/fnet"
	"github.com/frpc/frouter"
)

func main() {
	server := fnet.NewServer("frpc")

	server.SetOnConnStart(fhook_conn.DoConnectionBegin)
	server.SetOnConnStop(fhook_conn.DoConnectionEnd)
	server.AddRouter(1007, &frouter.PingRouter{})
	server.AddRouter(1009, &frouter.BuyRouter{})
	server.Serve()
}
