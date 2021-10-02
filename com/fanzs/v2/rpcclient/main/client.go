package main

import (
	"fmt"
	"frpc/com/fanzs/v2/rpcclient"
	"log"
)

func main() {
	client, err := rpcclient.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf(" rpc.Dial error err:%v", err)
	}

	var replay string
	err = client.Hello("HelloService.Hello",  &replay)
	if err != nil {
		log.Fatalf("rpc Call error! err:%v\n", err)
	}
	fmt.Println(replay)
}
