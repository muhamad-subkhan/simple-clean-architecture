package main

import (
	"clean/config"
	"clean/user/handler"
	"clean/user/repo"
	"clean/user/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.Database()

	router := gin.Default()

	port := os.Getenv("PORT")

	userRepo := repo.UserRepo(db)
	userUsecase := usecase.CreateUserUsecase(userRepo)

	handler.Routes(router, userUsecase)
	
	err := router.Run(":"+ port)
	if err != nil {
		log.Fatal(err)
	}

}