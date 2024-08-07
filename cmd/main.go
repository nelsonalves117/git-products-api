package main

import (
	"github.com/rs/zerolog/log"

	"github.com/nelsonalves117/go-products-api/internal/channels/rest"
	"github.com/nelsonalves117/go-products-api/internal/config"
)

func main() {
	config.Parse()

	server := rest.New()

	err := server.Start()
	if err != nil {
		log.Panic().Err(err).Msg("an error occurred while trying to start the server")
	}
}
