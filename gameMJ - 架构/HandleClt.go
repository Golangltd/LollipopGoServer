package main

import (
	"Proto"
	"Proto/Proto3"
	"glog-master"

	"code.google.com/p/go.net/websocket"
)

type Tsts struct {
	Id int
}

func (this *MJGame) HandleCltProtocol(protocol interface{}, protocol2 interface{}, ProtocolData map[string]interface{}, Connection *websocket.Conn) interface{} {
	switch protocol {
	case float64(Proto.G_GameBG_Proto):
		{
			this.HandleCltProtocol2(protocol2, ProtocolData, Connection)
		}
	default:
		glog.Info("protocol default")
	}
	return 0
}

func (this *MJGame) HandleCltProtocol2(protocol2 interface{}, ProtocolData map[string]interface{}, Connection *websocket.Conn) interface{} {
	ConnectionData := &MJGame{
		Connection: Connection,
		MapSafe:    M,
	}
	switch protocol2 {
	case float64(Proto3_Data.C2GW_ConnLoginProto2):
		{
			ConnectionData.User_Login(ProtocolData)
		}
	case float64(Proto3_Data.MJS2GW_ConnLoginProto2):
		{
			ConnectionData.Server_Login(ProtocolData)
		}
	default:
		glog.Info("protocol2 default")
	}
	return 0
}
