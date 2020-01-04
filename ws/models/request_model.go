package models

type Request struct {
	Seq  string      `json:"seq"`
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data,omitempty"`
}

type Login struct {
	ServiceToken string `json:"serviceToken"`
	AppId        uint32 `json:"appId,omitempty"`
	UserId       string `json:"userId,omitempty"`
}

type HeartBeat struct {
	UserId string `json:"userId,omitempty"`
}
