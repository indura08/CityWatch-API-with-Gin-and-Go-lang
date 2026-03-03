package handler

import (
	"citywatch/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (h *UserHandler) Register(c *gin.Context) {
	//var registerDto *user.RegisterDto
}
