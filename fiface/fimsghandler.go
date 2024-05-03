package fiface

type IMsgHandle interface {
	DoMsgHandler(request IRequest)
	AddRouter(msgID int64, router IRouter)
}
