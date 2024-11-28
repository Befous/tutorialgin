package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var headers = []string{
	"Origin",
	"Content-Type",
	"Accept",
	"Authorization",
	"Access-Control-Request-Headers",
	"Token",
	"Login",
	"Access-Control-Allow-Origin",
	"Bearer",
	"X-Requested-With",
}

var CorsAllowAll = cors.Config{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	AllowHeaders: headers,
	MaxAge:       2 * time.Hour,
}

func NewCorsMiddleware() gin.HandlerFunc {
	return cors.New(CorsAllowAll)
}
