package routes

import (
	"github.com/Edmond-develop/subscription-tracker/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(handler *controller.SubscriptionHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/subscriptions", handler.Create)
	r.GET("/subscriptions", handler.GetAll)
	r.GET("/subscriptions/:id", handler.GetID)
	r.DELETE("/subscriptions/:id", handler.Delete)
	r.GET("/subscriptions/summary", handler.Summary)

	return r
}
