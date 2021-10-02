package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf(" rpc.Dial error err:%v", err)
	}

	var replay string
	err = client.Call("HelloService.Hello", "fengzhengshu", &replay)
	if err != nil {
		log.Fatalf("rpc Call error! err:%v\n", err)
	}
	fmt.Println(replay)
}
