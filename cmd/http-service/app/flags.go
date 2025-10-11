package app

import "github.com/urfave/cli/v2"

var serveFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "8080",
		Usage:   "Port to run the server on",
		EnvVars: []string{"PORT"},
	},
	&cli.BoolFlag{
		Name:    "production",
		Aliases: []string{"prod"},
		Value:   false,
		Usage:   "Environment mode to use (true = production)",
		EnvVars: []string{"PRODUCTION"},
	},
}
