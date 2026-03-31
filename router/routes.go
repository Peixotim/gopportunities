package router

import (
	handler "github.com/Peixotim/gopportunities/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/health", handler.Health)
		v1.GET("/openning", handler.ShowOpeningHandler)
		v1.GET("/opennings", handler.ListOpeningsHandler)
		v1.POST("/openning", handler.CreateOpeningHandler)
		v1.PUT("/openning", handler.UpdateOpeningHandler)
		v1.DELETE("/openning", handler.DeleteOpeningHandler)
	}
}
