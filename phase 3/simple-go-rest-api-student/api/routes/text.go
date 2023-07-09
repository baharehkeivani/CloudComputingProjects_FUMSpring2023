package routes

import (
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/api/controller"
	"github.com/alireza-dehghan-nayeri/simple-go-rest-api/infrastructure"
)

// TextRoute -> Route for question module
type TextRoute struct {
	Controller controller.TextController
	Handler    infrastructure.GinRouter
}

// NewTextRoute -> initializes new choice rouets
func NewTextRoute(
	controller controller.TextController,
	handler infrastructure.GinRouter,

) TextRoute {
	return TextRoute{
		Controller: controller,
		Handler:    handler,
	}
}

// Setup -> setups new choice Routes
func (p TextRoute) Setup() {
	text := p.Handler.Gin.Group("/texts")
	{
		text.GET("/", p.Controller.GetTexts)
		text.POST("/", p.Controller.AddText)
		text.GET("/:id", p.Controller.GetText)
		text.DELETE("/:id", p.Controller.DeleteText)
		text.PUT("/:id", p.Controller.UpdateText)
	}
}
