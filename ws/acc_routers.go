package ws

func WebSocketRoutersInit() {
	Register("login", LoginController)
	Register("heartbeat", HeartbeatController)
	Register("ping", PingController)
}
