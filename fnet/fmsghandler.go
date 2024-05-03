package fnet

import (
	"github.com/frpc/fiface"
	"log"
)

type MsgHandle struct {
	APIMap map[int64]fiface.IRouter
}

func (m *MsgHandle) DoMsgHandler(request fiface.IRequest) {
	if handler, ok := m.APIMap[request.GetData().GetMsgID()]; ok {
		handler.PreHandle(request)
		handler.Handle(request)
		handler.PostHandle(request)
	} else {
		log.Println("MsgHandle.DoMsgHandler msgID=", request.GetData().GetMsgID())
	}
}

func (m *MsgHandle) AddRouter(msgID int64, router fiface.IRouter) {
	m.APIMap[msgID] = router
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{APIMap: make(map[int64]fiface.IRouter)}
}
