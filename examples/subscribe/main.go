package main

import (
	"context"
	"net/url"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/beeper/faye-client/pkg/client"
	"github.com/beeper/faye-client/pkg/message"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal().Msg("missing url or channel")
	}

	url, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse url")
	}

	log.Info().Msg("sending handshake")
	c := client.NewClient(url, &log.Logger)
	err = c.Handshake(context.TODO())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to send handshake")
	}
	log.Info().Msg("handshake successful")

	log.Info().Msg("subscribing to channel")
	channel := os.Args[2]
	err = c.Subscribe(context.TODO(), channel, func(m *message.Message) {
		log.Info().Interface("msg", m).Msg("received message")
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to subscribe")
	}
}
