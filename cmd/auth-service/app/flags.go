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
}

var swaggerFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "8081",
		Usage:   "Port to run the Swagger UI server on",
		EnvVars: []string{"SWAGGER_PORT"},
	},
}
