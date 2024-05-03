package fiface

type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(msgID int64, router IRouter)
	GetConnMgr() IConnManager
}
