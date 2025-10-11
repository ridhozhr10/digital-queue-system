package auth

import "github.com/gin-gonic/gin"

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func (h *AuthHandler) Register(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "register",
	})
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}
