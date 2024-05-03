package fiface

type IRequest interface {
	GetConnection() IConnection
	GetData() IMessage
}
