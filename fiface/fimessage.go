package fiface

type IMessage interface {
	GetDataLen() int64
	GetData() []byte
	GetMsgID() int64

	SetData([]byte)
	SetMsgID(msgID int64)
	SetDataLen(length int64)
}
