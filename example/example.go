package main

import (
	phuslulogger "github.com/attapon-th/phuslulogger"
	"github.com/phuslu/log"
)

func main() {
	// Set Default log

	phuslulogger.SetDefaultlogger()

	log.Info().Msg("Default logger INFO")
	log.Info().Msg("Default logger INFO")
}
