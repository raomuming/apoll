package ws

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	clientManager = NewClientManager()
	appIds        = []uint32{101, 102}
)

func GetAppIds() []uint32 {
	return appIds
}

func StartWebSocket() {
	http.HandleFunc("/acc", wsHandler)
	go clientManager.start()

	fmt.Println("Websocket start")
	http.ListenAndServe(":"+"8089", nil)
}

func wsHandler(w http.ResponseWriter, req *http.Request) {
	// upgrade protocol to websocket
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("upgrade protocol", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)
		return
	}

	fmt.Println("webSocket established:", conn.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

	go client.read()
	go client.write()

	// user connect event
	clientManager.Register <- client
}
