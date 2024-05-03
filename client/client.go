package main

import (
	"github.com/frpc/fnet"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		log.Println("net.Dial err=", err)
		return
	}
	//for {
	SendMsg(conn, 1007, "hello")
	SendMsg(conn, 1009, "变形金刚")
	time.Sleep(3 * time.Second)

	conn.Close()
	//}
}

func SendMsg(conn net.Conn, msgId int64, message string) {
	var dp fnet.DataPack
	packData, err := dp.PackData(&fnet.Message{
		MsgID:   msgId,
		DataLen: int64(len(message)),
		Data:    []byte(message),
	})
	if err != nil {
		return
	}
	_, err = conn.Write(packData)
	if err != nil {
		log.Println("conn.Write err=", err)
		return
	}
	headData := make([]byte, dp.GetHeadLen())
	_, err = io.ReadFull(conn, headData)
	if err != nil {
		log.Println("Connection.StartReader Conn.Read err", err)
		return
	}
	log.Println("Connection.StartReader Conn.Read headData=", headData)

	msg, err := dp.UnPackData(headData)
	if err != nil {
		return
	}
	log.Println("Connection.StartReader Conn.Read msg=", msg)
	buf := make([]byte, msg.GetDataLen())
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		log.Println("conn.Read err=", err)
		return
	}
	log.Println("conn.Read success buf=", string(buf))
}
