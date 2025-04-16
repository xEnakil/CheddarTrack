package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xenakil/cheddartrack/internal/service"
)

type AuthHandler struct {
	userService service.UserService
}

func NewAuthHandler(us service.UserService) *AuthHandler {
	return &AuthHandler{us}
}

func (h *AuthHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// @Summary Register new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body RegisterInput true "Register Input"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userService.Register(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered"})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body LoginInput true "Login Input"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.userService.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token, "token_type": "bearer", "expires_in": 3600})
}
