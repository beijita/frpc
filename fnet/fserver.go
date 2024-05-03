package fnet

import (
	"errors"
	"fmt"
	"github.com/frpc/fconfig"
	"github.com/frpc/fiface"
	"log"
	"net"
)

type Server struct {
	Name        string
	IPVersion   string
	IP          string
	Port        int
	ApiHandle   fiface.IMsgHandle
	ConnMgr     fiface.IConnManager
	OnConnStart func(connection fiface.IConnection)
	OnConnStop  func(connection fiface.IConnection)
}

func (s *Server) GetConnMgr() fiface.IConnManager {
	return s.ConnMgr
}

func (s *Server) AddRouter(msgId int64, router fiface.IRouter) {
	s.ApiHandle.AddRouter(msgId, router)
}

func NewServer(name string) fiface.IServer {
	fconfig.GlobalConf.Reload()
	return &Server{
		Name:      fconfig.GlobalConf.Name,
		IPVersion: "tcp4",
		IP:        fconfig.GlobalConf.Host,
		Port:      fconfig.GlobalConf.TcpPort,
		ApiHandle: NewMsgHandle(),
		ConnMgr:   NewConnManager(),
	}
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	log.Println("CallBackToClient Start")
	if _, err := conn.Write(data[:cnt]); err != nil {
		return errors.New("CallBackToClient conn.write err")
	}
	return nil
}

func (s *Server) Start() {
	log.Println("FServer Start")

	go func() {
		s.ApiHandle.StartWorkerPool()
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Println("net.ResolveTCPAddr err=", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Println("net.ListenTCP err=", err)
			return
		}

		var cid int64
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("listener.AcceptTCP err=", err)
				continue
			}
			dealConn := NewConnection(s, conn, cid, s.ApiHandle)
			cid++
			go dealConn.Start()

			if s.GetConnMgr().Len() > fconfig.GlobalConf.MaxConn {
				dealConn.Stop()
			}
		}
	}()
}

func (s *Server) Stop() {
	log.Println("FServer Stop")
	s.GetConnMgr().ClearAll()
}

func (s *Server) Serve() {
	log.Println("FServer Serve")
	s.Start()
	select {}
}

func (s *Server) SetOnConnStart(f func(connection fiface.IConnection)) {
	s.OnConnStart = f
}

func (s *Server) SetOnConnStop(f func(connection fiface.IConnection)) {
	s.OnConnStop = f
}

func (s *Server) CallOnConnStart(connection fiface.IConnection) {
	log.Println("Server.CallOnConnStart start, OnConnStart=", s.OnConnStart)
	if s.OnConnStart != nil {
		s.OnConnStart(connection)
	}
}

func (s *Server) CallOnConnStop(connection fiface.IConnection) {
	log.Println("Server.CallOnConnStop start, OnConnStop=", s.OnConnStop)
	if s.OnConnStop != nil {
		s.OnConnStop(connection)
	}
}
