package main

import (
	"github.com/Befous/tutorialgin/middleware"
	"github.com/Befous/tutorialgin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(middleware.NewCorsMiddleware())
	routes.ContohRoute(app)
	app.Run(":3000")
}
