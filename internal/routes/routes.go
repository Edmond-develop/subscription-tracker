package internal

import (
	"database/sql"
	internal "github.com/Edmond-develop/subscription-tracker/internal/hendlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.POST("/subscriptions", func(c *gin.Context) { internal.CreateSubscriptions(c, db) })
	r.GET("/subscriptions", func(c *gin.Context) { internal.ListSubscriptions(c, db) })
	r.GET("/subscriptions/:id", func(c *gin.Context) { internal.GetSubscription(c, db) })
	r.DELETE("/subscriptions/:id", func(c *gin.Context) { internal.DeleteSubscription(c, db) })
	r.GET("/subscriptions/summary", func(c *gin.Context) { internal.Summary(c, db) })

	return r
}
