package main

import (
	"log"

	"github.com/madyar997/practice_7/config"
	"github.com/madyar997/practice_7/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
