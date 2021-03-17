package main

import (
	"bwastartup/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "root:blackblair@tcp(127.0.0.1:3306)/akubisa?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if guard if there is error between code and db, program will throw err and not continue
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// assign variable users with new array of struct User
	// var users []user.User
	// length := len(users)

	// check length of users
	//fmt.Println(length)

	// commanding db to find all users and store into variable users with array of User struct type
	//db.Find(&users)

	// now check again the length of users
	// length = len(users)
	// fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// }

	router := gin.Default()
	router.GET("/users", handler)
	router.Run()
}

func handler(c *gin.Context) {
	dsn := "root:blackblair@tcp(127.0.0.1:3306)/akubisa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var users []user.User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}
