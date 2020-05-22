package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	engine.Use(CORSMiddleware())

	api := engine.Group("/cms/api")
	{
		api.StaticFile("/doc", "../web/static/redoc.html")
		api.StaticFile("/openapi.yaml", "../api/openapi.yaml")
	}

	return engine
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
