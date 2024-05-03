package fnet

import (
	"log"
	"testing"
)

func TestDataPack_PackData(t *testing.T) {
	msg := Message{
		MsgID:   1007,
		DataLen: int64(len("hello,world")),
		Data:    []byte("hello,world"),
	}
	pack := DataPack{}
	data, err := pack.PackData(&msg)
	if err != nil {
		return
	}
	log.Println("dataPack=", data)

	unPackData, err := pack.UnPackData(data)
	if err != nil {
		return
	}

	log.Println("unPackData=", unPackData)
	log.Println("unPackData.GetData=", string(unPackData.GetData()))
	log.Println("unPackData.GetData=", string(data[pack.GetHeadLen():pack.GetHeadLen()+unPackData.GetDataLen()]))
}
