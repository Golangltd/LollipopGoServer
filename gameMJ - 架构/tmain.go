package main

import (
	"LollipopGo/LollipopGo/network"
	"cache2go"
	"glog-master"
	"go-concurrentMap-master"
	"net/http"

	"code.google.com/p/go.net/websocket"
)

type MJGame struct {
	Connection *websocket.Conn
	StrMD5     string
	MapSafe    *concurrent.ConcurrentMap
}

var cache *cache2go.CacheTable
var M *concurrent.ConcurrentMap

func init() {
	impl.IMsg = new(MJGame)
	cache = cache2go.Cache("MJgame")
	M = concurrent.NewConcurrentMap()
	return
}

func main() {
	http.Handle("/BaBaLiuLiu_MJ", websocket.Handler(BuildConnection))
	if err := http.ListenAndServe(":8867", nil); err != nil {
		glog.Info("Entry nil", err.Error())
		glog.Flush()
		return
	}
	return
}

func BuildConnection(ws *websocket.Conn) {
	data := ws.Request().URL.Query().Get("data")
	glog.Info(data)
	if data == "" {
		glog.Info("data is Nil")
		glog.Flush()
		return
	}
	impl.InitConnection(ws)
}
