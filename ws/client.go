package ws

import (
	//"fmt"
	"github.com/gorilla/websocket"
	//"runtime/debug"
)

const (
	heartbeatExpirationTime = 6 * 60
)

type login struct {
	AppId  uint32
	UserId string
	Client *Client
}

type Client struct {
	Addr          string
	Socket        *websocket.Conn
	Send          chan []byte
	AppId         uint32
	UserId        string
	FirstTime     uint64
	HeartbeatTime uint64
	LoginTime     uint64
}

func NewClient(addr string, socket *websocket.Conn, firstTime uint64) (client *Client) {
	client = &Client{
		Addr:          addr,
		Socket:        socket,
		Send:          make(chan []byte, 100),
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
	}

	return
}

func (c *Client) read() {

}

func (c *Client) write() {

}
