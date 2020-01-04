package ws

import (
	"encoding/json"
	"fmt"
	"sync"

	"apoll/ws/models"
)

type DisposeFunc func(client *Client, seq string, message []byte) (code uint32, msg string, data interface{})

var (
	handlers        = make(map[string]DisposeFunc)
	handlersRWMutex sync.RWMutex
)

func Register(key string, value DisposeFunc) {
	handlersRWMutex.Lock()
	defer handlersRWMutex.Unlock()
	handlers[key] = value
}

func getHandlers(key string) (value DisposeFunc, ok bool) {
	handlersRWMutex.RLock()
	defer handlersRWMutex.RUnlock()

	value, ok = handlers[key]
	return
}

func ProcessData(client *Client, message []byte) {
	fmt.Println("process data", client.Addr, string(message))

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("process data stop", r)
		}
	}()

	request := &models.Request{}

	err := json.Unmarshal(message, request)
	if err != nil {
		fmt.Println("process data, json Unmarshal", err)
		client.SendMsg([]byte("illegal data"))
		return
	}

	requestData, err := json.Marshal(request.Data)
	if err != nil {
		fmt.Println("process data, json Marshal", err)
		client.SendMsg([]byte("process data failed"))
		return
	}

	seq := request.Seq
	cmd := request.Cmd

	var (
		code uint32
		msg  string
		data interface{}
	)

	// request
	fmt.Println("acc_request", cmd, client.Addr)

	if value, ok := getHandlers(cmd); ok {
		code, msg, data = value(client, seq, requestData)
	} else {
		code = RoutingNotExist
		fmt.Println("process data, route not exists", client.Addr, "cmd", cmd)
	}

	msg = GetErrorMessage(code, msg)

	responseHead := models.NewResponseHead(seq, cmd, code, msg, data)

	headByte, err := json.Marshal(responseHead)
	if err != nil {
		fmt.Println("process data json Marshal ", err)
		return
	}

	client.SendMsg(headByte)

	fmt.Println("acc_response send", client.Addr, client.AppId, client.UserId, "cmd", cmd, "code", code)
}
