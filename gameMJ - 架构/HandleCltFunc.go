package main

import (
	"glog-master"
)

func (this *MJGame) User_Login(ProtocolData map[string]interface{}) {

	OpenID := ""
	onlineUser := &MJGame{
		Connection: this.Connection,
		MapSafe:    this.MapSafe,
	}
	this.MapSafe.Put(OpenID+"|User", onlineUser)

	glog.Info("User_Login")
	data := &Tsts{
		Id: 2,
	}
	this.PlayerSendMessage(data)
	return
}

func (this *MJGame) Server_Login(ProtocolData map[string]interface{}) {
	onlineUser := &MJGame{
		Connection: this.Connection,
		MapSafe:    this.MapSafe,
	}
	serverid := ""
	this.MapSafe.Put(serverid+"|Server", onlineUser)

	glog.Info("User_Login")
	data := &Tsts{
		Id: 2,
	}
	this.PlayerSendMessage(data)
	return
}

func (this *MJGame) Get_UserConn(strOpenID string) {
	val, _ := this.MapSafe.Get(strOpenID + "|connect")
	if val == nil {
	} else {
		val.(interface{}).(*MJGame).PlayerSendMessage("test")
		return
	}
}
