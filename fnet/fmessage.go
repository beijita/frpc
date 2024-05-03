package fnet

type Message struct {
	MsgID   int64
	DataLen int64
	Data    []byte
}

func (m *Message) GetDataLen() int64 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) GetMsgID() int64 {
	return m.MsgID
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) SetMsgID(msgID int64) {
	m.MsgID = msgID
}

func (m *Message) SetDataLen(length int64) {
	m.DataLen = length
}
