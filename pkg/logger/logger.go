package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Setup(env string) {
	if env == "prod" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if env == "local" || env == "dev" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			NoColor:    false,
		})
	} else {
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
}
