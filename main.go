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

	user := user.User{
		ID:             3,
		Name:           "Naura Safa",
		Occupation:     "Bermain",
		Email:          "naura@dududu.co",
		PasswordHash:   "pwhash",
		AvatarFileName: "avatarnau.jpg",
		Role:           "user",
	}

	userRepository.Save(user)
}
