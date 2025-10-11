package app

import (
	"digital-queue-system/internal/http"
	"digital-queue-system/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func beforeActions(c *cli.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("Error loading .env file, using default environment variables")
	}
	logger.InitLogger()
	return nil
}

func serveActions(c *cli.Context) error {
	r := gin.Default()

	if c.Bool("production") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		r.Static("/swagger-ui", "./third_party/swagger-ui")
		r.StaticFile("/swagger.yaml", "./api/auth-service.yaml")
	}

	http.SetupRoutes(r)
	port := c.String("port")
	log.Info().Msgf("Starting server on port %s", port)
	r.Run(":" + port)
	return nil
}
