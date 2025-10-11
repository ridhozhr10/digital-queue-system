package app

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	app := &cli.App{
		Name:   "auth-service",
		Usage:  "Authentication service for the digital queue system",
		Before: beforeActions,
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Usage:  "Start the web server",
				Flags:  serveFlags,
				Action: serveActions,
			},
			{
				Name:   "swagger",
				Usage:  "Serve Swagger for auth-service api",
				Flags:  swaggerFlags,
				Action: swaggerAction,
			},
		},
	}
	return app
}

func Run() {
	app := NewApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to run application")
	}
}
