package routes

import (
	"github.com/gin-gonic/gin"
	"gorm-gin-postgres/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/items", handlers.CreateItem)
		api.GET("/items", handlers.GetAllItems)
		api.GET("/items/:id", handlers.GetItem)
		api.PUT("/items/:id", handlers.UpdateItem)
		api.DELETE("/items/:id", handlers.DeleteItem)
	}
	return router
}
