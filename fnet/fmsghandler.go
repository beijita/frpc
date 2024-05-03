package fnet

import (
	"github.com/frpc/fconfig"
	"github.com/frpc/fiface"
	"log"
)

type MsgHandle struct {
	APIMap       map[int64]fiface.IRouter
	WorkPoolSize int
	TaskQueue    []chan fiface.IRequest
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
	return &MsgHandle{APIMap: make(map[int64]fiface.IRouter),
		WorkPoolSize: fconfig.GlobalConf.MaxPacketSize,
		TaskQueue:    make([]chan fiface.IRequest, fconfig.GlobalConf.MaxPacketSize)}
}

func (m *MsgHandle) StartOneWorker(workerID int, taskQueue chan fiface.IRequest) {
	log.Println("MsgHandle.StartOneWorker start workerID=", workerID)
	for {
		select {
		case request := <-taskQueue:
			m.DoMsgHandler(request)
		}
	}
}

func (m *MsgHandle) StartWorkerPool() {
	for i := 0; i < m.WorkPoolSize; i++ {
		m.TaskQueue[i] = make(chan fiface.IRequest, fconfig.GlobalConf.MaxPacketSize)
		go m.StartOneWorker(i, m.TaskQueue[i])
	}
}

func (m *MsgHandle) SendMsgToTaskQueue(request fiface.IRequest) {
	workerID := request.GetConnection().GetConnID() % int64(m.WorkPoolSize)
	m.TaskQueue[workerID] <- request
}
