package client

import (
	"context"
	"fmt"
	"net/url"

	"github.com/rs/zerolog"
	"golang.org/x/exp/slices"

	"github.com/beeper/faye-client/pkg/message"
	"github.com/beeper/faye-client/pkg/messages"
	"github.com/beeper/faye-client/pkg/transport"
)

type Client struct {
	url *url.URL
	log *zerolog.Logger

	clientID string
	advice   *message.Advice

	longPollingTransport *transport.LongPollingTransport

	websocketsSupported bool
	websocketTransport  *transport.WebsocketTransport

	subscriptions *Subscriptions
}

func NewClient(url *url.URL, log *zerolog.Logger) *Client {
	return &Client{
		url: url,
		log: log,

		longPollingTransport: transport.NewLongPollingTransport(url),
		websocketTransport:   transport.NewWebsocketTransport(url),

		subscriptions: NewSubscriptions(),
	}
}

func (c *Client) Handshake(ctx context.Context) error {
	log := c.log.With().Str("method", "Handshake").Logger()
	ctx = log.WithContext(ctx)

	messages, err := c.longPollingTransport.Send(ctx, messages.NewHandshakeMessage())
	if err != nil {
		return err
	}

	if len(messages) != 1 {
		return fmt.Errorf("expected 1 message, got %d", len(messages))
	}

	msg := messages[0]

	if msg.Channel != "/meta/handshake" {
		return fmt.Errorf("unexpected message with channel %s", msg.Channel)
	}

	if !msg.Successful {
		return fmt.Errorf("handshake failed with error: %v", msg.Error)
	}

	c.websocketsSupported = slices.Contains(msg.SupportedConnectionTypes, "websocket")

	if c.clientID != msg.ClientID {
		c.clientID = msg.ClientID
		// TODO resubscribe all
	}

	c.handleAdvice(msg.Advice)

	return nil
}

func (c *Client) handleAdvice(advice *message.Advice) {
	if advice != nil {
		c.log.Debug().Interface("advice", advice).Msg("received advice")
		c.advice = advice
	}
}

func (c *Client) SubscribeWithExt(ctx context.Context, channel string, ext map[string]any, callback func(*message.Message)) error {
	log := c.log.With().Str("method", "Subscribe").Logger()
	ctx = log.WithContext(ctx)

	messages, err := c.longPollingTransport.Send(ctx, messages.NewSubscribeMessageWithExt(c.clientID, channel, ext))
	if err != nil {
		return err
	}

	if len(messages) != 1 {
		return fmt.Errorf("expected 1 message, got %d", len(messages))
	}

	msg := messages[0]

	if msg.Channel != "/meta/subscribe" {
		return fmt.Errorf("unexpected message with channel %s", msg.Channel)
	}

	if !msg.Successful {
		return fmt.Errorf("subscribe failed with error: %v", msg.Error)
	}

	c.handleAdvice(msg.Advice)

	c.subscriptions.Add(channel, callback)
	return nil
}
