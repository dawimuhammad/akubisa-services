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
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Login Failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}

		response := helper.APIResponse("Login Failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, "token")

	response := helper.APIResponse("Sucessfully Logged In!", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

	// flow process of Login API endpoint
	// client mengirimkan input berupa data email dan password user
	// input diterima handler
	// mapping dari input user ke Struct Input
	// Struct Input dikirimkan ke service
	// service akan mencari user dengan parameter email yang sama dibantu repository
	// mencocokkan password berdasarkan data user yang match
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	var metaMessage string = "Email has been registered"

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"error": errors}

		response := helper.APIResponse("Email Checking Failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"error": "Server Error"}

		response := helper.APIResponse("Email Checking Failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// melakukan wrapping data yang dibutuhkan ke dalam bentuk objek
	data := gin.H{
		"is_available": isEmailAvailable,
	}

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

	// flow process of check email availability API endpoint
	// client mengirimkan input berupa email
	// input diterima handler
	// mapping dari input user ke Struct Input
	// Struct Input dikirimkan ke service
	// service akan memanggil repository (existing) untuk mencari apakah sudah ada user yang menggunakan email tersebut
	// repository db mencari email berdasarkan parameter input yang disediakan
}
