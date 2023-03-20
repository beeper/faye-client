package transport

import (
	"context"
	"net/url"

	"github.com/beeper/faye-client/pkg/message"
)

type WebsocketTransport struct {
	url *url.URL
}

func NewWebsocketTransport(url *url.URL) *WebsocketTransport {
	return &WebsocketTransport{
		url: url,
	}
}

func (t *WebsocketTransport) Send(ctx context.Context, msg *message.Message) ([]*message.Message, error) {
	panic("not implemented")
}
