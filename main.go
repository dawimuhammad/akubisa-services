package main

import (
	"bwastartup/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:blackblair@tcp(127.0.0.1:3306)/akubisa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if guard if there is error between code and db, program will throw err and not continue
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}

	userInput.Name = "Nurul Zahwa"
	userInput.Email = "nurulzahwa@mail.co"
	userInput.Occupation = "Accounting"
	userInput.Password = "ya password"

	userService.RegisterUserInput(userInput)
}
