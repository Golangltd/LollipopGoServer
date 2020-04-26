package main

import (
	"LollipopGo/LollipopGo/network"
	"LollipopGo/LollipopGo/util"
	"Proto"
	"Proto/Proto3"
	"flag"
	"gameMJ/db"
	"glog-master"

	"code.google.com/p/go.net/websocket"
)

var addr = flag.String("addr", "127.0.0.1:8866", "http service address") // 链接gateway
var Conn *websocket.Conn                                                 // 保存用户的链接信息，数据会在主动匹配成功后进行链接

func init() {
	initGateWayNet()
}

func initGateWayNet() bool {

	url := "ws://" + *addr + "/BaBaLiuLiu?data=LollipopGo"
	conn, err := websocket.Dial(url, "", "test://golang/")
	if err != nil {
		glog.Info("err:", err.Error())
		return false
	}
	Conn = conn
	initConn(Conn)
	DB.InitNetRPC()
	impl.InitConnection(conn)
	return true
}

func initConn(conn *websocket.Conn) {
	data := &Proto3_Data.MJ2GW_ConnInit{
		Protocol:  Proto.G_GameBGServer_Proto, // 游戏主要协议
		Protocol2: Proto3_Data.MJ2GW_ConnInitProto2,
		ServerID:  util.MD5_LollipopGO("MJserver"),
	}
	impl.PlayerSendToServer(conn, data)
	return
}
