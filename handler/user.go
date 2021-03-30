package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
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

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Failed to create account!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Failed to create account!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "token")

	response := helper.APIResponse("Account has successfully created!", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// flow process of Login API endpoint
	// client mengirimkan input berupa data email dan password user
	// input diterima handler
	// mapping dari input user ke Struct Input
	// Struct Input dikirimkan ke service
	// service akan mencari user dengan parameter email yang sama dibantu repository
	// mencocokkan password berdasarkan data user yang match
}
