package fnet

import (
	"errors"
	"github.com/frpc/fiface"
	"io"
	"log"
	"net"
)

type Connection struct {
	Conn         *net.TCPConn
	ConnID       int64
	isClosed     bool
	handleAPI    fiface.HandFunc
	ExitBuffChan chan bool
	ApiHandle    fiface.IMsgHandle
}

func (c *Connection) SendMsg(msgID int64, data []byte) error {
	if c.isClosed {
		return errors.New("connection is closed")
	}
	var dp DataPack
	packData, err := dp.PackData(&Message{
		MsgID:   msgID,
		DataLen: int64(len(data)),
		Data:    data,
	})
	if err != nil {
		return err
	}
	_, err = c.Conn.Write(packData)
	if err != nil {
		c.ExitBuffChan <- true
		return err
	}
	return nil
}

func NewConnection(conn *net.TCPConn, connID int64, apiHandle fiface.IMsgHandle) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		ApiHandle:    apiHandle,
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
	var dp DataPack
	for {
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			log.Println("Connection.StartReader Conn.Read err", err)
			c.ExitBuffChan <- true
			continue
		}
		message, err := dp.UnPackData(headData)
		if err != nil {
			return
		}
		if message.GetDataLen() > 0 {
			data := make([]byte, message.GetDataLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				return
			}
			message.SetData(data)
		}
		req := Request{
			conn: c,
			msg:  message,
		}
		go func(request fiface.IRequest) {
			c.ApiHandle.DoMsgHandler(request)
		}(&req)
	}
}
