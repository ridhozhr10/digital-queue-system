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
		Before: beforeAction,
		Commands: []*cli.Command{
			{
				Name:   "auth-service",
				Usage:  "Start auth service web server",
				Flags:  authServeFlags,
				Action: authServeAction,
			},
			{
				Name:   "user-service",
				Usage:  "Start user service web server",
				Flags:  userServeFlags,
				Action: userServeAction,
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
