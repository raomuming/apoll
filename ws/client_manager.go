package ws

import (
	"fmt"
	"sync"
	//"time"
)

type ClientManager struct {
	Clients     map[*Client]bool
	ClientsLock sync.RWMutex
	Users       map[string]*Client
	UserLock    sync.RWMutex
	Register    chan *Client
	Login       chan *login
	Unregister  chan *Client
	Broadcast   chan []byte
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
		Login:      make(chan *login, 1000),
		Unregister: make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}

	return
}

func GetUserKey(appId uint32, userId string) (key string) {
	key = fmt.Sprintf("%d_%s", appId, userId)
	return
}

func (manager *ClientManager) start() {

}
