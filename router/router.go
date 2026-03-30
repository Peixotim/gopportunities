package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()
	port := ":8080"
	router.Run(port)
}
