package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello world")
}
