package message

type ConnectionType string

const (
	ConnectionTypeLongPolling ConnectionType = "long-polling"
	ConnectionTypeWebSocket   ConnectionType = "websocket"
)
