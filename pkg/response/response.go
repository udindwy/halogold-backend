package response

import (
	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error": message,
	})
}

func Success(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}
