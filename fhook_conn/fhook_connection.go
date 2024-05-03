package fhook_conn

import (
	"github.com/frpc/fiface"
	"log"
)

func DoConnectionBegin(conn fiface.IConnection) {
	log.Println("DoConnectionBegin start")
}

func DoConnectionEnd(conn fiface.IConnection) {
	log.Println("DoConnectionEnd start")
}
