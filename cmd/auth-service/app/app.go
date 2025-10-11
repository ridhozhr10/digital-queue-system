package app

import (
	"digital-queue-system/internal/http"
	"digital-queue-system/pkg/logger"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	app := &cli.App{
		Name:  "auth-service",
		Usage: "Authentication service for the digital queue system",
		Before: func(c *cli.Context) error {
			err := godotenv.Load()
			if err != nil {
				log.Warn().Msg("Error loading .env file, using default environment variables")
			}
			logger.InitLogger()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "Start the web server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   "8080",
						Usage:   "Port to run the server on",
						EnvVars: []string{"PORT"},
					},
				},
				Action: func(c *cli.Context) error {
					r := gin.Default()
					http.SetupRoutes(r)
					port := c.String("port")
					log.Info().Msgf("Starting server on port %s", port)
					r.Run(":" + port)
					return nil
				},
			},
			{
				Name:  "swagger",
				Usage: "Serve Swagger for auth-service api",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   "8081",
						Usage:   "Port to run the Swagger UI server on",
						EnvVars: []string{"SWAGGER_PORT"},
					},
				},
				Action: func(c *cli.Context) error {
					r := gin.Default()
					r.Static("/swagger-ui", "./third_party/swagger-ui")
					r.StaticFile("/swagger.yaml", "./api/auth-service.yaml")
					swaggerPort := c.String("port")
					log.Info().Msgf("Swagger UI available at http://localhost:%s/swagger-ui", swaggerPort)
					r.Run(":" + swaggerPort)
					return nil
				},
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
