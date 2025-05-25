package main

import (
	"bwastartup/user"
  "bwastartup/handler"
	"bwastartup/auth"
	"log"

	"fmt"
  "github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:MorishitaAi87*%@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	token,err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMn0.HIRKlgpxIlLOsIfsbGI5OqNlDMvszpE0HKCxaYtPu_8")
	if err != nil{
						fmt.Println("ERROR")
						fmt.Println("ERROR")
						fmt.Println("ERROR")
	}
	if token.Valid {
					 fmt.Println("VALID")
					 fmt.Println("VALID")
					 fmt.Println("VALID")
	} else{
					fmt.Println("INVALID")
					fmt.Println("INVALID")
					fmt.Println("INVALID")
	}

	fmt.Println(authService.GenerateToken(1001))
	
  userHandler := handler.NewUserHandler(userService,authService)

  router := gin.Default()
  api:= router.Group("/api/v1")

  api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)
 
  router.Run()

}
