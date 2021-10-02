package main

import (
	"frpc/com/fanzs/v2/rpcserver"
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpcserver.RegisterHelloService(new(rpcserver.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf(" net.Listen error err:%v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf(" listener.Accept error err:%v", err)
		}
		go rpc.ServeConn(conn)
	}
}
