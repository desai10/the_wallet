package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
