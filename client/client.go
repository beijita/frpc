package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		log.Println("net.Dial err=", err)
		return
	}
	for {
		_, err := conn.Write([]byte("hello frpc"))
		if err != nil {
			log.Println("conn.Write err=", err)
			continue
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			log.Println("conn.Read err=", err)
			continue
		}
		log.Println("conn.Read success buf=", string(buf), "cnt=", cnt)

		time.Sleep(3 * time.Second)
	}
}
