package main

import (
	"context"
	"net/url"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/beeper/faye-client/pkg/client"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal().Msg("missing url")
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
}
