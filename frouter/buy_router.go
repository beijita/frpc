package frouter

import (
	"github.com/frpc/fiface"
	"log"
)

type BuyRouter struct {
	BaseRouter
}

func (r *BuyRouter) PreHandle(request fiface.IRequest) {
	log.Println("BuyRouter.PreHandle start")
	//_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))
	//if err != nil {
	//	log.Println("PingRouter.PreHandle GetTCPConnection.Write err", err)
	//	return
	//}
}

func (r *BuyRouter) Handle(request fiface.IRequest) {
	log.Println("BuyRouter.Handle start")

	log.Println("BuyRouter.Handle data=", string(request.GetData().GetData()))
	err := request.GetConnection().SendMsg(request.GetData().GetMsgID(), []byte("请支付现金 over"))
	if err != nil {
		log.Println("BuyRouter.Handle GetTCPConnection.Write err", err)
		return
	}
}

func (r *BuyRouter) PostHandle(request fiface.IRequest) {
	log.Println("PostHandle start")
	//_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n"))
	//if err != nil {
	//	log.Println("PingRouter.PostHandle GetTCPConnection.Write err", err)
	//	return
	//}
}
