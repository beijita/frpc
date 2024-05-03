package main

import "github.com/frpc/fnet"

func main() {
	server := fnet.NewServer("frpc")
	server.Serve()
}
