package fiface

type IDataPack interface {
	GetHeadLen() int64
	PackData(message IMessage) ([]byte, error)
	UnPackData([]byte) (IMessage, error)
}
