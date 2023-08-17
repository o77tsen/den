package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/o77tsen/den/controllers"
)

func MessageRoutes(route *gin.Engine) {
	v1 := route.Group("/api")
	{
		v1.GET("/messages", controllers.GetAllMessages)
		v1.GET("/message/:id", controllers.GetMessage)
		v1.POST("/messages", controllers.CreateMessage)
		v1.DELETE("/message/:id", controllers.DeleteMessage)
	}
}