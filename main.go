package main

import (
	"GoMvcPattern/app"
	"GoMvcPattern/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	app.Boot()
}

func main() {
	fmt.Print("Hello World!")
	if os.Getenv("APP_ENV") == "local" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	routes.WebRoutes(r)

	r.Run(":4000")
}
