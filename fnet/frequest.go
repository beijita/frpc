package fnet

import "github.com/frpc/fiface"

type Request struct {
	conn fiface.IConnection
	msg  fiface.IMessage
}

func (r *Request) GetData() fiface.IMessage {
	return r.msg
}

func (r *Request) GetConnection() fiface.IConnection {
	return r.conn
}

func (r *Request) GetMessageID() int64 {
	return r.msg.GetMsgID()
}
