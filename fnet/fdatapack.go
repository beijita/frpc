package fnet

import (
	"bytes"
	"encoding/binary"
	"github.com/frpc/fiface"
)

type DataPack struct {
}

func (d *DataPack) GetHeadLen() int64 {
	return 16
}

func (d *DataPack) PackData(message fiface.IMessage) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := binary.Write(buffer, binary.LittleEndian, message.GetDataLen())
	if err != nil {
		return nil, err
	}
	err = binary.Write(buffer, binary.LittleEndian, message.GetMsgID())
	if err != nil {
		return nil, err
	}
	err = binary.Write(buffer, binary.LittleEndian, message.GetData())
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (d *DataPack) UnPackData(data []byte) (fiface.IMessage, error) {
	reader := bytes.NewReader(data)
	msg := Message{}
	err := binary.Read(reader, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}
	err = binary.Read(reader, binary.LittleEndian, &msg.MsgID)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}
