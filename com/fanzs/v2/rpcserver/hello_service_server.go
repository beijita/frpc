package rpcserver

import (
	"fmt"
	"net/rpc"
)

const HelloServiceName = "HelloService"

type HelloServiceInterface = interface {
	Hello(request string, replay *string)
}

func RegisterHelloService(svc *HelloService) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloService struct {
}

func (p *HelloService) Hello(request string, resp *string) error {
	*resp = fmt.Sprintf("hello, %v", request)
	return nil
}
