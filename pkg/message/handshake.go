package message

const ChannelHandshake = "/meta/handshake"

func NewHandshakeMessage() *Message {
	return &Message{
		Channel:                  ChannelHandshake,
		Version:                  "1.0",
		SupportedConnectionTypes: []ConnectionType{ConnectionTypeLongPolling},
	}
}
