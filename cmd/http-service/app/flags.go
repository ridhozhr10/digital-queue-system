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
	&cli.StringFlag{
		Name:    "database-url",
		Aliases: []string{"db-url"},
		Usage:   "Database connection string",
		EnvVars: []string{"DATABASE_URL"},
	},
	&cli.StringFlag{
		Name:    "jwt-secret",
		Aliases: []string{"jwt"},
		Usage:   "JWT secret key for authentication",
		EnvVars: []string{"JWT_SECRET"},
		Required: true,
	},
}
