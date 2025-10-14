package app

import "github.com/urfave/cli/v2"

var authServeFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "8080",
		Usage:   "Port to run the server on",
		EnvVars: []string{"AUTH_SVC_PORT"},
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
		Name:     "jwt-secret",
		Aliases:  []string{"jwt"},
		Usage:    "JWT secret key for authentication",
		EnvVars:  []string{"JWT_SECRET"},
		Required: true,
	},
}

var userServeFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Value:   "8082",
		Usage:   "Port to run the server on",
		EnvVars: []string{"USER_SVC_PORT"},
	},
	&cli.BoolFlag{
		Name:    "production",
		Aliases: []string{"prod"},
		Value:   false,
		Usage:   "Environment mode to use (true = production)",
		EnvVars: []string{"PRODUCTION"},
	},
}
