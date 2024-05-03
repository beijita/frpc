package fnet

import (
	"errors"
	"fmt"
	"github.com/frpc/fiface"
	"log"
	"net"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    fiface.IRouter
}

func (s *Server) AddRouter(router fiface.IRouter) {
	s.Router = router
}

func NewServer(name string) fiface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "127.0.0.1",
		Port:      7777,
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
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	log.Println("FServer Stop")
}

func (s *Server) Serve() {
	log.Println("FServer Serve")
	s.Start()
	select {}
}
