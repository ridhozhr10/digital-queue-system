package app_test

import (
	"digital-queue-system/cmd/auth-service/app"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNewAppCommands(t *testing.T) {
	cliApp := app.NewApp()

	assert.NotNil(t, cliApp)
	assert.Equal(t, "auth-service", cliApp.Name)

	// Check for the 'serve' command
	serveCommand := cliApp.Command("serve")
	assert.NotNil(t, serveCommand)
	assert.Equal(t, "Start the web server", serveCommand.Usage)
	assert.Len(t, serveCommand.Flags, 1)
	servePortFlag := serveCommand.Flags[0].(*cli.StringFlag)
	assert.Equal(t, "port", servePortFlag.Name)
	assert.Equal(t, "8080", servePortFlag.Value)
	assert.Contains(t, servePortFlag.EnvVars, "PORT")

	// Check for the 'swagger' command
	swaggerCommand := cliApp.Command("swagger")
	assert.NotNil(t, swaggerCommand)
	assert.Equal(t, "Serve Swagger for auth-service api", swaggerCommand.Usage)
	assert.Len(t, swaggerCommand.Flags, 1)
	swaggerPortFlag := swaggerCommand.Flags[0].(*cli.StringFlag)
	assert.Equal(t, "port", swaggerPortFlag.Name)
	assert.Equal(t, "8081", swaggerPortFlag.Value)
	assert.Contains(t, swaggerPortFlag.EnvVars, "SWAGGER_PORT")
}

func TestEnvLoading(t *testing.T) {
	// Create a temporary .env file
	envContent := []byte("TEST_VAR=test_value")
	err := os.WriteFile(".env", envContent, 0644)
	assert.NoError(t, err)
	defer os.Remove(".env") // Clean up the .env file after the test

	cliApp := app.NewApp()
	assert.NotNil(t, cliApp)

	// Create a dummy cli.Context
	ctx := cli.NewContext(cliApp, nil, nil)

	// Call the Before hook directly
	err = cliApp.Before(ctx)
	assert.NoError(t, err)

	// Check if the environment variable is loaded
	assert.Equal(t, "test_value", os.Getenv("TEST_VAR"))

	// Clean up the environment variable to avoid side effects on other tests
	os.Unsetenv("TEST_VAR")
}
