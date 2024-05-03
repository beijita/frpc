package fhook_conn

import (
	"github.com/frpc/fiface"
	"log"
)

func DoConnectionBegin(conn fiface.IConnection) {
	log.Println("DoConnectionBegin start")
	conn.SetProperty("year", "2024")
}

func DoConnectionEnd(conn fiface.IConnection) {
	log.Println("DoConnectionEnd start")
	log.Println("property date year=", conn.GetProperty("year"))
}
