package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	logLevel := strings.ToLower(os.Getenv("LOG_LEVEL"))
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	switch logLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
