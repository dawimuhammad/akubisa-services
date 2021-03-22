package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

// NewUserHandler ..
func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// get input from user
	// map input dari user struct RegisterUserInput
	// struct di atas dipassing sebagai parameter service

	var input user.RegisterUserInput

	fmt.Println("userInput--> ", input)

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.RegisterUser(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	response := helper.APIResponse("Account has successfully created!", http.StatusOK, "success", user)

	c.JSON(http.StatusOK, response)
}
