package http_test

import (
	"digital-queue-system/internal/http"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetupRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	http.SetupRoutes(r)

	// Test /ping route
	w := httptest.NewRecorder()
	req, _ := nethttp.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}
