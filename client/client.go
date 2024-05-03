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
	var dp fnet.DataPack
	for {
		packData, err := dp.PackData(&fnet.Message{
			MsgID:   1007,
			DataLen: int64(len("hello,world")),
			Data:    []byte("hello,world"),
		})
		if err != nil {
			return
		}
		_, err = conn.Write(packData)
		if err != nil {
			log.Println("conn.Write err=", err)
			continue
		}
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			log.Println("Connection.StartReader Conn.Read err", err)
			continue
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
			continue
		}
		log.Println("conn.Read success buf=", string(buf))

		time.Sleep(3 * time.Second)
	}
}
