package DB

import (
	"LollipopGo/LollipopGo/db"
	_ "LollipopGo/LollipopGo/util"
	"Proto/Proto2"
	"glog-master"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var ConnRPC *rpc.Client // 链接DB server

func InitNetRPC() {
	client, err := jsonrpc.Dial("tcp", db.Service)
	if err != nil {
		glog.Info("dial error:", err)
	}
	ConnRPC = client
}

// 修改数据
func ModefyGamePlayerDataGameInfo(data *Proto2.DB_GameOver) interface{} {

	args := data
	// 返回的数据
	var reply bool
	//--------------------------------------------------------------------------
	// 同步调用
	// err = ConnRPC_GM.Call("Arith.SavePlayerDataST2DB", args, &reply)
	// if err != nil {
	// 	glog.Info("Arith.SavePlayerDataST2DB call error:", err)
	// }
	// 异步调用
	divCall := ConnRPC.Go("Arith.SavePlayerDataST2DB", args, &reply, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	glog.Info(replyCall.Reply)
	//--------------------------------------------------------------------------
	// 返回的数据
	glog.Info("ModefyGamePlayerDataGameInfo :", reply)
	return reply
}
