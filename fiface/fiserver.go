package fiface

type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(msgID int64, router IRouter)
	GetConnMgr() IConnManager

	SetOnConnStart(func(connection IConnection))
	SetOnConnStop(func(connection IConnection))
	CallOnConnStart(connection IConnection)
	CallOnConnStop(connection IConnection)
}
