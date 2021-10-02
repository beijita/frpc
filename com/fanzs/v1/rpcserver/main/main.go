package main

import (
	"frpc/com/fanzs/v1/rpcserver"
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("HelloService", new(rpcserver.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf(" net.Listen error err:%v", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf(" listener.Accept error err:%v", err)
	}
	rpc.ServeConn(conn)
}

