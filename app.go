package application

import (
	"log"

	"github.com/ericoliveiras/gate-guard/internal/config"
	"github.com/ericoliveiras/gate-guard/internal/server"
)

func Start(config *config.Config) {
	app := server.NewServer(config)

	err := app.Start(config.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
