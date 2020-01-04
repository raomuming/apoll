package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"runtime/debug"
)

const (
	heartbeatExpirationTime = 6 * 60
)

type login struct {
	AppId  uint32
	UserId string
	Client *Client
}

// get client info
func (l *login) GetKey() (key string) {
	key = GetUserKey(l.AppId, l.UserId)
	return
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

// get client info
func (c *Client) GetKey() (key string) {
	key = GetUserKey(c.AppId, c.UserId)
	return
}

func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("read stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		fmt.Println("close send chan", c)
		close(c.Send)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			fmt.Println("read data from client error", c.Addr, err)
			return
		}

		// process
		fmt.Println("process data:", string(message))
		ProcessData(c, message)
	}
}

func (c *Client) write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		clientManager.Unregister <- c
		c.Socket.Close()
		fmt.Println("Client send message defer", c)
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				fmt.Println("Client send data, close connection", c.Addr, "ok", ok)
				return
			}

			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) SendMsg(msg []byte) {
	if c == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()

	c.Send <- msg
}

func (c *Client) close() {
	close(c.Send)
}

func (c *Client) Login(appId uint32, userId string, loginTime uint64) {
	c.AppId = appId
	c.UserId = userId
	c.LoginTime = loginTime

	c.Heartbeat(loginTime)
}

func (c *Client) Heartbeat(currentTime uint64) {
	c.HeartbeatTime = currentTime
}

func (c *Client) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.HeartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}
	return
}

func (c *Client) IsLogin() (isLogin bool) {
	if c.UserId != "" {
		isLogin = true
		return
	}
	return
}
