package rpcclient

import (
	"frpc/com/fanzs/v2/rpcserver"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

var _HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: client}, nil
}

func (p *HelloServiceClient) Hello(request string, response *string) error {
	return p.Client.Call(rpcserver.HelloServiceName+".Hello", request, response)
}
