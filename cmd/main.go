package main

import (
	application "github.com/ericoliveiras/gate-guard"
	"github.com/ericoliveiras/gate-guard/internal/config"
)

func main() {
	config := config.NewConfig()

	application.Start(config)
}
