package app

import (
	"digital-queue-system/internal/http"
	"digital-queue-system/internal/repository"
	"digital-queue-system/internal/service"
	"digital-queue-system/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func beforeAction(c *cli.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Msg("Error loading .env file, using default environment variables")
	}
	logger.InitLogger()
	return nil
}

func authServeAction(c *cli.Context) error {
	r := gin.Default()

	if c.Bool("production") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		r.Static("/swagger-ui", "./third_party/swagger-ui")
		r.StaticFile("/swagger.yaml", "./api/auth-service.yaml")
	}

	connStr := c.String("database-url")
	if connStr == "" {
		log.Fatal().Msg("DATABASE_URL environment variable is not set")
	}

	pgRepo, err := repository.NewPostgresRepository(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	userRepo := repository.NewUserRepository(pgRepo)
	jwtSecret := c.String("jwt-secret")
	authSvc := service.NewAuthService(userRepo, jwtSecret)

	http.SetupAuthRoutes(r, authSvc)
	port := c.String("port")
	log.Info().Msgf("Starting server on port %s", port)
	r.Run(":" + port)
	return nil
}

func userServeAction(c *cli.Context) error {
	r := gin.Default()

	if c.Bool("production") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		// Add swagger-ui for user-service if needed
	}

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "User service is running!"})
	})

	port := c.String("port")
	log.Info().Msgf("Starting user service on port %s", port)
	r.Run(":" + port)
	return nil
}
