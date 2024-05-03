package fnet

import (
	"github.com/frpc/fiface"
	"log"
	"net"
)

type Connection struct {
	Conn         *net.TCPConn
	ConnID       int64
	isClosed     bool
	handleAPI    fiface.HandFunc
	ExitBuffChan chan bool
	Router       fiface.IRouter
}

func NewConnection(conn *net.TCPConn, connID int64, router fiface.IRouter) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}
}

func (c *Connection) Start() {
	log.Println("Connection.Start start")
	go c.StartReader()
	for {
		select {
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Stop() {
	log.Println("Connection.Stop start")
	if c.isClosed {
		return
	}
	c.isClosed = true
	c.ExitBuffChan <- true
	c.Conn.Close()
	close(c.ExitBuffChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() int64 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) StartReader() {
	log.Println("Connection.StartReader start")
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			log.Println("Connection.StartReader Conn.Read err", err)
			c.ExitBuffChan <- true
			continue
		}
		req := Request{
			conn: c,
			data: buf,
		}
		go func(request fiface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}
