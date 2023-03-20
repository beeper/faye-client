package messages

import "github.com/beeper/faye-client/pkg/message"

const ChannelHandshake = "/meta/handshake"

func NewHandshakeMessage() *message.Message {
	return &message.Message{
		Channel:                  ChannelHandshake,
		Version:                  "1.0",
		SupportedConnectionTypes: []message.ConnectionType{message.ConnectionTypeLongPolling},
	}
}
