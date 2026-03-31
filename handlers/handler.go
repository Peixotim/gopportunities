package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "API IS RUNNING !",
	})
}
func ShowOpeningHandler(ctx *gin.Context) {

}
func ListOpeningsHandler(ctx *gin.Context) {

}
func CreateOpeningHandler(ctx *gin.Context) {

}

func UpdateOpeningHandler(ctx *gin.Context) {

}

func DeleteOpeningHandler(ctx *gin.Context) {

}
