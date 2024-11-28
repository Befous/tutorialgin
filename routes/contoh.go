package routes

import (
	"github.com/Befous/tutorialgin/handler"
	"github.com/gin-gonic/gin"
)

func ContohRoute(page *gin.Engine) {
	page.GET("/routeget/:z", handler.ContohHandler)
	page.POST("/routepost/:z", handler.ContohHandler)
	page.PUT("/routeput/:z", handler.ContohHandler)
	page.DELETE("/routedelete/:z", handler.ContohHandler)
}
