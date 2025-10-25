package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	SetupRoutes(r)

	r.Run(":3001")
}
