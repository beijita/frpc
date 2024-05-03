package fiface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() int64
	RemoteAddr() net.Addr
	SendMsg(msgID int64, data []byte) error
}

type HandFunc func(*net.TCPConn, []byte, int) error
