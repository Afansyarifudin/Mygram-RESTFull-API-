package app

import (
	"fmt"
	"mygram/config"
	"mygram/controllers"
	"mygram/repository"
	"mygram/route"
	"mygram/services"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.GORM.DB)
	service := services.NewService(repo)
	server := controllers.NewHttpServer(service)

	route.RegisterApi(router, server)
	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))

}
