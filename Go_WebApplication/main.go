package main

import (
	"Go_WebApplication/config"
	"Go_WebApplication/logger"
	"Go_WebApplication/routes"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	var err error
	// log initiation
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false})

	// =========================== Setup logger =========================== //
	err = logger.SetupLogger(logger.Log)
	if err != nil {
		log.Error().Err(err).Msg("error while setup logger")
		logger.Log.Error().Err(err).Msg("error while setup logger")
		return
	}
	logger.Log.Info().Msg("Enter for main()...")

	config.ConnectDatabase()
	r := routes.SetupRoutes()
	r.Run(":8080") // Start the server
}
