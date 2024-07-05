package routes

import (
	"GoMvcPattern/handlers"

	"github.com/gin-gonic/gin"
)

func WebRoutes(r *gin.Engine) {
	r.GET("/", handlers.HomeHandler)
}
