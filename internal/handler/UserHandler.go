package handler

import (
	auth "citywatch/internal/Dto/Auth"
	user "citywatch/internal/Dto/User"
	"citywatch/internal/enums"
	"citywatch/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

// citizen registration
func (h *UserHandler) CitizenRegister(c *gin.Context) {
	var registerDto *user.RegisterDto

	if err := c.ShouldBindJSON(&registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request body",
		})
		return
	}

	err := h.userService.Register(registerDto, enums.Citizen)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created succssfully",
	})
}

// worker registration
func (h *UserHandler) WorkerRegister(c *gin.Context) {
	var registerDto *user.RegisterDto

	if err := c.ShouldBindJSON(&registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request body",
		})
		return
	}

	err := h.userService.Register(registerDto, enums.Worker)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

// admin registration
func (h *UserHandler) AdminRegistration(c *gin.Context) {
	var registerDto *user.RegisterDto

	if err := c.ShouldBindJSON(&registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid Request body",
		})
		return
	}

	err := h.userService.Register(registerDto, enums.Admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Admin created",
	})
}

// login function
func (h *UserHandler) Login(c *gin.Context) {
	var loginDto *auth.LoginDto

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})

		return
	}

	token, err := h.userService.Login(loginDto)

	if err != nil || token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})
}
