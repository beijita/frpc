package fiface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() int64
	RemoteAddr() net.Addr
	SendMsg(msgID int64, data []byte) error
	SendBuffMsg(msgID int64, data []byte) error

	SetProperty(key string, value interface{})
	GetProperty(key string) interface{}
	RemoveProperty(key string)
}

type HandFunc func(*net.TCPConn, []byte, int) error
