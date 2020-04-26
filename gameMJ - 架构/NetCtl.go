package main

import (
	"glog-master"

	"encoding/json"

	"code.google.com/p/go.net/websocket"
)

func (this *MJGame) PlayerSendMessage(senddata interface{}) int {

	b, err1 := json.Marshal(senddata)
	if err1 != nil {
		glog.Error("PlayerSendMessage json.Marshal data fail ! err:", err1.Error())
		glog.Flush()
		return 1
	}
	glog.Info("json.Marshal(b) :", string(b))
	data := ""
	data = "data" + "=" + string(b[0:len(b)])
	glog.Info("json.Marshal(data) :", data)
	glog.Flush()
	err := websocket.JSON.Send(this.Connection, b)
	if err != nil {
		glog.Error("PlayerSendMessage send data fail ! err:", err.Error())
		glog.Flush()
		return 2
	}
	return 0
}
