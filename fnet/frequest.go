package fnet

import "github.com/frpc/fiface"

type Request struct {
	conn fiface.IConnection
	data []byte
}

func (r *Request) GetConnection() fiface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
