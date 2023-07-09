package main

import (
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/api/controller"
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/api/repository"
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/api/routes"
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/api/service"
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/infrastructure"
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {

	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()

	textRepository := repository.NewTextRepository(db)
	textService := service.NewTextService(textRepository)
	textController := controller.NewTextController(textService)
	textRoute := routes.NewTextRoute(textController, router)
	textRoute.Setup()

	db.DB.AutoMigrate(&models.Text{})

	router.Gin.Run(":8000")
}
