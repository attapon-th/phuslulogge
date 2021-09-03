package main

import (
	"github.com/attapon-th/phuslulogger"
	"github.com/phuslu/log"
)

func main() {
	// Set Default log

	phuslulogger.SetDefaultlogger()

	log.Debug().Msg("Default logger DEGUG")
	log.Info().Msg("Default logger INFO")
	log.Warn().Msg("Default logger Warning")
	log.Error().Msg("Default logger Error")
	log.Fatal().Msg("Default logger Fatal")
}
