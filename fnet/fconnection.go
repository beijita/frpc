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
}

func NewConnection(conn *net.TCPConn, connID int64, callbackAPI fiface.HandFunc) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		handleAPI:    callbackAPI,
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
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			log.Println("Connection.StartReader Conn.Read err", err)
			c.ExitBuffChan <- true
			continue
		}
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			log.Println("Connection.StartReader handleAPI err", err)
			c.ExitBuffChan <- true
			return
		}
	}
}
