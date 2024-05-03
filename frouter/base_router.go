package frouter

import (
	"github.com/frpc/fiface"
	"log"
)

type BaseRouter struct {
}

func (r *BaseRouter) PreHandle(request fiface.IRequest) {
	log.Println("BaseRouter.PreHandle start")
}

func (r *BaseRouter) Handle(request fiface.IRequest) {
	log.Println("BaseRouter.Handle start")
	log.Println("BaseRouter.Handle request=", request)
}

func (r *BaseRouter) PostHandle(request fiface.IRequest) {
	log.Println("BaseRouter.PostHandle start")
}
