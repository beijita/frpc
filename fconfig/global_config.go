package fconfig

import (
	"encoding/json"
	"github.com/frpc/fiface"
	"io/ioutil"
	"log"
	"os"
)

type GlobalConfig struct {
	TcpServer     fiface.IServer
	Host          string
	TcpPort       int
	Name          string
	Version       string
	MaxPacketSize int
	MaxConn       int
}

var GlobalConf *GlobalConfig

func (g *GlobalConfig) Reload() {
	if "test" == os.Getenv("GO_ENV") {
		log.Println("GlobalConfig.Reload test env continue")
		return
	}
	log.Println("GlobalConfig.Reload start")
	pwd, _ := os.Getwd()
	log.Println("GlobalConfig.Reload pwd=", pwd)
	data, err := ioutil.ReadFile("conf/fconfig.json")
	if err != nil {
		//log.Fatal("GlobalConfig file is not exist! pwd=", pwd)
		return
	}
	err = json.Unmarshal(data, GlobalConf)
	if err != nil {
		panic(err)
	}
	log.Println("GlobalConfig.Reload GlobalConf=", GlobalConf)
}

func init() {
	GlobalConf = &GlobalConfig{
		Host:          "127.0.0.1",
		TcpPort:       6666,
		Name:          "frpc",
		Version:       "V0.4",
		MaxPacketSize: 4096,
		MaxConn:       8,
	}
	GlobalConf.Reload()
}
