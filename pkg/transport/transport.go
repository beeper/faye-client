package transport

import (
	"context"

	"github.com/beeper/faye-client/pkg/message"
)

type Transport interface {
	Send(ctx context.Context, msg *message.Message) ([]*message.Message, error)
}
