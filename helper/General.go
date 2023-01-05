package helper

import (
	"github.com/gin-gonic/gin"
)

func AppendHeader(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if c.Request.Method == "OPTIONS" {
		c.JSON(200, gin.H{"success": 1})
		panic("done")
	}
}
