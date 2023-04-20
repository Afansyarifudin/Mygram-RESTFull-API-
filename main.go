package main

import (
	"mygram/app"
	"mygram/config"
	_ "mygram/docs"

	"github.com/joho/godotenv"
)

// @title MyGram RESTFull-API
// @Version 1.0
// @description A My-Gram RESTFull API using Gin + GORM with JWT Auth
// @contact.name Nur Afan Syarifudin
// @contact.email afan.syarifudin10@gmail.com
// @license.name MIT License

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used
// @host localhost:8081
// @BasePath /

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	_, err = config.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
