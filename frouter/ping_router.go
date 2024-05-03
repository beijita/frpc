package frouter

import (
	"github.com/frpc/fiface"
	"log"
)

type PingRouter struct {
	BaseRouter
}

func (p *PingRouter) PreHandle(request fiface.IRequest) {
	log.Println("PingRouter.PreHandle start")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))
	if err != nil {
		log.Println("PingRouter.PreHandle GetTCPConnection.Write err", err)
		return
	}
}

func (p *PingRouter) Handle(request fiface.IRequest) {
	log.Println("PingRouter.Handle start")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("handle ping\n"))
	if err != nil {
		log.Println("PingRouter.Handle GetTCPConnection.Write err", err)
		return
	}
}

func (p *PingRouter) PostHandle(request fiface.IRequest) {
	log.Println("PingRouter.PostHandle start")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n"))
	if err != nil {
		log.Println("PingRouter.PostHandle GetTCPConnection.Write err", err)
		return
	}
}
