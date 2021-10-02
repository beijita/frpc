package rpcserver

import (
	"fmt"
)

type HelloService struct {
}

func (p *HelloService) Hello(request string, resp *string) error {
	*resp = fmt.Sprintf("hello, %v", request)
	return nil
}
