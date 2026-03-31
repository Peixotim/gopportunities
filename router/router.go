package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	port := ":8080"
	initializeRoutes(router)
	router.Run(port)
}
