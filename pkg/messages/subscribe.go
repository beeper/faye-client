package messages

import "github.com/beeper/faye-client/pkg/message"

const ChannelSubscribe = "/meta/subscribe"

func NewSubscribeMessage(clientID, subscription string) *message.Message {
	return &message.Message{
		Channel:                  ChannelSubscribe,
		ClientID:                 clientID,
		Subscription:             subscription,
		SupportedConnectionTypes: []message.ConnectionType{message.ConnectionTypeLongPolling},
	}
}
