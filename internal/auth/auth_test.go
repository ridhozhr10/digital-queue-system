package auth_test

import (
	"digital-queue-system/internal/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAuthHandler(t *testing.T) {
	h := auth.NewAuthHandler()
	assert.NotNil(t, h, "NewAuthHandler should return a non-nil handler")
}
